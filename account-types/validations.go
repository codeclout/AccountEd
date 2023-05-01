package accountTypes

import (
	"errors"
	"fmt"
	"regexp"
)

func ValidateAccountTypeId(id AccountTypeIn) (bool, error) {

	if ok, e := regexp.MatchString(`(?m)^[a-zA-Z0-9]{24}$`, id.Id); ok && e == nil {
		return true, nil
	}

	return false, errors.New(fmt.Sprintf("validation failed for %T", id))
}
