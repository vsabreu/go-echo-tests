"use strict";

(function() {
    document.getElementById("createuser").addEventListener("click", createUser);
}());

function loadUsers() {
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/users", true);
    console.log(getCookieValue("_csrf"));
    xhr.setRequestHeader("X-CSRF-Token", getCookieValue("_csrf"));
    xhr.onload = function () {
        let users = JSON.parse(this.responseText);
        
    }
    xhr.send(null);
}

function createUser() {
    
    let username = document.querySelector("#username");
    let email = document.querySelector("#email");

    const user = {
        "name": username.value,
        "email": email.value
    }

    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/users", true);
    console.log(getCookieValue("_csrf"));
    xhr.setRequestHeader("X-CSRF-Token", getCookieValue("_csrf"));
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.onload = function () {
        let newUser = JSON.parse(this.responseText);
        username.value = "";
        email.value = "";
    }
    xhr.send(JSON.stringify(user));
}

function getCookieValue(a) {
    let b = document.cookie.match('(^|;)\\s*' + a + '\\s*=\\s*([^;]+)');
    return b ? b.pop() : '';
}