
function getApplications(paneId) {
    $(function() {
        $.ajax({
            'url': '../getApplications',
            'success': function(d, s, x) {
                console.log(d);
            }
        });
    });
}
