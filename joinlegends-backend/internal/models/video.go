package models

import "time"

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
