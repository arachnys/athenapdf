"use strict";

const crypto = require("crypto");
const fs = require("fs");
const path = require("path");
const rw = require("rw");
const url = require("url");

const athena = require("commander");
const setCookie = require("set-cookie-parser");
const electron = require("electron");
const app = electron.app;
const BrowserWindow = electron.BrowserWindow;

const mediaPlugin = fs.readFileSync(
  path.join(__dirname, "./plugin_media.js"),
  "utf8"
);

var bw = null;
var ses = null;
var uriArg = null;
var outputArg = null;

if (!process.defaultApp) {
  process.argv.unshift("--");
}

const addHeader = (header, arr) => {
  arr.push(header);
  return arr;
};

athena
  .version("2.13.0")
  .description("convert HTML to PDF via stdin or a local / remote URI")
  .option("--debug", "show GUI", false)
  .option(
    "-T, --timeout <seconds>",
    "seconds before timing out (default: 120)",
    parseInt
  )
  .option(
    "-D, --delay <milliseconds>",
    "milliseconds delay before saving (default: 200)",
    parseInt
  )
  .option(
    "-P, --pagesize <size>",
    "page size of the generated PDF (default: A4)",
    /^(A3|A4|A5|Legal|Letter|Tabloid)$/i,
    "A4"
  )
  .option(
    "-M, --margins <marginsType>",
    "margins to use when generating the PDF (default: standard)",
    /^(standard|none|minimal)$/i,
    "standard"
  )
  .option(
    "-Z --zoom <factor>",
    "zoom factor for higher scale rendering (default: 1 - represents 100%)",
    parseInt
  )
  .option("-S, --stdout", "write conversion to stdout")
  .option("-A, --aggressive", "aggressive mode / runs dom-distiller")
  .option(
    "-B, --bypass",
    "bypasses paywalls on digital publications (experimental feature)"
  )
  .option(
    "-H, --http-header <key:value>",
    "add custom headers to request",
    addHeader,
    []
  )
  .option("--proxy <url>", "use proxy to load remote HTML")
  .option("--no-portrait", "render in landscape")
  .option("--no-background", "omit CSS backgrounds")
  .option("--no-cache", "disables caching")
  .option("--ignore-certificate-errors", "ignores certificate errors")
  .option("--ignore-gpu-blacklist", "Enables GPU in Docker environment")
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
  let base64Html = new Buffer(
    rw.readFileSync("/dev/stdin", "utf8"),
    "utf8"
  ).toString("base64");
  uriArg = "data:text/html;base64," + base64Html;
  // Handle local paths
} else if (
  !uriArg.toLowerCase().startsWith("http") &&
  !uriArg.toLowerCase().startsWith("chrome://")
) {
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
  if (!athena.stdout) {
    console.info("Using proxy: ", athena.proxy);
  }
  app.commandLine.appendSwitch("proxy-server", athena.proxy);
}

if (athena.ignoreCertificateErrors) {
  app.commandLine.appendSwitch("ignore-certificate-errors");
}

app.commandLine.appendSwitch(
  "ignore-gpu-blacklist",
  athena.ignoreGpuBlacklist || "false"
);

// Preferences
var bwOpts = {
  show: athena.debug || false,
  webPreferences: {
    nodeIntegration: false,
    webSecurity: false,
    zoomFactor: athena.zoom || 1
  },
  width: 1920,
  height: 1080
};

if (process.platform === "linux") {
  bwOpts["webPreferences"]["defaultFontFamily"] = {
    standard: "Liberation Serif",
    serif: "Liberation Serif",
    sansSerif: "Liberation Sans",
    monospace: "Liberation Mono"
  };
}

// Add custom headers if specified
var extraHeaders = athena.httpHeader;

// Toggle cache headers
if (!athena.cache) {
  extraHeaders.push("pragma: no-cache");
}
const loadOpts = {
  extraHeaders: extraHeaders.join("\n")
};

// Enum for Electron's marginType codes
const MarginEnum = {
  standard: 0,
  none: 1,
  minimal: 2
};

