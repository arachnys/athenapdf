# Aggressive Mode

Ever used [Readability][readability], [Instapaper][instapaper] or [Pocket][pocket]?

Do you like their distraction-free reading experience / content extraction?

Enable aggressive mode when rendering PDFs, and get rid of web page clutter / noise! Get a mobile-optimised / screen reader friendly version of a HTML page as a PDF.

Simply pass a `-A` or `--aggressive` flag to render with [`dom-distiller`][domdistiller] (by Chromium), our page simplifier.

**Example:**

```bash
docker run --rm -v $(pwd):/converted/ arachnysdocker/athenapdf athenapdf -A http://blog.arachnys.com/google-isnt-even-close-to-proper-due-diligence.-why-not
```



[readability]: https://www.readability.com/
[instapaper]: https://www.instapaper.com/
[pocket]: https://getpocket.com/
[domdistiller]: https://github.com/chromium/dom-distiller
