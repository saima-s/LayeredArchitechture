package errorPackage

import "fmt"

type MissingParams struct {
	Params []string
}

func (m MissingParams) Error() string {
	return fmt.Sprintf("You are missing %v parameter", m.Params)
}

type Error string

const (
	ErrEligibility      Error = "you are not eligible"
	DbError             Error = "No connection to DB"
	NoDataError         Error = "No record found"
	QueryExecutionError Error = "Query execution error"
	InvalidId           Error = "Invalid request ID"
	MissingBodyJson     Error = "Missing JSON Body"
	JsonParsingError    Error = "Missing JSON Body"
)

func (e Error) Error() string {
	return string(e)
}
