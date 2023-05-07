package server

import (
  "regexp"
  "strconv"
)

func ValidateRequestLimit(limit *string) bool {
  if ok, e := regexp.MatchString(`(?m)^[1-9]{1,3}$`, *limit); ok && e == nil {
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
