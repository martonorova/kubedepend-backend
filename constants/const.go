package constants

import "fmt"

const (
	ROUTE_ALL_JOB = "/jobs"
)

var (
	ROUTE_SINGLE_JOB = fmt.Sprintf("%s/:id", ROUTE_ALL_JOB)
)
