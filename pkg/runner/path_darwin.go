// +build darwin
// Taken from: https://github.com/knq/chromedp/blob/master/runner/path_darwin.go

package runner

const (
	DefaultChromePath = `/Applications/Google Chrome Canary.app/Contents/MacOS/Google Chrome Canary`
)

func getChromePath() string {
	return DefaultChromePath
}
