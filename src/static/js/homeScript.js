console.log("Script loaded");
const tbody = document.getElementById('tbody');


fetch("/home", {
    method: "SELECT"
}).then((response) => {
    response.text().then(function (data) {
        let result = JSON.parse(data);
        console.log(result);


        for(let i=0;i<result.length;i++){
            let row = document.createElement('tr');
            let html = "";
            for(var property in result[i]) {
                html += "<td>" + result[i][property] + "</td>";
            }
            row.innerHTML = html;
            tbody.appendChild(row);
        }

    });
}).catch((error) => {
    console.log(error)
});

