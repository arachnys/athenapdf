package main

import (
	"os"
	"strconv"
)

// CloudConvert configuration.
type CloudConvert struct {
	// API key to access CloudConvert's API.
	// It can be found in https://cloudconvert.com/user/profile
	APIKey string
	// Base API URL (without backslashes).
	// Defaults to https://api.cloudconvert.com
	APIUrl string
}

// Statsd configuration.
// It contains a HOST:PORT address to a statsd server, and a prefix (namespace)
// for the recorded stats.
type Statsd struct {
	Address string
	Prefix  string
}

// Config for Weaver.
// It contains all the configuration variables that will be used by the
// microservice.
type Config struct {
	CloudConvert
	// Defaults to none.
	Statsd
	// The address:port for the HTTP server to listen on.
	// Defaults to ':8080'
	HTTPAddr string
	// The address:port for the HTTPS server to listen on.
	// Defaults to none
	HTTPSAddr string
	// The TLS certificate to use if the HTTPS listener is enabled.
	// Defaults to none
	TLSCertFile string
	// The TLS key to use if the HTTPS listener is enabled.
	// Defaults to none
	TLSKeyFile string
	// Authentication key to be used with AuthorizationMiddleware.
	// It will be used to protect all conversion routes.
	// Defaults to 'arachnys-weaver'.
	AuthKey string
	// See AthenaPDF CMD.
	// Defaults to 'athenapdf -S'.
	AthenaCMD string
	// The maximum number of workers / concurrent conversions that can be
	// running at any one time.
	// Defaults to 10.
	MaxWorkers int
	// The maximum number of conversion jobs that can be held in a queue at
	// any one time without blocking a NewWork Goroutine.
	// Defaults to 50.
	MaxConversionQueue int
	// Seconds until a conversion job is terminated, and a handler is returned.
	// Defaults to 90.
	WorkerTimeout int
	// Toggles falling back to CloudConvert if athenapdf CLI fails to convert.
	// The failure may also be due to a timeout.
	// Defaults to false.
	ConversionFallback bool
	// The data source name (DSN) for a Sentry server (used for logging errors).
	// Defaults to none.
	SentryDSN string
}

// NewEnvConfig initialises configuration variables from the environment.
func NewEnvConfig() Config {
	// Set defaults
	cloudconvert := CloudConvert{APIUrl: "https://api.cloudconvert.com"}
	conf := Config{
		CloudConvert:       cloudconvert,
		HTTPAddr:           ":8080",
		AuthKey:            "arachnys-weaver",
		AthenaCMD:          "athenapdf -S",
		MaxWorkers:         10,
		MaxConversionQueue: 50,
		WorkerTimeout:      90,
		ConversionFallback: false,
	}

	if httpAddr := os.Getenv("WEAVER_HTTP_ADDR"); httpAddr != "" {
		conf.HTTPAddr = httpAddr
	}

	if httpsAddr := os.Getenv("WEAVER_HTTPS_ADDR"); httpsAddr != "" {
		conf.HTTPSAddr = httpsAddr
	}

	if tlsCertFile := os.Getenv("WEAVER_TLS_CERT_FILE"); tlsCertFile != "" {
		conf.TLSCertFile = tlsCertFile
	}

	if tlsKeyFile := os.Getenv("WEAVER_TLS_KEY_FILE"); tlsKeyFile != "" {
		conf.TLSKeyFile = tlsKeyFile
	}

	if authKey := os.Getenv("WEAVER_AUTH_KEY"); authKey != "" {
		conf.AuthKey = authKey
	}

	if athenaCMD := os.Getenv("WEAVER_ATHENA_CMD"); athenaCMD != "" {
		conf.AthenaCMD = athenaCMD
	}

	// NOTE: we aren't handle the _unlikely_ event of errors properly (they are being suppressed)
	if maxWorkers := os.Getenv("WEAVER_MAX_WORKERS"); maxWorkers != "" {
		conf.MaxWorkers, _ = strconv.Atoi(maxWorkers)
	}

	if maxConversionQueue := os.Getenv("WEAVER_MAX_CONVERSION_QUEUE"); maxConversionQueue != "" {
		conf.MaxConversionQueue, _ = strconv.Atoi(maxConversionQueue)
	}

	if workerTimeout := os.Getenv("WEAVER_WORKER_TIMEOUT"); workerTimeout != "" {
		conf.WorkerTimeout, _ = strconv.Atoi(workerTimeout)
	}

	if conversionFallback := os.Getenv("WEAVER_CONVERSION_FALLBACK"); conversionFallback != "" {
		conf.ConversionFallback, _ = strconv.ParseBool(conversionFallback)
	}

	if cloudConvertAPI := os.Getenv("CLOUDCONVERT_API"); cloudConvertAPI != "" {
		conf.CloudConvert.APIUrl = cloudConvertAPI
	}

	if cloudConvertKey := os.Getenv("CLOUDCONVERT_KEY"); cloudConvertKey != "" {
		conf.CloudConvert.APIKey = cloudConvertKey
	}

	if statsdAddress := os.Getenv("STATSD_ADDRESS"); statsdAddress != "" {
		conf.Statsd.Address = statsdAddress
	}

	if statsdPrefix := os.Getenv("STATSD_PREFIX"); statsdPrefix != "" {
		conf.Statsd.Prefix = statsdPrefix
	}

	if sentryDSN := os.Getenv("SENTRY_DSN"); sentryDSN != "" {
		conf.SentryDSN = sentryDSN
	}

	return conf
}
