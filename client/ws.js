var ws = new WebSocket("ws://" + location.host + "/ws");

// set event handler
ws.onopen = onOpen;
ws.onmessage = onMessage;
ws.onclose = onClose;
ws.onerror = onError;

// open event
function onOpen(event) {
	console.log("onOpen()");
	console.log(event);

	sendData();
}

// receive data event
function onMessage(event) {
	console.log("onMessage()");

	if(event && event.data){
		console.log(event);

		var data = JSON.parse(event.data);
		var elapse = new Date().getTime() - data.time

		document.getElementById("res").innerHTML =
			elapse + "ms<br>" +
			document.getElementById("res").innerHTML;
	}

	setTimeout(sendData, 1000);
}

// error event
function onError(event) {
	console.log("onError()");
	console.log(event);
}

// close event
function onClose(event) {
	console.log("onClose()");
	console.log(event);
}

// send data
function sendData() {
	console.log("sendData()");

	var data = {
		cmd:  "ping",
		time: new Date().getTime()
	};
	var json = JSON.stringify(data);

	console.log("data: " + data);
	console.log("json: " + json);
	ws.send(json);
}
