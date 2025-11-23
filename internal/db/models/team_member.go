package models

import (
	"time"
)

type TeamMember struct {
	ID          uint          `gorm:"primaryKey;autoIncrement"`
	UserID      string        `gorm:"uniqueIndex:idx_team_user;not null"`
	User        User          `gorm:"foreignKey:UserID;references:ID"`
	TeamID      uint          `gorm:"uniqueIndex:idx_team_user;not null"`
	Team        Team          `gorm:"foreignKey:TeamID;references:ID"`
	ReviewerPRs []PullRequest `gorm:"many2many:pull_request_reviewers;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
