var delay = function(ms) {
    return new Promise(function(resolve) {
        setTimeout(function() {
            resolve(true);
        }, 1000);
    });
};

if (typeof ATHENAPDF_DELAY !== "undefined") {
    var ATHENAPDF_DELAY = 1000;
}

delay(ATHENAPDF_DELAY);
