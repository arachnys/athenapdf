# athenapdf

[![](https://badge.imagelayers.io/arachnysdocker/athenapdf:latest.svg)](https://imagelayers.io/?images=arachnysdocker/athenapdf:latest 'Get your own badge on imagelayers.io')

**The CLI component of athenapdf.**

A simple CLI tool to convert HTML to PDF from a local file or a URL to a web page using [Electron][electron] ([Chromium][chromium]).

It can be used anywhere with [Docker][docker] installed, even in a headless environment (i.e. no display server).


## Features

- Alternative / drop-in replacement for [wkhtmltopdf]
- Uses the latest, stable release of [Chromium][chromium], and its Blink web rendering engine:
    - Supports the latest HTML5, and CSS3
    - Supports JavaScript, and AJAX
    - Supports internal linking / page anchors within a PDF
    - More: https://www.chromestatus.com/features
- Supports local HTML files, and remote URLs
- Supports writing to standard output: you can take advantage of UNIX piping / redirects
- Supports rendering in:
    - Landscape or portrait
    - A3, A4, A5, Legal, Letter, and Tabloid page size
- Adjustable PDF generation delay
- Adjustable, built-in timeout mechanism
- Adjustable cache control
- Adjustable browser zoom settings
- Adjustable margin sizes
- Automatically falls back to `screen` stylesheets if no `print` stylesheet is defined
- [Aggressive mode](docs/aggressive.md): declutter web pages, and improves readability
- Bypass paywalls for most digital publications with a single `-B` flag (experimental feature)
- Dockerized:
    - Easy to set up, distribute, and run
    - Runs in [headless] mode (the [display server][xvfb] is handled for you)
    - Out-of-the-box support for a broad range of foreign characters
- Actively maintained, and production tested


## Quick Start

See [`docs/quick-start.md`](docs/quick-start.md).


## Development

See [`docs/development.md`](docs/development.md).


## Building

See [`docs/building.md`](docs/building.md).



[docker]: https://www.docker.com/
[electron]: http://electron.atom.io/
[chromium]: https://www.chromium.org/
[wkhtmltopdf]: http://wkhtmltopdf.org/
[headless]: http://internetofthingsagenda.techtarget.com/definition/headless-system
[xvfb]: https://en.wikipedia.org/wiki/Xvfb
