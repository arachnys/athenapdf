# Building

`athenapdf` can be packaged as a _binary_ for easy distribution using `electron-packager`.

In fact, the Docker version of `athenapdf` runs on the packaged Electron app, not the unpacked source code which the NPM version uses.


## Docker

To build a new Docker image run `make buildcli` in the project's root directory.

Open `Makefile` for more information or if you would like to run the build commands manually.


## NPM

Currently, the following commands will package `athenapdf` for OSX, and GNU / Linux x64 distribution:

1. `npm install`
2. `npm run build`
3. Check the `build/` directory for the respective binaries
