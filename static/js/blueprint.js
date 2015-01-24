/*global $*/
/*global document*/
/*global GK*/
/*global ace*/

$(document).ready(function() {
    var editor = ace.edit("editor");
    editor.setTheme("ace/theme/tomorrow_night_eighties");
    editor.getSession().setMode("ace/mode/golang");
});