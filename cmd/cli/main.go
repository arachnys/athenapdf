package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"time"

	"github.com/arachnys/athenapdf/pkg/proto"
	"github.com/arachnys/athenapdf/pkg/runner"
	"github.com/arachnys/athenapdf/pkg/runner/plugin"
)

const (
	appName        = "athenapdf"
	appVersion     = "3.0.0-b"
	appDescription = "convert (M)HTML to PDF using headless Chromium / Blink (DevTools Protocol)"

	defaultPageHeight = "11"
	defaultPageWidth  = "8.5"
	defaultPageMargin = "0.4"
)

var (
	defaultUserAgent = fmt.Sprintf("%s/%s", appName, appVersion)

	app = kingpin.New(appName, appDescription).Version(appVersion)

	debug    = app.Flag("debug", "enable debug mode / verbose logging").Short('D').Bool()
	dryRun   = app.Flag("dry-run", "do not render the PDF output, useful for testing").Bool()
	timeout  = app.Flag("timeout", "duration to wait for the page to load before exiting").Default("60s").Duration()
	insecure = app.Flag("insecure", "allow running insecure content / bypass certificate errors").Bool()

	server = app.Flag("server", "run in client-server mode by connecting to a local running instance of Chromium's Remote Debugging Protocol").URL()
	proxy  = app.Flag("proxy", "use a proxy server for HTTP(S) requests (only works in non-client-server mode, default)").URL()

	builtinPlugins = app.Flag("plugin", "built-in JavaScript plugin to execute on page load (repeatable)").Default(plugin.Default()...).Enums(plugin.List()...)
	customScripts  = app.Flag("run-script", "custom JavaScript file to execute on page load (repeatable)").ExistingFiles()

	marginBottom = app.Flag("margin-bottom", "bottom margin of PDF in inches").Short('B').Default(defaultPageMargin).Float64()
	marginLeft   = app.Flag("margin-left", "left margin of PDF in inches").Short('L').Default(defaultPageMargin).Float64()
	marginRight  = app.Flag("margin-right", "right margin of PDF in inches").Short('R').Default(defaultPageMargin).Float64()
	marginTop    = app.Flag("margin-top", "top margin of PDF in inches").Short('T').Default(defaultPageMargin).Float64()
	pageHeight   = app.Flag("page-height", "height of PDF in inches").Short('H').Default(defaultPageHeight).Float64()
	pageWidth    = app.Flag("page-width", "width of PDF in inches").Short('W').Default(defaultPageWidth).Float64()
	orientation  = app.Flag("orientation", "orientation of PDF").Short('O').Default(proto.Conversion_Dimensions_PORTRAIT.String()).Enum(proto.Conversion_Dimensions_PORTRAIT.String(), proto.Conversion_Dimensions_LANDSCAPE.String())
	scale        = app.Flag("scale", "scale of PDF rendering").Short('S').Default("1").Float64()

	headers   = app.Flag("custom-header", "set an additional HTTP header (repeatable)").PlaceHolder("NAME=VALUE").StringMap()
	cookies   = app.Flag("cookie", "set an additional cookie, the value must be URL encoded (repeatable)").PlaceHolder("NAME=VALUE").StringMap()
	userAgent = app.Flag("user-agent", "set user agent").Short('U').Default(defaultUserAgent).String()
	mediaType = app.Flag("media-type", "set media type to emulate").Short('M').Default("print").String()

	noCache      = app.Flag("no-cache", "do not use cache for any request").Bool()
	noJavaScript = app.Flag("no-javascript", "do not execute JavaScript").Bool()
	noBackground = app.Flag("no-background", "do not print background graphics").Bool()

	input  = app.Arg("input", "URL / file name of target to convert to PDF (use `-` if piping via stdin)").Required().String()
	output = app.Arg("output", "path to store conversion output (use `-` if piping to stdout)").String()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	if *debug || *dryRun {
		startTime := time.Now()
		defer func() {
			fmt.Printf("\nProcessing time: %s\n", time.Since(startTime).String())
		}()
	}

	// Create new runner
	r := newRunner()

	exit, err := r.AutoTarget()
	defer func() {
		kingpin.FatalIfError(exit(), "")
	}()
	if err != nil {
		kingpin.Errorf("%s", err)
		return
	}

	*input, err = handleInput(*input)
	if err != nil {
		kingpin.Errorf("%s", err)
		return
	}

	// Create new conversion request
	req := newConversion()

	// Set runner specific / supported options
	req.Options = getOptions()

	// Set custom headers
	req.Headers = getHeaders()

	// Set custom cookies
	req.Cookies = getCookies()

	// Run conversion
	b, err := r.Convert(req)
	if err != nil {
		kingpin.Errorf("%s", err)
		return
	}

	if err := handleOutput(b, *output, *dryRun); err != nil {
		kingpin.Errorf("%s", err)
		return
	}
}

func newRunner() *runner.Runner {
	return &runner.Runner{
		Debug:   *debug,
		DryRun:  *dryRun,
		Timeout: *timeout,
		Server:  *server,
		Proxy:   *proxy,
		Plugins: struct {
			Builtin []string
			Custom  []string
		}{
			*builtinPlugins,
			*customScripts,
		},
	}
}

func newConversion() *proto.Conversion {
	return &proto.Conversion{
		Uri: *input,
		Dimensions: &proto.Conversion_Dimensions{
			*marginBottom,
			*marginLeft,
			*marginRight,
			*marginTop,
			*pageHeight,
			*pageWidth,
			proto.Conversion_Dimensions_Orientation(
				proto.Conversion_Dimensions_Orientation_value[*orientation],
			),
		},
	}
}

func getOptions() map[string]*proto.Option {
	options := make(map[string]*proto.Option, 6)
	options["scale"] = &proto.Option{
		&proto.Option_DoubleValue{
			*scale,
		},
	}
	options["user_agent"] = &proto.Option{
		&proto.Option_StringValue{
			*userAgent,
		},
	}
	options["media_type"] = &proto.Option{
		&proto.Option_StringValue{
			*mediaType,
		},
	}
	options["insecure"] = &proto.Option{
		&proto.Option_BoolValue{
			*insecure,
		},
	}
	options["no_cache"] = &proto.Option{
		&proto.Option_BoolValue{
			*noCache,
		},
	}
	options["no_javascript"] = &proto.Option{
		&proto.Option_BoolValue{
			*noJavaScript,
		},
	}
	options["no_background"] = &proto.Option{
		&proto.Option_BoolValue{
			*noBackground,
		},
	}

	return options
}

func getHeaders() []*proto.Header {
	protoHeaders := make([]*proto.Header, len(*headers))
	for name, value := range *headers {
		protoHeaders = append(
			protoHeaders,
			&proto.Header{
				Name:  name,
				Value: value,
			},
		)
	}

	return protoHeaders
}

func getCookies() []*proto.Cookie {
	protoCookies := make([]*proto.Cookie, len(*cookies))
	for name, value := range *cookies {
		protoCookies = append(
			protoCookies,
			&proto.Cookie{
				Name:  name,
				Value: value,
				Url:   *input,
			},
		)
	}

	return protoCookies
}
