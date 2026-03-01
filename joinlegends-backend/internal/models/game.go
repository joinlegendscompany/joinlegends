package models

import "time"

type Game struct {
	ID          string     `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	BannerID    *string    `json:"banner_id" db:"banner_id"`
	Category    string     `json:"category" db:"category"`
	Description string     `json:"description" db:"description"`
	Developer   string     `json:"developer" db:"developer"`
	Publisher   *string    `json:"publisher" db:"publisher"`
	ReleaseYear int        `json:"release_year" db:"release_year"`
	IsActive    bool       `json:"is_active" db:"is_active"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
}
