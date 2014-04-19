
var applications;
var aicdLinksId;
var aicdJsonId;

function ajax(options) {
    var ws = new WebSocket("ws://" + document.domain + ":8080/" + options['url']);
    ws.onmessage = function(e) {
        options['handleMessage'](JSON.parse(e.data));
        ws.onclose = function() {};
        ws.close();
    }

    ws.onopen = function() {
        var data = options['data']
        if (data != undefined) {
            ws.send(JSON.stringify(data));
        }
    }
}

function handleFileSelect(evt) {
    var files = evt.target.files;

    for (var i = 0, f; f = files[i]; i++) {
        var reader = new FileReader();
        reader.onload = function(e) {
            var text = reader.result;
            var o = JSON.parse(text);
            ajax({
                'url': 'uploadAicd',
                'data': o,
                'handleMessage': function(received) {
                    _getApplications();
                }
            });
        }
        reader.readAsText(f, 'utf-8');
    }
}

function showAicd(eId, appName, ver) {
    var e = $('#' + eId);
    var app = applications[appName][ver];
    var html = JSON.stringify(app, null, '  ');
    e.html(appName + " / " + ver + "<br /><pre>" + html + "</pre>");
}

function renderAicd(d) {
    applications = d;
    var html = "";
    $.each(applications, function(key, value) {
        html += "<strong>" + key + "</strong><br /><ul>"; 
        $.each(value, function(k, v) {
            html += '<li><button onclick="showAicd(\'' + aicdJsonId + '\', \'' + key + '\', \'' + k + '\')" type="button">Version ' + k + '</button></li>';
        });
        html += "</ul>";
    });
    $('#' + aicdLinksId).html(html);
}

function _getApplications() {
    ajax({
        'url': 'getApplications',
        'handleMessage': function(d) {
            renderAicd(d);
        }
    });
}

function getApplications(_aicdLinksId, _aicdJsonId) {
    aicdLinksId = _aicdLinksId;
    aicdJsonId = _aicdJsonId;
    $(function() {
        _getApplications();
    });
}

$(function() {
    document.getElementById('fileupload').addEventListener('change', handleFileSelect, false);
});

