# Quick Start


## Install

Currently, `athenapdf` is packaged, and distributed as a [Docker][docker] image.

Before starting, ensure your Docker environment is [set up][docker], and ready-to-use. **For OSX / Windows users**, ensure your [Docker Machine][docker-machine] is prepared, and the appropriate environment variables are established.

It can be installed using:

```bash
docker pull arachnysdocker/athenapdf
```

You can run the same command above to get updates.


## Usage

### Syntax

```bash
docker run --rm -v $(pwd):/converted/ arachnysdocker/athenapdf athenapdf <input_path> [output_path]
```

The `<input_path>` can either be a local HTML file or a URL to a web page.

**The `[output_path]` is optional.** If it is not specified, the output path is set to the current working directory with a SHA1 hash of the file / URL as its file name.

Due to a recent change in [Docker v1.10.0][1.10.0], you will have to add `--security-opt seccomp:unconfined` after `docker run` to suppress the `libudev: udev_has_devtmpfs: name_to_handle_at on /dev: Operation not permitted` error which may prevent you from successfully generating a PDF.

#### Windows Usage

For Windows users, binding a volume with the preceding command may cause an error when using Git Bash / MinGW. In that case, adding an additional forward slash before the volume will remove the error:

```bash
docker run --rm -v /$(pwd):/converted/ arachnysdocker/athenapdf athenapdf <input_path> [output_path]
```

Alternatively, if using the Windows command prompt, `$(pwd)` must be replaced by `%cd%`:

```cmd
docker run --rm -v %cd%:/converted/ arachnysdocker/athenapdf athenapdf <input_path> [output_path]
```

### Examples

#### Local file

```bash
docker run --rm -v $(pwd):/converted/ arachnysdocker/athenapdf athenapdf local_file.html
```

#### Remote URL

```bash
docker run --rm -v $(pwd):/converted/ arachnysdocker/athenapdf athenapdf http://blog.arachnys.com/
```

#### Arguments / flags

To see a list of available CLI options:

```bash
docker run --rm -v $(pwd):/converted/ arachnysdocker/athenapdf athenapdf --help
```

For example, you can write the output of the conversion to [standard output][stdout] using the `-S` flag, e.g.

```bash
docker run --rm -v $(pwd):/converted/ arachnysdocker/athenapdf athenapdf -S http://blog.arachnys.com/
```

There is also a [flag][aggressive] for rendering a HTML document to a screen reader / mobile-friendly PDF. It is perfect for news articles, and blog posts. See [aggressive.md][aggressive].


## Tips / Tricks

See [`tips.md`](tips.md).



[docker]: https://www.docker.com/
[docker-machine]: https://docs.docker.com/mac/step_one/
[1.10.0]: https://github.com/docker/docker/releases/tag/v1.10.0
[stdout]: https://en.wikipedia.org/wiki/Standard_streams#Standard_output_.28stdout.29
[aggressive]: aggressive.md
