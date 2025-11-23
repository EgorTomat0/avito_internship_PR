package models

import (
	"time"
)

type User struct {
	ID          string        `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Username    string        `gorm:"not null"`
	IsActive    bool          `gorm:"default:true"`
	TeamMembers []TeamMember  `gorm:"foreignKey:UserID"`
	CreatedPRs  []PullRequest `gorm:"foreignKey:AuthorID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
