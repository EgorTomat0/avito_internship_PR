package dto

type TeamDto struct {
	Name    string
	Members []TeamMemberDto
}

type TeamMemberDto struct {
	UserID   string
	Username string
	IsActive bool
}
