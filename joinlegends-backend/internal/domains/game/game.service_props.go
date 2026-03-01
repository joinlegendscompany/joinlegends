package game

type CreateGameServiceProps struct {
	Name        string
	BannerID    *string
	Category    string
	Description string
	Developer   string
	Publisher   *string
	ReleaseYear int
}

type UpdateGameServiceProps struct {
	ID          string
	Name        *string
	BannerID    *string
	Category    *string
	Description *string
	Developer   *string
	Publisher   *string
	ReleaseYear *int
	IsActive    *bool
}
