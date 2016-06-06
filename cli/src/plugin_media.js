var print = document.querySelectorAll('[rel="stylesheet"][media*="print"]');
var screen = document.querySelectorAll('[rel="stylesheet"][media*="screen"]');

if (print.length === 0) {
    for (var i = 0, l = screen.length; i < l; i++) {
        screen[i].removeAttribute("media");
    }
}