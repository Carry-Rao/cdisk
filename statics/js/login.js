var xhr = new XMLHttpRequest();
xhr.open("GET", "/api/system/get_name", true);
xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
        var res = xhr.responseText;
        //使用vue.min.js渲染网盘名称
        document.title = "登录 - " + res;
        document.querySelector("#app h1").innerHTML = res;
    }
}
xhr.send();

function login() {
    var password = document.querySelector("#password");
    password.value = CryptoJS.SHA256(password.value);
    document.getElementById("login-form").submit()
}
