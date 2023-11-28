
var userRole = getAccessLvl(); /*GET accesslvl from API, think it's stored in userDetails under Login in db.handle.go??*/

if (userRole == 'admin') {
    renderAdminView();
}
else if (userRole = 'verified-user') {
    renderVerfiedView();
}
else{
    renderDefaultView();
}


function renderAdminView(){
    window.location.href=""
}
function renderVerifiedView(){
    window.location.href="verified_view.html"
}
function renderDefaultView(){
    window.location.href="guest_view.html"
}