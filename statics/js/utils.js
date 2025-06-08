console.log("utils.js loaded");
xhr = new XMLHttpRequest();

function get_name() {
xhr.open("GET", "/api/system/get_name", true);
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4 && xhr.status == 200) {
            var res = xhr.responseText;
            document.title = res;
            document.querySelector("#header h1").innerHTML = res;
        }
    }
    xhr.send();
}

get_name();

function register() {
    window.location.href = "/register";
}