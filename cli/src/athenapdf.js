"use strict";

const crypto = require("crypto");
const fs = require("fs");
const path = require("path");
const rw = require("rw");
const url = require("url");

const athena = require("commander");
const electron = require("electron");
const app = electron.app;
const BrowserWindow = electron.BrowserWindow;

const mediaPlugin = fs.readFileSync(path.join(__dirname, "./plugin_media.js"), "utf8");

var bw = null;
var ses = null;
var uriArg = null;
var outputArg = null;

athena
    .version("2.1.3")
    .description("convert HTML to PDF via stdin or a local / remote URI")
    .option("--debug", "show GUI", false)
    .option("-T, --timeout <seconds>", "seconds before timing out (default: 120)", parseInt)
    .option("-D, --delay <milliseconds>", "milliseconds delay before saving (default: 200)", parseInt)
    .option("-P, --pagesize <size>", "page size of the generated PDF (default: A4)", /^(A3|A4|A5|Legal|Letter|Tabloid)$/i, "A4")
    .option("-S, --stdout", "write conversion to stdout")
    .option("-A, --aggressive", "aggressive mode / runs dom-distiller")
    .option("-B, --bypass", "bypasses paywalls on digital publications (experimental feature)")
    .option("--no-portrait", "render in landscape")
    .option("--no-background", "omit CSS backgrounds")
    .option("--no-cache", "disables caching")
    .option("--proxy <url>", "use proxy to load remote HTML")
    .arguments("<URI> [output]")
    .action((uri, output) => {
        uriArg = uri;
        outputArg = output;
    })
    .parse(process.argv);

// Display help information by default
if (!process.argv.slice(2).length) {
    athena.outputHelp();
}

if (!uriArg) {
    console.error("No URI given. Set the URI to `-` to pipe HTML via stdin.");
    process.exit(1);
}

// Handle stdin
if (uriArg === "-") {
    let base64Html = new Buffer(rw.readFileSync("/dev/stdin", "utf8"), "utf8").toString("base64");
    uriArg = "data:text/html;base64," + base64Html;
// Handle local paths
} else if (uriArg.toLowerCase().indexOf("http") !== 0) {
    uriArg = url.format({
        protocol: "file",
        pathname: path.resolve(uriArg),
        slashes: true
    });
}

// Generate SHA1 hash if no output is specified
if (!outputArg) {
    const shasum = crypto.createHash("sha1");
    shasum.update(uriArg);
    outputArg = shasum.digest("hex") + ".pdf";
}

// Built-in timeout (exit) when debugging is off
if (!athena.debug) {
    setTimeout(() => {
        console.error("PDF generation timed out.");
        app.exit(2);
    }, (athena.timeout || 120) * 1000);
}

if (athena.proxy) {
    console.log("Using proxy", athena.proxy);
    app.commandLine.appendSwitch("proxy-server", athena.proxy);
}

// Preferences
const bwOpts = {
    show: (athena.debug || false),
    webPreferences: {
        nodeIntegration: false,
        webSecurity: false
    }
};

const loadOpts = {
    "extraHeaders": athena.cache ? "" : "pragma: no-cache\n"
};

const pdfOpts = {
    pageSize: athena.pagesize,
    printBackground: athena.background,
    landscape: !athena.portrait
};

// Utils
const _complete = () => {
    console.timeEnd("PDF Conversion");
    athena.debug || app.quit();
};

const _output = (data) => {
    const outputPath = path.join(process.cwd(), outputArg);
    if (athena.stdout) {
        process.stdout.write(data, _complete);
    } else {
        fs.writeFile(outputPath, data, (err) => {
            if (err) console.error(err);
            console.info(`Converted '${uriArg}' to PDF: '${outputArg}'`);
            _complete();
        });
    }
};

app.on("ready", () => {
    console.time("PDF Conversion");

    bw = new BrowserWindow(bwOpts);

    bw.on("closed", () => { bw = null; });

    bw.loadURL(uriArg, loadOpts);

    if (athena.bypass) {
        const _cookieWhitelist = ["nytimes", "ft.com"];
        const _inCookieWhitelist = (url) => {
            let matches = _cookieWhitelist.filter((safe) => {
                return url.indexOf(safe) !== -1;
            });
            return (matches.length !== 0);
        };
        ses = bw.webContents.session;
        ses.webRequest.onBeforeSendHeaders((details, callback) => {
            if (details.resourceType === "mainFrame") {
                if (!_inCookieWhitelist(details.url)) {
                   delete details.requestHeaders["Cookie"];
                }
                details.requestHeaders["Referer"] = "https://www.google.com/";
                details.requestHeaders["User-Agent"] = "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)";
            }
            callback({cancel: false, requestHeaders: details.requestHeaders});
        });
    }

    bw.webContents.on("did-fail-load", (e, code, desc, url, isMainFrame) => {
        if (parseInt(code, 10) >= -3) return;
        console.error(`Failed to load: ${code} ${desc} (${url})`);
        if (isMainFrame) {
            app.exit(1);
        }
    });

    bw.webContents.on("crashed", () => {
        console.error(`The renderer process has crashed.`);
        app.exit(1);
    });

    // Load plugins
    let plugins = mediaPlugin + "\n";
    if (athena.aggressive) {
        const distillerPlugin = fs.readFileSync(path.join(__dirname, "./plugin_domdistiller.js"), "utf8");
        plugins += distillerPlugin;
    }
    bw.webContents.executeJavaScript(plugins);

    bw.webContents.on("did-finish-load", () => {
        setTimeout(() => {
            bw.webContents.printToPDF(pdfOpts, (err, data) => {
                if (err) console.error(err);
                _output(data);
            });
        }, (athena.delay || 200));
    });
});

app.on("window-all-closed", () => {
    if (process.platform !== "darwin") {
        app.quit();
    }
});
