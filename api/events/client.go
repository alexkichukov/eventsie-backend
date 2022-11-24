package events

import (
	"context"
	"events/api/config"
	"events/api/models"
	pb "events/pb/events"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	validator "gopkg.in/validator.v2"
)

type ServiceClient struct {
	Client pb.EventsClient
}

type CreateEventBody struct {
	Name string `json:"name" form:"name"`
	Pass string `json:"pass" form:"pass"`
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
	city := c.Query("city")
	tags := c.Query("tags")
	priceFrom := c.Query("priceFrom")
	priceTo := c.Query("priceTo")

	request := &pb.FindManyRequest{}

	if len(category) > 0 {
		c := strings.Split(category, ",")
		request.Categories = c
	}
	if len(city) > 0 {
		c := strings.Split(city, ",")
		request.Cities = c
	}
	if len(tags) > 0 {
		c := strings.Split(tags, ",")
		request.Tags = c
	}
	if len(priceFrom) > 0 {
		from, err := strconv.ParseFloat(priceFrom, 64)
		if err == nil && from > 0 {
			request.PriceFrom = &from
		}
	}
	if len(priceTo) > 0 {
		to, err := strconv.ParseFloat(priceTo, 64)
		if err == nil && to > 0 {
			request.PriceTo = &to
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

func (svc *ServiceClient) CreateEvent(c *fiber.Ctx) error {
	e := &models.CreateEventBody{}

	if err := c.BodyParser(e); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid event data",
		})
	}

	if err := validator.Validate(e); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid event data",
		})
	}

	event := &pb.Event{
		Title:       e.Title,
		Date:        e.Date,
		Description: e.Description,
		Tags:        e.Tags,
		Category:    e.Category,
		Location: &pb.Location{
			Address:  e.Location.Address,
			City:     e.Location.City,
			Postcode: e.Location.Postcode,
		},
		Price: &pb.Price{
			From: e.Price.From,
			To:   e.Price.To,
		},
	}

	svc.Client.Add(context.TODO(), &pb.AddRequest{Event: event})

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Event created successfully",
	})
}
