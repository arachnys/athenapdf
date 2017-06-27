package converter

import (
	"fmt"
)

type ConverterError struct {
	err           error
	converterName string
}

func (e ConverterError) Error() string {
	if e.converterName != "" {
		return fmt.Sprintf("%s: %s", e.converterName, e.err.Error())
	}
	return e.err.Error()
}

func (e ConverterError) Cause() error {
	return e.err
}
