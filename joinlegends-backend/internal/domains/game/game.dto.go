package game

type CreateGameDto struct {
	Name        string  `json:"name"`
	BannerID    *string `json:"banner_id"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Developer   string  `json:"developer"`
	Publisher   *string `json:"publisher"`
	ReleaseYear int     `json:"release_year"`
}

type UpdateGameDto struct {
	Name        *string `json:"name"`
	BannerID    *string `json:"banner_id"`
	Category    *string `json:"category"`
	Description *string `json:"description"`
	Developer   *string `json:"developer"`
	Publisher   *string `json:"publisher"`
	ReleaseYear *int    `json:"release_year"`
	IsActive    *bool   `json:"is_active"`
}
