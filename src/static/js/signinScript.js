console.log("Script loaded");

const input = document.getElementsByTagName("input");

const inputForm = document.getElementById("inputForm");

inputForm.addEventListener("submit", (e)=> {

    e.preventDefault();

    var check = true;

    for (var i = 0; i < input.length; i++) {
        if (validate(input[i]) == false) {
            showValidate(input[i]);
            check = false;
        }
    }

    if (!check) {
        return check;
    }

    let data = {
        Username: document.getElementById("username").value,
        Password: document.getElementById("password").value,
    };

    fetch("/signin", {
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        method: "POST",
        body: JSON.stringify(data)
    }).then((response) => {
        response.text().then(function (data) {
            let result = JSON.parse(data);
            if (result["Result"] = "ok") {
                //location.href("/home");
            }else{
                //limpiar campos
                alert(result["Content"]);
            }
        });
    }).catch((error) => {
        console.log(error)
    });

    return check;
});


$('.validate-form .input100').each(function () {
    $(this).focus(function () {
        hideValidate(this);
    });
});

function validate(input) {
    if ($(input).val().trim() == '') {
        return false;
    }
}

function showValidate(input) {
    var thisAlert = $(input).parent();

    $(thisAlert).addClass('alert-validate');
}

function hideValidate(input) {
    var thisAlert = $(input).parent();

    $(thisAlert).removeClass('alert-validate');
}