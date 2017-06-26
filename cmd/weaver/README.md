# weaver

[![](https://badge.imagelayers.io/arachnysdocker/athenapdf-service:latest.svg)](https://imagelayers.io/?images=arachnysdocker/athenapdf-service:latest 'Get your own badge on imagelayers.io')
[![GoDoc](https://godoc.org/github.com/arachnys/athenapdf/weaver?status.svg)](https://godoc.org/github.com/arachnys/athenapdf/weaver)

**The microservice component of athenapdf.**

A scalable, Go microservice for running HTML to PDF conversions using [`athenapdf`][athenapdf].

Although it was predominantly designed for generating a PDF from a HTML using [`athenapdf`][athenapdf], it is agnostic about the converter it uses. i.e. You can build an adapter for other CLI or web-based converters.


## Features

- Extensible converter backend:
    - [`athenapdf`][athenapdf]
    - [CloudConvert][cloudconvert]
- Hosts blocking:
    - Blocks unwanted ads, and trackers
    - Speeds up PDF generation
- Supports uploading conversions to S3
- Supports returning conversions to the browser (`application/pdf`)
- Concurrent workers, and internal job queue:
    - Stateless
    - Easy to scale horizontally, and vertically
- Strong service visibility for quality control:
    - Metrics collection ([statsd])
    - Error logging ([Sentry][sentry])
- Dockerized:
    - Easy to set up, distribute, and deploy
    - Runs in headless mode (the display server is handled for you)
    - Out-of-the-box support for a broad range of foreign characters
- Actively maintained, and production tested


## Quick Start

See [`docs/quick-start.md`](docs/quick-start.md).


## Development

See [`docs/development.md`](docs/development.md).


## Building

See [`docs/building.md`](docs/building.md).


[athenapdf]: ../cli
[cloudconvert]: https://cloudconvert.com/
[statsd]: https://github.com/etsy/statsd
[sentry]: https://getsentry.com/
