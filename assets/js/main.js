"use strict";

(function() {
    document.getElementById("login").addEventListener("click", login);
}());

function login() {
    let username = document.querySelector("#username").value;
    let passphrase = document.querySelector("#passphrase").value;

    let data = new FormData();
    data.append("username", username);
    data.append("passphrase", passphrase);

    makeRequest("/login", data);
}

function makeRequest(endpoint, data) {
    let xhr = new XMLHttpRequest();
    xhr.open('POST', endpoint, true);
    xhr.onload = function () {
        // do something to response
        console.log(this.responseText);
    };
    xhr.send(data);
}