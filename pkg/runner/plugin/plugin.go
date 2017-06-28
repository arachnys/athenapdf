package plugin

//go:generate go-bindata -pkg $GOPACKAGE js

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func List() []string {
	plugins := make([]string, 0, len(AssetNames()))
	for _, assetPath := range AssetNames() {
		plugin := strings.Split(path.Base(path.Clean(assetPath)), ".")[0]
		plugins = append(plugins, plugin)
	}
	return plugins
}

func Default() []string {
	return []string{"no-css-media"}
}

func Get(builtin []string, custom []string) ([]io.Reader, error) {
	br, err := GetBuiltin(builtin...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	cr, err := GetCustom(custom...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return append(br, cr...), nil
}

func GetBuiltin(pluginsToLoad ...string) ([]io.Reader, error) {
	loaded := make([]io.Reader, 0, len(pluginsToLoad))

	for _, pluginName := range pluginsToLoad {
		var buf bytes.Buffer

		plugin, err := GetBuiltinRaw(pluginName)
		if err != nil {
			return nil, err
		}
		if _, err := buf.Write(plugin); err != nil {
			return nil, errors.WithStack(err)
		}

		loaded = append(loaded, &buf)
	}

	return loaded, nil
}

func GetBuiltinRaw(pluginName string) ([]byte, error) {
	if !Exists(pluginName) {
		return nil, errors.Errorf(
			"plugin `%s` does not exist, available plugins: %s",
			pluginName,
			strings.Join(List(), ", "),
		)
	}

	b, err := Asset(fmt.Sprintf("js/%s.js", pluginName))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return b, nil
}

func GetCustom(pluginsToLoad ...string) ([]io.Reader, error) {
	loaded := make([]io.Reader, 0, len(pluginsToLoad))

	for _, pluginPath := range pluginsToLoad {
		var buf bytes.Buffer

		p, err := filepath.Abs(path.Clean(pluginPath))
		if err != nil {
			return nil, errors.WithStack(err)
		}

		f, err := os.Open(p)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if f != nil {
			defer f.Close()
		}

		if _, err := io.Copy(&buf, f); err != nil {
			return nil, errors.WithStack(err)
		}

		loaded = append(loaded, &buf)
	}

	return loaded, nil
}

func Exists(pluginName string) bool {
	for _, availablePlugin := range List() {
		if pluginName == availablePlugin {
			return true
		}
	}

	return false
}
