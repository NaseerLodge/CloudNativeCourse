function multiply() {
	var string = document.getElementById("string").value;
	var number = document.getElementById("number").value;
	var xhr = new XMLHttpRequest();
	xhr.open("POST", "/multiply?string=" + string + "&number=" + number);
	xhr.onload = function() {
		document.getElementById("result").innerHTML = xhr.responseText;
	};
	xhr.send();
}
