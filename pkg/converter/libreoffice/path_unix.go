// +build linux freebsd netbsd openbsd

package libreoffice

import "os/exec"

const (
	DefaultLibreOfficePath = "/usr/lib/openoffice/program/soffice"
)

var libreOfficeExecutables = []string{
	"soffice",
	"soffice.bin",
	"libreoffice",
}

func libreOfficePath() string {
	for _, p := range libreOfficeExecutables {
		path, err := exec.LookPath(p)
		if err == nil {
			return path
		}
	}

	return DefaultLibreOfficePath
}
