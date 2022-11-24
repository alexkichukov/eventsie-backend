package models

type EventPrice struct {
	From float64 `json:"from" validate:"nonzero"`
	To   float64 `json:"to"`
}

type EventLocation struct {
	Address  string `json:"address" validate:"nonzero"`
	City     string `json:"city" validate:"nonzero"`
	Postcode string `json:"postcode" validate:"nonzero"`
}

type CreateEventBody struct {
	Title       string         `json:"title" validate:"nonzero"`
	Date        string         `json:"date" validate:"nonzero"`
	Description string         `json:"description" validate:"nonzero"`
	Tags        []string       `json:"tags" validate:"nonzero"`
	Category    string         `json:"category" validate:"nonzero"`
	Location    *EventLocation `json:"location" validate:"nonnil"`
	Price       *EventPrice    `json:"price" validate:"nonnil"`
}
