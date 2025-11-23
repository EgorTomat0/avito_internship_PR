package error

var NotFound = CustomError{
	Code:    ErrNotFound,
	Message: "resource not found",
}

var NotAssigned = CustomError{
	Code:    ErrNotAssigned,
	Message: "reviewer is not assigned to this PR",
}

var NoCandidate = CustomError{
	Code:    ErrNoCandidate,
	Message: "no active replacement candidate in team",
}

var PrMerged = CustomError{
	Code:    ErrPRMerged,
	Message: "cannot reassign on merged PR",
}

var PrExists = CustomError{
	Code:    ErrPRExists,
	Message: "PR id already exists",
}

var TeamExists = CustomError{
	Code:    ErrTeamExists,
	Message: "team_name already exists",
}
