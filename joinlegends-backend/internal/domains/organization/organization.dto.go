package organization

type CreateOrganizationDto struct {
	Name string `json:"name"`
}

type UpdateOrganizationDto struct {
	Name string `json:"name"`
}

type AddMemberDto struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"` // ADMIN, MEMBER
}
