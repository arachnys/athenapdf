var print = document.querySelectorAll('[type="text/css"][media*="print"]');
var screen = document.querySelectorAll('[type="text/css"][media*="screen"]');

if (print.length === 0) {
    for (var i = 0, l = screen.length; i < l; i++) {
        screen[i].removeAttribute("media");
    }
}