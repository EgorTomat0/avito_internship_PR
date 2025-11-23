package output

import "avito_internship_PR/internal/pr/application/dto"

type ReassignReviewerOutput struct {
	PR         *dto.PullRequestDto
	ReplacedBy string
}
