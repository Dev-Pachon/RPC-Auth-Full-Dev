console.log("Script loaded");

const input = document.getElementsByTagName("input");

const inputForm = document.getElementById("inputForm");

//Submit action

inputForm.addEventListener("submit", (e)=> {

        e.preventDefault();

    var check = true;

    for (let i = 0; i < input.length; i++) {
        if (!validate(input, i)) {
            showValidate(input[i]);
            check = false;
        }
    }

    if (!check) {
        return check;
    }

    let data = {
        Username: document.getElementById("username").value,
        Email: document.getElementById("email").value,
        Password: document.getElementById("password").value,
        Firstname: document.getElementById("firstname").value,
        Lastname: document.getElementById("lastname").value,
        Birthdate: document.getElementById("birthdate").value,
        Country: document.getElementById("country").value,
        University: document.getElementById("university").value
    };

    fetch("/signup", {
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        method: "POST",
        body: JSON.stringify(data)
    }).then((response) => {
        response.text().then(function (data) {
            let result = JSON.parse(data);
            if (result["Result"] ==="ok") {
                location.href = "/signin";
            }else{
                alert(result["Content"]);
            }
        });
    }).catch((error) => {
        console.log(error)
    });

    return check;
});

//Hide the alert when the user focus the input.
$('.validate-form .input100').each(function () {
    $(this).focus(function () {
        hideValidate(this);
    });
});

//Validate the input element with certains constraints.
function validate(input, i) {

    var element = input[i];

    if (element.value == "") {
        parent = $(element).parent();
        parent.attr("data-validate", element.getAttribute("name") + " is required!");
        return false;
    }
    switch (element.getAttribute("name")) {
        case "Username":
            if (element.value.length < 8) {
                parent = $(element).parent();
                parent.attr("data-validate", "Use 8 characters or more for your username");
                return false;
            }
            regularExpr = /[A-Za-z0-9]/;
            if ($(element).val().match(regularExpr) == null) {
                parent = $(element).parent();
                parent.attr("data-validate", "Use only letters and number for your username");
                return false;
            }
            regularExpr = /[ ]/;
            if ($(element).val().match(regularExpr) != null) {
                parent = $(element).parent();
                parent.attr("data-validate", "Your username can’t have a blank space");
                return false;
            }
            break;
        case "Password":
            if (element.value.length < 8) {
                parent = $(element).parent();
                parent.attr("data-validate", "Use 8 characters or more for your password");
                return false;
            }
            regularExpr = /[ ]/;
            if ($(element).val().match(regularExpr) != null) {
                parent = $(element).parent();
                parent.attr("data-validate", "Your password can’t have a blank space");
                return false;
            }
            break;

        case "Email":
            regularExpr = /[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}/;
            if ($(element).val().match(regularExpr) == null) {
                parent = $(element).parent();
                parent.attr("data-validate", "Use a valid email");
                return false;
            }
            break;
        case "Confirm password":
            
            if (element.value.localeCompare(input[i - 1].value) != 0) {
                parent = $(element).parent();
                parent.attr("data-validate", "The passwords aren't the same");
                return false;
            }
            break;

        case "Birthdate":
            yearInput = element.value.split("-")[0];
            yearActual = (new Date()).getFullYear();
            if (!(yearActual - yearInput >= 5 && yearActual - yearInput <= 200)) {
                parent = $(element).parent();
                parent.attr("data-validate", "The birthdate is invalid");
                return false;
            }
            break;

    }

    return true;

}

//show an alert
function showValidate(input) {
    var thisAlert = $(input).parent();

    $(thisAlert).addClass('alert-validate');
}

//hide the alert
function hideValidate(input) {
    var thisAlert = $(input).parent();

    $(thisAlert).removeClass('alert-validate');
}