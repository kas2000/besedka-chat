// window.onload = function () {
var connWithUrl;
var conn;
var msg = document.getElementById("msg");
var userID = document.getElementById("user-id");
var log = document.getElementById("log");

function test() {
    console.log(conn)
}

function appendLog(item) {
    var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
    log.appendChild(item);
    if (doScroll) {
        log.scrollTop = log.scrollHeight - log.clientHeight;
    }
}

if (window["WebSocket"]) {
    conn = new WebSocket("ws://" + document.location.host + "/ws");
    // conn = connWithUrl
    conn.onclose = function (evt) {
        var item = document.createElement("div");
        item.innerHTML = "<b>Connection closed.</b>";
        appendLog(item);
    };
    conn.onmessage = function (evt) {
        var messages = evt.data.split('\n');
        for (var i = 0; i < messages.length; i++) {
            var item = document.createElement("div");
            item.innerText = messages[i];
            appendLog(item);
        }
    };
} else {
    var item = document.createElement("div");
    item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
    appendLog(item);
}

function sendMess() {
    console.log(conn)
    console.log("HELLO world");
    if (!conn) {
        console.log("HELLO world)");
        return false;
    }
    if (!msg.value) {
        console.log("HELLO world:)");
        return false;
    }
    conn.send(msg.value);
    // msg.value = "";
    return false;
}