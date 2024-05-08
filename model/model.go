package model

// employee represents data about a record employee.
type Employee struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Position string  `json:"title"`
	Salary   float64 `json:"salary"`
}
