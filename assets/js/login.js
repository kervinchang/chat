const rememberMeCheckbox = document.getElementById("rememberMe");
const usernameInput = document.getElementById("username");

// 在页面加载时检查是否存在cookie，若存在则自动填入用户名
if(getCookie("remembered_username")) {
    usernameInput.value = getCookie("remembered_username");
    rememberMeCheckbox.checked = true;
}

// 当用户取消勾选Remember Me时，删除保存在cookie中的用户名
rememberMeCheckbox.addEventListener('change', function() {
    if(!rememberMeCheckbox.checked) {
        deleteCookie("remembered_username");
    }
});

// 当用户选中Remember Me时，保存用户名至cookie中
document.querySelector('form').addEventListener('submit', function() {
    if(rememberMeCheckbox.checked) {
        setCookie("remembered_username", usernameInput.value, 7);
    }
});

// 设置cookie
function setCookie(key, value, day) {
    var expires = new Date();
    expires.setTime(expires.getTime() + (day * 24 * 60 * 60 * 1000));
    document.cookie = key + '=' + value + ';expires=' + expires.toUTCString();
}

// 获取cookie
function getCookie(key) {
    var keyValue = document.cookie.match('(^|;) ?' + key + '=([^;]*)(;|$)');
    return keyValue ? keyValue[2] : null;
}

// 删除cookie
function deleteCookie(key) {
    document.cookie = key +'=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
}