const pdfOpts = {
  pageSize: athena.pagesize,
  marginsType: MarginEnum[athena.margins],
  printBackground: athena.background,
  landscape: !athena.portrait
};

// Utils
const _complete = () => {
  if (!athena.stdout) {
    console.timeEnd("PDF Conversion");
  }
  athena.debug || app.quit();
};

const _output = data => {
  const outputPath = path.join(process.cwd(), outputArg);
  if (athena.stdout) {
    process.stdout.write(data, _complete);
  } else {
    fs.writeFile(outputPath, data, err => {
      if (err) console.error(err);
      console.info(`Converted '${uriArg}' to PDF: '${outputArg}'`);
      _complete();
    });
  }
};

app.on("ready", () => {
  if (!athena.stdout) {
    console.time("PDF Conversion");
  }
  bw = new BrowserWindow(bwOpts);

  bw.on("closed", () => {
    bw = null;
    ses = null;
  });

  bw.loadURL(uriArg, loadOpts);

  ses = bw.webContents.session;
  if (athena.bypass) {
    const _cookieWhitelist = ["nytimes", "ft.com"];
    const _inCookieWhitelist = url => {
      let matches = _cookieWhitelist.filter(safe => {
        return url.indexOf(safe) !== -1;
      });
      return matches.length !== 0;
    };
    ses.webRequest.onBeforeSendHeaders((details, callback) => {
      if (details.resourceType === "mainFrame") {
        if (!_inCookieWhitelist(details.url)) {
          delete details.requestHeaders["Cookie"];
        }
        details.requestHeaders["Referer"] = "https://www.google.com/";
        details.requestHeaders["User-Agent"] =
          "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)";
      }
      callback({ cancel: false, requestHeaders: details.requestHeaders });
    });
  }

  ses.on("will-download", (e, item, webContents) => {
    e.preventDefault();
    console.error(`Unable to convert an octet-stream, use stdin.`);
    app.exit(1);
  });

  const cookieHeaders = athena.httpHeader
    .filter(header => header.includes("Cookie:"))
    .map(header =>
      header.substring(header.indexOf("Cookie:") + "Cookie:".length).trim()
    );
  if (cookieHeaders) {
    setCookie.parse(cookieHeaders, { decodeValues: false }).forEach(cookie => {
      ses.cookies.set(Object.assign(cookie, { url: uriArg }), error => {
        if (error) {
          console.error(error);
        }
      });
    });
  }

  bw.webContents.on("did-fail-load", (e, code, desc, url, isMainFrame) => {
    if (parseInt(code, 10) >= -3) return;
    console.error(`Failed to load: ${code} ${desc} (${url})`);
    if (isMainFrame) {
      app.exit(1);
    }
  });

  bw.webContents.on(
    "did-get-response-details",
    (
      e,
      status,
      newURL,
      originalURL,
      httpResponseCode,
      requestMethod,
      referrer,
      headers,
      resourceType
    ) => {
      if (httpResponseCode >= 400) {
        console.error(
          `Failed to load ${newURL} - got HTTP code ${httpResponseCode}`
        );
        if (resourceType === "mainFrame") {
          app.exit(1);
        }
      }
    }
  );

  bw.webContents.on("crashed", () => {
    console.error(`The renderer process has crashed.`);
    app.exit(1);
  });

  // Load plugins
  let plugins = mediaPlugin + "\n";
  if (athena.aggressive) {
    const distillerPlugin = fs.readFileSync(
      path.join(__dirname, "./plugin_domdistiller.js"),
      "utf8"
    );
    plugins += distillerPlugin;
  }
  bw.webContents.executeJavaScript(plugins);

  bw.webContents.on("did-finish-load", () => {
    setTimeout(() => {
      bw.webContents.printToPDF(pdfOpts, (err, data) => {
        if (err) console.error(err);
        _output(data);
      });
    }, athena.delay || 200);
  });
});

app.on("window-all-closed", () => {
  if (process.platform !== "darwin") {
    app.quit();
  }
});
