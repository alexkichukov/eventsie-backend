package events

import (
	"context"
	"events/api/config"
	pb "events/pb/events"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.EventsClient
}

func NewServiceClient() *ServiceClient {
	cfg := config.GetConfig()
	serviceURL := fmt.Sprintf("localhost:%d", cfg.EVENTS_SERVICE_PORT)

	conn, err := grpc.Dial(serviceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Couldn't connect to service:", err)
	}

	return &ServiceClient{
		Client: pb.NewEventsClient(conn),
	}
}

func (svc *ServiceClient) GetEventByID(c *fiber.Ctx) error {
	resp, err := svc.Client.FindOne(context.TODO(), &pb.FindOneRequest{Id: c.Params("id")})

	// Could not connect to the events service
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not establish connection to the events service.",
		})
	}

	// There is an error
	if len(resp.Error) != 0 {
		return c.Status(int(resp.Status)).JSON(fiber.Map{
			"message": resp.Error,
		})
	}

	// There was no event found
	if resp.Event == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("No event with ID %s was found", c.Params("id")),
		})
	}

	return c.Status(int(resp.Status)).JSON(resp.Event)
}

func (svc *ServiceClient) GetAllEvents(c *fiber.Ctx) error {
	category := c.Query("category")
	priceFrom := c.Query("priceFrom")
	priceTo := c.Query("priceTo")

	request := &pb.FindManyRequest{}

	if len(category) > 0 {
		request.Category = &category
	}
	if len(priceFrom) > 0 || len(priceTo) > 0 {
		request.Price = &pb.Price{}
	}
	if len(priceFrom) > 0 {
		from, err := strconv.ParseFloat(priceFrom, 64)
		if err == nil {
			request.Price.From = from
		}
	}
	if len(priceTo) > 0 {
		to, err := strconv.ParseFloat(priceTo, 64)
		if err == nil {
			request.Price.To = to
		}
	}

	resp, err := svc.Client.FindMany(context.TODO(), request)

	// Could not connect to the events service
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not establish connection to the events service.",
		})
	}

	// There is an error
	if len(resp.Error) != 0 {
		return c.Status(int(resp.Status)).JSON(fiber.Map{
			"message": resp.Error,
		})
	}

	// There was no event found
	if resp.Events == nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "No events found.",
		})
	}

	return c.Status(int(resp.Status)).JSON(resp.Events)
}
