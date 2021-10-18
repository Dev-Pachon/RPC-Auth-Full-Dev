console.log("Script loaded");

fetch("/home", {
    method: "SELECT"
}).then((response) => {
    response.text().then(function (data) {
        let result = JSON.parse(data);
        console.log(result)
    });
}).catch((error) => {
    console.log(error)
});