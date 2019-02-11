# Athena

[![Build Status](https://travis-ci.org/arachnys/athenapdf.svg?branch=master)](https://travis-ci.org/arachnys/athenapdf)
[![License](http://img.shields.io/badge/license-MIT-orange.svg?style=flat)](http://opensource.org/licenses/MIT)
[![Gitter chat](https://badges.gitter.im/arachnys/athenapdf.png)](https://gitter.im/arachnys/athenapdf)

[**Live Demo / Screenshots**][athenapdf]


Simple, Docker-powered PDF conversions.

Athena is comprised of an [Electron][electron] command line interface (CLI) tool, and a [Go][go] microservice for converting HTML to PDF documents.

> Athena transformed Arachne into a spider for challenging her as a weaver and/or weaving a tapestry that insulted the gods.

**Examples:**

- Original: [Google isn’t even close as a tool for proper due diligence. Why not?][example-1] **(Converted: [PDF][example-1-pdf] | [Aggressive][example-1-aggressive])**
- Original: [Panamanian Law Firm Is Gatekeeper To Vast Flow of Murky Offshore Secrets][example-2] **(Converted: [PDF][example-2-pdf] | [Aggressive][example-2-aggressive])**

_When [aggressive mode][aggressive] is enabled, only the essential contents of a page are kept in the generated PDF document. It is a clutter-free version of the web page, perfect for reading._


## Background

Athena is an open source project.

It was designed to [do one thing and to do it well][unixphil] - PDF conversions; to work together with other programs; and to be able to handle text streams, because that is a universal interface.

It aims to give users an on-demand capability to convert HTML to PDF without frills.

At the lowest level, its [CLI][cli] component ([`athenapdf`][cli]) was designed to be an alternative / drop-in replacement for [wkhtmltopdf], a popular CLI tool for HTML to PDF conversions. Because of Docker the CLI syntax is a bit more complex but it's much more reliable.

(For what it's worth, wkhtmltopdf is great, but it has a horrible habit of crashing unexpectedly - especially when printing documents with invalid HTML, problematic CSS or other issues).

There is also a [microservice][weaver] component ([`weaver`][weaver]), allowing you to leverage Athena over HTTP.

## Getting Started

### CLI vs Microservice

Our [CLI][cli] tool will suffice for most simple, and everyday HTML to PDF conversions.

However, for conversions at scale / PDF conversion as a service, we recommend getting started with our [microservice][weaver] component instead.

![CLI vs Microservice](https://s3-eu-west-1.amazonaws.com/athena-pdf-public/examples/docker.png)

The microservice is packaged with `athenapdf`, and you can run both components independently.

### Docker

Both components are packaged, and distributed as [Docker][docker] images.

The only dependency you will need is Docker, and the rest will be handled for you (even if you are running in an environment without a display server - [headless environment][headless]).

### Quick Start

Before starting, ensure your [Docker][docker] environment is set up, and ready-to-use.

**For OSX / Windows users**, ensure your [Docker Machine][docker-machine] is prepared, and the appropriate environment variables are established.

#### CLI

[![asciicast](https://asciinema.org/a/c1fbvtdnrctfq6baipfox00ct.png)](https://asciinema.org/a/c1fbvtdnrctfq6baipfox00ct)

1. `docker pull arachnysdocker/athenapdf`
2. `docker run --rm -v $(pwd):/converted/ arachnysdocker/athenapdf athenapdf <input_path> [output_path]`
3. See [`cli`][cli] for full documentation

The `[output_path]` can be omitted.

Example: `docker run --rm -v $(pwd):/converted/ arachnysdocker/athenapdf athenapdf http://blog.arachnys.com/`

**For Windows users**, an additional forward slash must precede the volume when using Git Bash / MinGW:

```bash
docker run --rm -v /$(pwd):/converted/ arachnysdocker/athenapdf athenapdf http://blog.arachnys.com/
```

Alternatively, if using the Windows command prompt:

```cmd
docker run --rm -v %cd%:/converted/ arachnysdocker/athenapdf athenapdf http://blog.arachnys.com/
```

#### Microservice

[![asciicast](https://asciinema.org/a/41247.png)](https://asciinema.org/a/41247)

1. `docker pull arachnysdocker/athenapdf-service`
2. `docker run -p 8080:8080 --rm arachnysdocker/athenapdf-service`
3. Inline conversion: `http://<docker-address>:8080/convert?auth=arachnys-weaver&url=http://blog.arachnys.com/`
4. OR cURL, and redirect output to file: `curl http://dockerhost:8080/convert\?auth\=arachnys-weaver\&url\=http://blog.arachnys.com/ |> out.pdf`
5. See [`weaver`][weaver] for full documentation

The default authentication key is `arachnys-weaver`. This can be changed through the `WEAVER_AUTH_KEY` environment variable.

The microservice can be deployed scalably to [ECS][ecs] if you want to build your own conversion farm.

## License

Please note `athenapdf` is NEITHER affiliated with NOR endorsed by Google Inc. and GitHub Inc.

See [`LICENSE`](LICENSE).

---

![Arachnys](https://s3-eu-west-1.amazonaws.com/athena-pdf-public/examples/logo.png)

An [Arachnys][arachnys] Christmas project.


[athenapdf]: http://www.athenapdf.com/
[unixphil]: https://en.wikipedia.org/wiki/Unix_philosophy
[example-1]: http://blog.arachnys.com/google-isnt-even-close-to-proper-due-diligence.-why-not?utm_campaign=athena&utm_medium=external%20website&utm_source=github&utm_content=readme
[example-1-pdf]: https://s3-eu-west-1.amazonaws.com/athena-pdf-public/examples/d0ae9e636bb7edceac48bd9b0780ad9dd69397cc.pdf
[example-1-aggressive]: https://s3-eu-west-1.amazonaws.com/athena-pdf-public/examples/d0ae9e636bb7edceac48bd9b0780ad9dd69397cc-aggressive.pdf
[example-2]: https://panamapapers.icij.org/20160403-mossack-fonseca-offshore-secrets.html
[example-2-pdf]: https://s3-eu-west-1.amazonaws.com/athena-pdf-public/examples/f86e19ea13661e47c1f5f83bbde650eb3fe85aad.pdf
[example-2-aggressive]: https://s3-eu-west-1.amazonaws.com/athena-pdf-public/examples/f86e19ea13661e47c1f5f83bbde650eb3fe85aad-aggressive.pdf
[aggressive]: cli/docs/aggressive.md
[electron]: http://electron.atom.io/
[go]: https://golang.org/
[cli]: cli/
[weaver]: weaver/
[wkhtmltopdf]: http://wkhtmltopdf.org/
[cloudconvert]: https://cloudconvert.com/
[docker]: https://www.docker.com/
[docker-machine]: https://docs.docker.com/mac/step_one/
[headless]: http://internetofthingsagenda.techtarget.com/definition/headless-system
[arachnys]: https://www.arachnys.com/?utm_campaign=athena&utm_medium=external%20website&utm_source=github&utm_content=readme
[ecs]: https://aws.amazon.com/ecs/details/
