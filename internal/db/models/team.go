package models

import (
	"time"
)

type Team struct {
	ID           uint          `gorm:"primaryKey;autoIncrement"`
	TeamName     string        `gorm:"unique;uniqueIndex;not null"`
	TeamMembers  []TeamMember  `gorm:"foreignKey:TeamID"`
	PullRequests []PullRequest `gorm:"foreignKey:TeamID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
