
var ansi_up = new AnsiUp;

var socket = new WebSocket("ws://localhost:9000/ws/");

socket.onopen = function() {
	console.log("WS opened");
}

socket.onmessage = function(event) {
	let data = event.data;
	var html = ansi_up.ansi_to_html(data);
	console.log(html);
	$('#terminal').append(html+"</br>");
	$('#terminal').animate({ scrollTop: $('#terminal')[0].scrollHeight }, 10);
};



socket.onerror = function(error) {
  alert(`[error] ${error.message}`);
};

socket.onclose = function(event) {
	if (event.wasClean) {
    alert(`[close] Соединение закрыто чисто, код=${event.code} причина=${event.reason}`);
  } else {
    alert('[close] Соединение прервано');
  }
}

function sendData() {
	var data = $('#in').val();
	console.log(data);
	socket.send(data);
}







