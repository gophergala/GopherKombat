/*global $*/
/*global document*/
/*global GK*/
/*global ace*/
/*global console*/

$(document).ready(function() {
    var editor1 = ace.edit("editor1");
    editor1.setTheme("ace/theme/tomorrow_night_eighties");
    editor1.getSession().setMode("ace/mode/golang");
    var editor2 = ace.edit("editor2");
    editor2.setTheme("ace/theme/tomorrow_night_eighties");
    editor2.getSession().setMode("ace/mode/golang");
    $(".submitCode").click(function() {
        var code1 = editor1.getSession().getValue(),
            code2 = editor2.getSession().getValue(),
            data,
            url = "/perf/submit";
        data = {
            code1: code1,
            code2: code2
        };
        GK.requestAgent().doPOST(url, data, function(resp) {
            console.log(resp);
            if (!resp.success) {
                $("#validation1").html(resp.resp.err1);
                $("#validation2").html(resp.resp.err2);
                $("#validation1").show();
                $("#validation2").show();
            } else {
                $(".submitCode").hide();
                $('.results').show();
                console.log(resp.resp);
                $("#res1 .time").html(compareTime(resp.resp.t1, resp.resp.t2));
                $("#res2 .time").html(compareTime(resp.resp.t2, resp.resp.t1));
                $("#res1 .memory").html(formatMem(resp.resp.m1));
                $("#res2 .memory").html(formatMem(resp.resp.m2));
                $("#res1 .lines").html(editor1.session.getLength());
                $("#res2 .lines").html(editor2.session.getLength());
            }          
        });
    });
    $(".another-one").click(function() {
        $('.results').hide();
        $(".submitCode").show();
    });
});

function compareTime(v1, v2) {
    var ratio = v1/v2*100 - 100, 
        sufix = v1>v2 ? "% slower" : "% faster";
    return Math.abs(ratio.toFixed(2)) + sufix;
}
function formatMem(v1) {
    var val = v1/1000,
        sufix = " kbytes";
    return Math.abs(val.toFixed(2)) + sufix;
}
