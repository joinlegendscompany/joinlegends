package models

import "time"

type Organization struct {
	ID          string     `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	OwnerID     string     `json:"owner_id" db:"owner_id"`
	BannerID    *string    `json:"banner_id" db:"banner_id"`
	ValidatedAt *time.Time `json:"validated_at" db:"validated_at"` // nil = pendente, preenchido = pública/validada
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`

	Members []Member `json:"members" db:"-"`
}

type Member struct {
	ID             string    `json:"id" db:"id"`
	UserID         string    `json:"user_id" db:"user_id"`
	OrganizationID string    `json:"organization_id" db:"organization_id"`
	Role           string    `json:"role" db:"role"` // ADMIN, MEMBER, OWNER
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`

	Games []Game `json:"games" db:"-"`
}

type MemberGame struct {
	ID        string    `json:"id" db:"id"`
	MemberID  string    `json:"member_id" db:"member_id"`
	GameID    string    `json:"game_id" db:"game_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
