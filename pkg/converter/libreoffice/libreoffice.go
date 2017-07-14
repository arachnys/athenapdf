package libreoffice

import (
	"bytes"
	"context"
	"github.com/pkg/errors"
	"io"
	"os"
	"os/exec"
	"regexp"

	"github.com/arachnys/athenapdf/pkg/converter"
	"github.com/arachnys/athenapdf/pkg/proto"
	"github.com/arachnys/athenapdf/pkg/uri"
)

const (
	converterName = "libreoffice"
)

var defaultFlags = []string{
	"--headless",
}

type LibreOffice struct{}

func init() {
	converter.Register(converterName, &LibreOffice{})
}

func (*LibreOffice) String() string {
	return converterName
}

func (*LibreOffice) Convert(ctx context.Context, req *proto.Conversion, opts map[string]*proto.Option) (io.Reader, error) {
	if !uri.IsLocal(req.GetUri()) {
		return nil, errors.New("conversion uri is not local")
	}

	uri, err := uri.RemoveScheme(req.GetUri())
	if err != nil {
		return nil, err
	}

	tmpDir := os.TempDir()
	flags := append(defaultFlags, []string{
		"--convert-to",
		"pdf",
		"--outdir",
		tmpDir,
		uri,
	}...)
	cmd := exec.CommandContext(ctx, libreOfficePath(), flags...)
	out, err := cmd.Output()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	re := regexp.MustCompile(`\-\> (.+) using filter`)
	matches := re.FindStringSubmatch(string(out))
	if len(matches) != 2 {
		return nil, errors.Errorf("failed to identify pdf output: `%s`", string(out))
	}
	outputFilePath := matches[1]
	if _, err := os.Stat(outputFilePath); os.IsNotExist(err) {
		return nil, errors.Wrap(err, "pdf output does not exist")
	}
	defer os.Remove(outputFilePath)

	var buf bytes.Buffer
	f, err := os.Open(outputFilePath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if f != nil {
		defer f.Close()
	}

	if _, err := io.Copy(&buf, f); err != nil {
		return nil, errors.WithStack(err)
	}

	return &buf, nil
}

func (*LibreOffice) SupportedMimeTypes() []string {
	return []string{
		"application/msword",
		"application/rtf",
		"vnd.openxmlformats-officedocument",
	}
}
