package http

import "strconv"

var limit int64

func (a *Adapter) getRequestLimit(q string) int64 {
	if len(q) > 0 {
		n, e := strconv.ParseInt(q, 10, 16)
		if e != nil {
			limit = -1
		}

		limit = n
	}

	return limit
}
