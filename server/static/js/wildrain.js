
var applications;

function showAicd(eId, appName, ver) {
    var e = $('#' + eId);
    var app = applications[appName][ver];
    var html = JSON.stringify(app, null, '  ');
    e.html(appName + " / " + ver + "<br /><pre>" + html + "</pre>");
}

function renderAicd(d, aicdLinksId, aicdJsonId) {
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

function getApplications(aicdLinksId, aicdJsonId) {
    $(function() {
        $.ajax({
            'url': '../getApplications',
            'success': function(d, s, x) {
                renderAicd(d, aicdLinksId, aicdJsonId);
            }
        });
    });
}
