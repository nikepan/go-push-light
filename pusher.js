var pushServer = window.location.hostname + ':8095/';

var pusher = null;
var callbacks = {};

function makeEvent() {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', 'http://' + pushServer + 'pub?intent=channelname.taskname.12345&obj={"id":123}', false);
    xhr.send();
}


function connectPusher(onConnect) {
    pusher = new WebSocket("ws://" + pushServer + "sub");

    pusher.onopen = function () {
        console.log("Pusher Connected");
        onConnect();
    };

    pusher.onclose = function (event) {
        if (event.wasClean) {
            console.log('Pusher Connection closed gracefull');
        } else {
            console.log('Pusher Connection closed')
        }
        console.log('Pusher disconnect code: ' + event.code + ' reason: ' + event.reason);
    };

    pusher.onmessage = function (event) {
        console.log("Pusher Received data: " + event.data);
        var jdata = JSON.parse(event.data);
        var cb = callbacks[jdata.intent];
        if (cb != undefined) {
            var obj = jdata.obj;
            if (obj) {
                try {
                    obj = JSON.parse(obj);
                } catch (e) {

                }
            }
            cb(obj);
        }
    };

    pusher.onerror = function (error) {
        console.log("Pusher Error: " + error.message);
    };
}


function pushSub(intent, callback) {
    if (pusher == null) {
        connectPusher(function () {
            pusher.send('{"op": "sub", "intent": "' + intent + '"}');
            callbacks[intent] = callback;
        });
    } else {
        pusher.send('{"op": "sub", "intent": "' + intent + '"}');
        callbacks[intent] = callback;
    }
}

function pushUnsub(intent) {
    pusher.send('{"op": "unsub", "intent": "' + intent + '"}');
}