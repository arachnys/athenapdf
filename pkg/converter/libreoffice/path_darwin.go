// +build darwin

package libreoffice

const (
	DefaultLibreOfficePath = `/Applications/LibreOffice.app/Contents/MacOS/soffice`
)

func libreOfficePath() string {
	return DefaultLibreOfficePath
}
