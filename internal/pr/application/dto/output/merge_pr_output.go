package output

import (
	"time"

	"avito_internship_PR/internal/pr/application/dto"
)

type MergePROutput struct {
	PR       *dto.PullRequestDto
	MergedAt time.Time
}
