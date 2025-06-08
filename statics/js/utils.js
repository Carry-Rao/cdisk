//获取网盘名称
console.log("utils.js loaded");
window.onload = function () {
xhr = new XMLHttpRequest();
xhr.open("GET", "/api/system/get_name", true);
xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
        var res = JSON.parse(xhr.responseText);
        //使用vue.min.js渲染网盘名称
        new Vue({
            el: 'title',
            data: {
                title: res.name
            }
        })
        new Vue({
            el: '#headbar',
            data: {
                title: res.name
            }
        })
    }
}
xhr.send();
}