package http

import (
	"strconv"
)

func (a *Adapter) getRequestLimit(q *string) *int16 {
	var limit int16 = -1

	if len(*q) > 0 {
		n, e := strconv.ParseInt(*q, 10, 16)
		if e != nil {
			return &limit
		}

		limit = int16(n)
	}

	return &limit
}
