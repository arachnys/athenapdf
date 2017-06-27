package process

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/arachnys/athenapdf/pkg/converter"
	"github.com/arachnys/athenapdf/pkg/fetcher"
	"github.com/arachnys/athenapdf/pkg/proto"
	"github.com/arachnys/athenapdf/pkg/uploader"
)

func Process(ctx context.Context, p *proto.Process) (io.Reader, bool, error) {
	conversion := p.GetConversion()

	if converter.IsLocal(conversion) {
		// Clean up local conversion
		defer os.Remove(conversion.GetUri())
	}

	if p.GetFetcher() != "" {
		fr, mimeType, err := fetcher.Fetch(p.GetFetcher())(ctx, conversion.GetUri(), p.GetFetcherOptions())
		if err != nil {
			return nil, false, err
		}

		tmpFile, err := ioutil.TempFile("", "athenapdf-process")
		if err != nil {
			return nil, false, err
		}
		if tmpFile.Name() != "" {
			defer os.Remove(tmpFile.Name())
		}

		if _, err := io.Copy(tmpFile, fr); err != nil {
			return nil, false, err
		}

		// TODO: store the original URI
		p.Conversion.Uri = fmt.Sprintf("file://%s", tmpFile.Name())
		if conversion.GetMimeType() == "" {
			p.Conversion.MimeType = mimeType
		}
	}

	cr, err := converter.Convert(p.GetConverter())(ctx, p.GetConversion(), p.GetConverterOptions())
	if err != nil {
		return nil, false, err
	}

	if p.GetUploader() != "" {
		if err := uploader.Upload(p.GetUploader())(ctx, cr, p.GetUploaderOptions()); err != nil {
			return nil, false, err
		}
		return nil, true, nil
	}

	return cr, false, nil
}
