
var userRole = getAccessLvl(); /*GET accesslvl from Account structs*/

if (userRole == 'admin') {
    renderAdminView();
}
else if (userRole = 'verified-user') {
    renderVerfiedView();
}
else{
    renderDefaultView();
}


