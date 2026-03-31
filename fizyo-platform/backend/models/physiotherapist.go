package models

type Physiotherapist struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	City        string   `json:"city"`
	District    string   `json:"district"`
	Phone       string   `json:"phone"`
	Specialties []string `json:"specialties"`
}
