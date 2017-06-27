package process

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"

	"github.com/arachnys/athenapdf/pkg/converter"
	"github.com/arachnys/athenapdf/pkg/fetcher"
	"github.com/arachnys/athenapdf/pkg/mime"
	"github.com/arachnys/athenapdf/pkg/proto"
	"github.com/arachnys/athenapdf/pkg/uploader"
)

const (
	defaultMimeType = "text/plain; charset=\"UTF-8\""
)

func Process(ctx context.Context, p *proto.Process) (io.Reader, bool, error) {
	var converterFn converter.ConverterFunc
	var err error

	conversion := p.GetConversion()
	converterName := p.GetConverter()

	if converter.IsLocal(conversion) {
		// Clean up local conversion
		defer os.Remove(conversion.GetUri())
	}

	if conversion.GetMimeType() == "" {
		if !converter.IsLocal(conversion) {
			return nil, false, errors.New("process: mime type must be provided for non-local conversions")
		}

		mimeType, err := mime.TypeFromFile(conversion.GetUri())
		if err != nil {
			return nil, false, err
		}
		if mimeType == "" {
			mimeType = defaultMimeType
		}

		conversion.MimeType = mimeType
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

	if converterName == "" {
		c, err := converter.GetFromConversionExcluding(conversion)([]string{})
		if err != nil {
			return nil, false, err
		}
		converterFn = c.Convert
	} else {
		converterFn = converter.Convert(converterName)
	}

	cr, err := converterFn(ctx, p.GetConversion(), p.GetConverterOptions())
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
