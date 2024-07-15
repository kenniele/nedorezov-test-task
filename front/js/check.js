let form2 = document.getElementById("check")
let button = form.querySelectorAll("button")[0]

if (button) {
    button.onclick = function (e) {
        let inputs = form2.querySelectorAll(".form > input")

        let data = {}

        for (let i = 0; i < inputs.length; i++) {
            data[inputs[i].name] = inputs[i].value
        }

        let xhr = new XMLHttpRequest()
        xhr.open("POST", "/accounts")
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
}