package models

import "time"

type Exercise struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	MuscleGroup string `json:"muscle_group"`
	UserId      *int   `json:"user_id,omitempty"` // On utilise un pointeur pour omettre l'user id dans un premier temps pour le MVP
}

type Workout struct {
	ID     int	`json:"id"`
	UserId int	`json:"user_id"`
	Date   time.Time `json:"date"`
	DurationMinutes *int  `json:"duration_minutes,omitempty"`
	Notes	*string		`json:"notes,omitempty"`
	CreatedAt time.Time	`json:"created_at"`
}