package converter

// Conversion represents the most basic, skeleton conversion job.
// Currently, it is assumed that the input is a HTML (or similar) resource,
// and the output is a PDF.
// Conversion implements the Converter interface with dummy methods.
// Any child conversion type should overwrite these methods.
type Conversion struct{}

// Convert should return a byte slice containing the converted resource,
// synchronously.
// It should terminate any long-running processes (and Goroutines) if the done
// channel is returned.
func (c Conversion) Convert(s ConversionSource, done <-chan struct{}) ([]byte, error) {
	return []byte{}, nil
}

// Upload should take a byte slice, and return a boolean indicating if the data
// was used for post-processing, e.g. uploading to a remote host like S3.
// It should always return false if there is an error.
func (c Conversion) Upload(b []byte) (bool, error) {
	return false, nil
}
