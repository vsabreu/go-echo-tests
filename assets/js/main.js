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

    makeRequest("POST", "/login", data);
}

function makeRequest(method, endpoint, data, fn) {
    let xhr = new XMLHttpRequest();
    xhr.open(method, endpoint, true);
    xhr.onload = function () {
        let rs = JSON.parse(this.responseText);
        if (rs.token) {
            document.querySelector(".hid").classList.toggle("hid");
        }
    }
    xhr.send(data);
}