var waitForWindowStatus = function(desiredStatus) {
    return new Promise(function(resolve) {
        var poller = setInterval(function() {
            if (window.status && window.status == desiredStatus) {
                clearInterval(poller);
                resolve(desiredStatus);
            }
        }, 100);
    });
};

if (typeof WINDOW_STATUS === "undefined") {
    var WINDOW_STATUS = "ready";
}

waitForWindowStatus(WINDOW_STATUS);
