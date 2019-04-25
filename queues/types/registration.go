package qtypes

//go:generate easyjson

import (
	"fmt"

	"github.com/studtool/common/errs"
)

//easyjson:json
type RegistrationEmailData map[string]interface{}

func (d *RegistrationEmailData) Validate() *errs.Error {
	fields := []string{"email", "token"}

	if len(*d) != len(fields) {
		return d.makeErr()
	}

	for _, f := range fields {
		if _, ok := (*d)[f]; !ok {
			return d.makeErr()
		}
	}

	return nil
}

func (d *RegistrationEmailData) makeErr() *errs.Error {
	return errs.NewInternalError(fmt.Sprintf("wrong email data format: %v", *d))
}
