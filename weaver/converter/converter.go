package converter

type Converter interface {
	Convert(<-chan struct{}) ([]byte, error)
	Upload([]byte) (bool, error)
}
