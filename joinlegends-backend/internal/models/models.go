package models

import "time"

type User struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Role      string    `json:"role" db:"role"` // ADMIN, USER, ROOT
}

type Recovery struct {
	ID        int       `db:"id"`
	UserID    string    `db:"user_id"`
	Email     string    `db:"email"`
	Code      string    `db:"code"`
	Attempts  int       `db:"attempts"`
	Expired   bool      `db:"expired"`
	ExpiresAt time.Time `db:"expires_at"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Session struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	Browser   *string   `json:"browser" db:"browser"`
	IP        *string   `json:"ip" db:"ip"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
}

type Video struct {
	ID             string     `json:"id" db:"id"`
	FileName       string     `json:"file_name" db:"file_name"`
	FilePath       string     `json:"file_path" db:"file_path"`
	FileSize       int64      `json:"file_size" db:"file_size"`
	FileType       string     `json:"file_type" db:"file_type"`
	FileHash       string     `json:"file_hash" db:"file_hash"`
	FileDuration   int64      `json:"file_duration" db:"file_duration"`
	FileResolution string     `json:"file_resolution" db:"file_resolution"`
	FileBitrate    int64      `json:"file_bitrate" db:"file_bitrate"`
	FileCodec      string     `json:"file_codec" db:"file_codec"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at" db:"deleted_at"`
	OwnerID        string     `json:"owner_id" db:"owner_id"`
	OwnerName      string     `json:"owner_name" db:"owner_name"`
}

type Organization struct {
	ID        string     `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	OwnerID   string     `json:"owner_id" db:"owner_id"`
	BannerID  *string    `json:"banner_id" db:"banner_id"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`

	Members []Member `json:"members" db:"-"`
}

type Member struct {
	ID             string    `json:"id" db:"id"`
	UserID         string    `json:"user_id" db:"user_id"`
	OrganizationID string    `json:"organization_id" db:"organization_id"`
	Role           string    `json:"role" db:"role"` // ADMIN, MEMBER, OWNER
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

type Events struct {
}
