function register() {
    console.log("register");
    var password = document.getElementById("password");
    var confirm_password = document.getElementById("confirm_password");
    if (password.value!= confirm_password.value) {
        return; 
    }
    password.value = CryptoJS.SHA256(password.value);
    document.getElementById("register_form").submit();
}