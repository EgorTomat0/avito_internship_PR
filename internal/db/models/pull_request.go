package models

import (
	"time"
)

type PullRequest struct {
	ID              string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	PullRequestName string
	Status          PRStatus     `gorm:"type:pr_status;index;default:'OPEN';not null"`
	MergedAt        time.Time    `gorm:"type:timestamp"`
	AuthorID        string       `gorm:"not null"`
	Author          User         `gorm:"foreignKey:AuthorID;references:ID"`
	TeamID          uint         `gorm:"not null"`
	Team            Team         `gorm:"foreignKey:TeamID;references:ID"`
	Reviewers       []TeamMember `gorm:"many2many:pull_request_reviewers;"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
