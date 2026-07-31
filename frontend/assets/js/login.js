function login(){

    var emailInput = document.querySelector('input[type="email"]');
    var passwordInput = document.querySelector('input[type="password"]');

    var emailValue = emailInput.value.trim();
    var passwordValue = passwordInput.value.trim();

    var emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

    var isAdminUsername = (emailValue === "admin");
    var isValidEmail = emailPattern.test(emailValue);
    var isPasswordCorrect = (passwordValue === "admin");

    var isValid = (isAdminUsername || isValidEmail) && isPasswordCorrect;

    if(isValid){
        window.location.href = "dashboard.html";
    } else {
        showLoginError(emailInput);
    }

}

function showLoginError(emailInput){

    var existingError = document.querySelector('.login-error');

    if(existingError){
        existingError.remove();
    }

    var errorEl = document.createElement('p');
    errorEl.className = 'login-error';
    errorEl.textContent = "Invalid credentials. Enter valid email or password .";
    errorEl.style.color = "#f87171";
    errorEl.style.fontSize = "13px";
    errorEl.style.marginTop = "-6px";
    errorEl.style.marginBottom = "14px";
    errorEl.style.fontFamily = "'Inter', sans-serif";

    emailInput.parentNode.insertBefore(errorEl, emailInput);

}