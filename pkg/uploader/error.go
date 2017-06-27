package uploader

import (
	"fmt"
)

type UploaderError struct {
	err          error
	uploaderName string
}

func (e UploaderError) Error() string {
	if e.uploaderName != "" {
		return fmt.Sprintf("%s: %s", e.uploaderName, e.err.Error())
	}
	return e.err.Error()
}

func (e UploaderError) Cause() error {
	return e.err
}
