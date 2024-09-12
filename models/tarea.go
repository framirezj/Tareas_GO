package models

// Struct para tarea
type Tarea struct {
	Title       string `json: "title"`
	Description string `json: "description"`
	Completed   bool   `json: "completed"`
}
