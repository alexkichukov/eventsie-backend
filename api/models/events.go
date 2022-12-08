package models

type EventPrice struct {
	From float64 `json:"from"`
	To   float64 `json:"to"`
}

type EventLocation struct {
	Address  string `json:"address"`
	City     string `json:"city"`
	Postcode string `json:"postcode"`
}

type CreateEventBody struct {
	Title       string         `json:"title"`
	Date        string         `json:"date"`
	Description string         `json:"description"`
	Tags        []string       `json:"tags"`
	Category    string         `json:"category"`
	CreatedBy   string         `json:"createdBy"`
	Location    *EventLocation `json:"location"`
	Price       *EventPrice    `json:"price"`
}

type UpdateEventBody struct {
	Id          string         `json:"id"`
	Title       string         `json:"title"`
	Date        string         `json:"date"`
	Description string         `json:"description"`
	Tags        []string       `json:"tags"`
	Category    string         `json:"category"`
	Location    *EventLocation `json:"location"`
	Price       *EventPrice    `json:"price"`
}

type DeleteEventBody struct {
	Id string `json:"id"`
}

func (el *EventLocation) GetAddress() string {
	if el == nil {
		return ""
	}
	return el.Address
}

func (el *EventLocation) GetCity() string {
	if el == nil {
		return ""
	}
	return el.City
}

func (el *EventLocation) GetPostcode() string {
	if el == nil {
		return ""
	}
	return el.Postcode
}

func (ep *EventPrice) GetFrom() float64 {
	if ep == nil {
		return 0
	}
	return ep.From
}

func (ep *EventPrice) GetTo() float64 {
	if ep == nil {
		return 0
	}
	return ep.To
}
