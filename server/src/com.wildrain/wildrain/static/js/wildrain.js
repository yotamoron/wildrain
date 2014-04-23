
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

function renderAppsSelectors(app, ver, event) {
	$('#' + app).html("");
	$('#' + ver).html("");
	$('#' + event).html("");
	var appSelect = '<select id="appSelectDropdown">';
	$.each(applications, function(key, value) {
		appSelect += '<option value="' + key + '">' + key + '</option>';
	});
	appSelect += "</select>"
	$('#' + app).html(appSelect);
	$('#appSelectDropdown').prop("selectedIndex", -1);
	renderVersionsSelectors(ver, event);
}

function renderVersionsSelectors(ver, event) {
	$('#appSelectDropdown').change(function(e) {
		var app = $('#appSelectDropdown').val();
		var verSelect = '<select id="verSelectDropdown">';
		$.each(applications[app], function(k, v) {
			verSelect += '<option value="' + k + '">' + k + '</option>';
		});
		verSelect += "</select>"
		$('#' + ver).html(verSelect);
		$('#verSelectDropdown').prop("selectedIndex", -1);
		renderEventsSelectors(event);
	});

}

function renderEventsSelectors(event) {
	$('#verSelectDropdown').change(function(e) {
		var app = $('#appSelectDropdown').val();
		var ver = $('#verSelectDropdown').val();
		var eventSelect = '<select id="eventSelectDropdown">';
		$.each(applications[app][ver]['Events'], function(k, v) {
			var eventName = v['Name'];
			eventSelect += '<option value="' + eventName + '">' + eventName + '</option>';
		});
		eventSelect += "</select>"
		$('#' + event).html(eventSelect);
		$('#eventSelectDropdown').prop("selectedIndex", -1);
	});
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
			renderAppsSelectors("appSelect", "verSelect", "eventSelect");
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

