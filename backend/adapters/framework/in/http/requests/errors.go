package requests

import "regexp"

var (
	ErrorDuplicateKey            = []byte("duplicate")
	ErrorFailedAction            = []byte("requested action failed")
	ErrorFailedRequestValidation = []byte("payload received failed to validate")
	ErrorInvalidJSON             = []byte("invalid json")
)

type RequestErrorWithRetry struct {
	Msg         string `json:"msg"`
	ShouldRetry bool   `json:"shouldRetry"`
}

func ShouldRetryRequest(s int) bool {
	b, e := regexp.Match(`429|504`, []byte(string(rune(s))))

	if e != nil {
		return false
	}

	return b
}
