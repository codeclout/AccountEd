package http

import (
	"strconv"

	"github.com/codeclout/AccountEd/onboarding/internal"
)

func (a *Adapter) getRequestLimit(q *string) *int16 {
	var limit int16 = -1

	if ok := internal.ValidateRequestLimit(q); ok {
		n, _ := strconv.ParseInt(*q, 10, 16)
		limit = int16(n)
	}

	return &limit
}
