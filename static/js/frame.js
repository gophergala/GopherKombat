/*global $*/
/*global document*/
/*global GK*/
/*global window*/

GK.load = function(page) {
     $("#content").load(page);
};

$(document).ready(function() {
    GK.load("home");
    $(".nav").click(function() {
        var page = $(this).attr("rel");
        GK.load(page);
    });
    $(".login").click(function() {
        var url = "https://github.com/login/oauth/authorize",
            client_id = "fe6528d512e0697b7883";
        window.location.href = url + "?" + "client_id=" + client_id;
    });
    $(".logout").click(function() {
        window.location.href = "/logout";
    });
});