let form1 = document.getElementById("manipulate")
let button = form.querySelectorAll("button")

let processClick = (e, type) => {
    let inputs = form1.querySelectorAll(".form > input")

    let data = {}

    for (let i = 0; i < inputs.length; i++) {
        data[inputs[i].name] = inputs[i].value
    }

    let xhr = new XMLHttpRequest()
    xhr.open("POST", `/accounts/${id}/${type}`)
    xhr.onload = function (e) {
        let response = JSON.parse(e.currentTarget.response)
        if ("Error" in response) {
            if (response.Error == null) {
                console.log("Пользователь успешно зарегистрирован")
            } else {
                console.log(response.Error)
            }
        } else {
            console.log("Некорректные данные")
        }
    }
    xhr.send(JSON.stringify(data))
}

if (button) {
    button[0].onclick = processClick(e, "deposit")
    button[1].onclick = processClick(e, "withdraw")
}