package views

type PersonView struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type PersonName struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
