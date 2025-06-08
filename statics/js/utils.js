console.log("utils.js loaded");
xhr = new XMLHttpRequest();

function get_name() {
xhr.open("GET", "/api/system/get_name", true);
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4 && xhr.status == 200) {
            var res = xhr.responseText;
            //使用vue.min.js渲染网盘名称
            document.title = res;
            new Vue({
                el: '#headbar',
                data: {
                    title: res
                }
            })
        }
    }
    xhr.send();
}

get_name();