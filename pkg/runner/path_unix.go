// +build linux freebsd netbsd openbsd
// Taken from: https://github.com/knq/chromedp/blob/master/runner/path_unix.go

package runner

import "os/exec"

const (
	DefaultChromePath = "/usr/bin/google-chrome"
)

var chromeExecutables = []string{
	"chrome",
	"chromium",
	"chromium-browser",
	"google-chrome",
	"google-chrome-beta",
	"google-chrome-unstable",
}

func getChromePath() string {
	for _, p := range chromeExecutables {
		path, err := exec.LookPath(p)
		if err == nil {
			return path
		}
	}

	return DefaultChromePath
}
