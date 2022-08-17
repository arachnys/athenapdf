const config = {
    "paths": {
        "actualPdfRootFolder": process.cwd() + "/files",
        "baselinePdfRootFolder": process.cwd() + "/files",
        "actualPngRootFolder": process.cwd() + "/files/new",
        "baselinePngRootFolder": process.cwd() + "/files/baseline",
        "diffPngRootFolder": process.cwd() + "/files/diff"
    },
    "settings": {
        "imageEngine": "graphicsMagick",
        "density": 100,
        "quality": 70,
        "tolerance": 0,
        "threshold": 0.05,
        "cleanPngPaths": true,
        "matchPageCount": true,
        "disableFontFace": true,
	    "verbosity": 10
    }
}

module.exports = config;