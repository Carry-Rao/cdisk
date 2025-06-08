xhr.open("GET", "/api/system/get_name", true);
xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
        var res = xhr.responseText;
        //使用vue.min.js渲染网盘名称
        document.title = "登录 - " + res;
        new Vue({
            el: '#app h1',
            data: {
                title: res
            }
        })
    }
}
xhr.send();
document.addEventListener('DOMContentLoaded', function() {
    new Vue({
        el: '#app',
        methods: {
            register: function() {
                // 实现注册逻辑
                window.location.href = '/register'; // 跳转到注册页面
            },
            resetPassword: function() {
                // 实现忘记密码逻辑
                window.location.href = '/reset-password'; // 跳转到忘记密码页面
            }
        }
    });
});
