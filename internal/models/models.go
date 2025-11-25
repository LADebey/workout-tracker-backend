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


type WorkoutExercise struct {
	ID             int `json:"id"`
	WorkoutID      int `json:"workout_id"`
	ExerciseID     int `json:"exercise_id"`
	OrderInWorkout int `json:"order_in_workout"`
}


type Set struct {
	ID                  int         `json:"id"`
	WorkoutExerciseID   int         `json:"workout_exercise_id"`
	SetNumber           int         `json:"set_number"`
	Weight              float64     `json:"weight"`
	Repetitions         int         `json:"repetitions"`
	RPE                 *float32    `json:"rpe,omitempty"`
	Technique           *string     `json:"technique,omitempty"`
	Failed              bool        `json:"failed"`
	WarmUp              bool        `json:"warm_up"`
}