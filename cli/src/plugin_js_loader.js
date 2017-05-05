const jsLoaderPromise = new Promise((resolve, reject) => {
    try {
        let athenaPdfReadyTimer;
        const pollAthenaReadyState = () => {
            athenaPdfReadyTimer = setTimeout(() => {
                if (window.ATHENA_PDF_READY) {
                    resolve(true);
                } else {
                    pollAthenaReadyState();
                }
            }, 250);
        }

        // poll for `ATHENA_PDF_READY` present
        pollAthenaReadyState();

        // overall timeout
        setTimeout(() => {
            clearTimeout(athenaPdfReadyTimer);
            throw "Error: timeout: `window.ATHENA_PDF_READY` variable not found.";
        }, 30000);
    } catch (error) {
        reject(error)
    }
});

jsLoaderPromise.then(result => result);
