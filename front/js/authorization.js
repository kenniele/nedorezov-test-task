let form = document.getElementById("create")
let button = form.querySelectorAll("button")[0]

if (button) {
    button.onclick = function (e) {
        let block = form.querySelector(".info-id")

        let xhr = new XMLHttpRequest()
        xhr.open("POST", "/accounts")
        xhr.onload = function (e) {
            console.log(e.currentTarget.response)
            let response = JSON.parse(e.currentTarget.response)
            if ("Error" in response) {
                if (response.Error == null) {
                    console.log("Пользователь успешно зарегистрирован")
                    block.innerText = `Пользователь с ID=${response.ID} успешно зарегистрирован`
                } else {
                    console.log(response.Error)
                }
            } else {
                console.log("Некорректные данные")
            }
        }
        xhr.send(JSON.stringify({}))
    }
}