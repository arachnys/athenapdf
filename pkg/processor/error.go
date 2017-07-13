package processor

type ProcessError struct {
	err error
}

func (e ProcessError) Error() string {
	return e.err.Error()
}

func (e ProcessError) Cause() error {
	return e.err
}
