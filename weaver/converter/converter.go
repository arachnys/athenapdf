package converter

type Converter interface {
	Convert(ConversionSource, <-chan struct{}) ([]byte, error)
	Upload([]byte) (bool, error)
}
