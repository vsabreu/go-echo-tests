"use strict";

(function() {
    document.getElementById("login").addEventListener("click", login);
    document.getElementById("btnRestricted").addEventListener("click", goToRestricted);
}());

function login() {
    let username = document.querySelector("#username").value;
    let passphrase = document.querySelector("#passphrase").value;

    let data = new FormData();
    data.append("username", username);
    data.append("passphrase", passphrase);

    makeRequest("POST", "/login", data);
}

function goToRestricted() {
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/users", true);
    console.log(getCookieValue("_csrf"));
    xhr.setRequestHeader("X-CSRF-Token", getCookieValue("_csrf"));
    xhr.onload = function () {
        alert(this.responseText);
    }
    xhr.send(null);
}

function makeRequest(method, endpoint, data, fn) {
    let xhr = new XMLHttpRequest();
    xhr.open(method, endpoint, true);
    xhr.setRequestHeader("X-CSRF-Token", getCookieValue("_csrf"));
    xhr.onload = function () {
        let rs = JSON.parse(this.responseText);
        if (rs.token) {
            document.querySelector(".hid").classList.toggle("hid");
        }
    }
    xhr.send(data);
}

function getCookieValue(a) {
    let b = document.cookie.match('(^|;)\\s*' + a + '\\s*=\\s*([^;]+)');
    return b ? b.pop() : '';
}