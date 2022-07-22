function httpGet(lifter_name) {
    let xmlHttpReq = new XMLHttpRequest();
    xmlHttpReq.open("GET", "/api/lookup/".concat(lifter_name), false);
    xmlHttpReq.send(null);
    return xmlHttpReq.responseText;
}
console.log(httpGet('euan'));