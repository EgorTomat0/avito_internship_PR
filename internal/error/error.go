package error

type ErrCode string

const (
	ErrTeamExists  ErrCode = "TEAM_EXISTS"
	ErrPRExists    ErrCode = "PR_EXISTS"
	ErrPRMerged    ErrCode = "PR_MERGED"
	ErrNotAssigned ErrCode = "NOT_ASSIGNED"
	ErrNoCandidate ErrCode = "NO_CANDIDATE"
	ErrNotFound    ErrCode = "NOT_FOUND"
)

type CustomError struct {
	Code    ErrCode
	Message string
}

func (e CustomError) Error() string {
	return e.Message
}

type ErrResponse struct {
	Error CustomError `json:"error"`
}
