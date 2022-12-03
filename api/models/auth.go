package models

type RegisterBody struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

type LoginBody struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type FavouriteAttendBody struct {
	EventID string `json:"eventID"`
}
