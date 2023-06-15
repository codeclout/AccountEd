package server

import (
	"regexp"
	"strconv"
)

// ValidateRequestLimit checks if the given limit string is a valid request limit by verifying that it contains one or two digits. It returns
// true if the limit string matches the corresponding regular expression and there's no error, otherwise returns false.
func ValidateRequestLimit(limit *string) bool {
	if ok, e := regexp.MatchString(`(?m)^[\d]{1,2}$`, *limit); ok && e == nil {
		return true
	}

	return false
}

func getRequestLimit(q *string) *int16 {
	var limit int16 = -1

	if ok := ValidateRequestLimit(q); ok {
		n, _ := strconv.ParseInt(*q, 10, 16)
		limit = int16(n)
	}

	return &limit
}
