let clear_info = () => {
    let blocks = document.querySelectorAll(".info-id")
    blocks.forEach(block => {
        block.innerText = ""
    })
}

let create_button = document.querySelector("#create > button")

if (create_button) {
    create_button.onclick = function (e) {
        let block = document.querySelector("#create .info-id")

        clear_info()

        let xhr = new XMLHttpRequest()
        xhr.open("POST", "/accounts")
        xhr.setRequestHeader("Content-Type", "application/json")
        xhr.onload = function (e) {
            let response = JSON.parse(e.currentTarget.response)
            if ("Error" in response) {
                if (response.Error == null) {
                    console.log("Пользователь успешно зарегистрирован")
                    block.innerText = `Пользователь с ID=${response.ID} успешно зарегистрирован`
                } else {
                    block.innerText = `Возникла ошибка: ${response.Error}`
                }
            } else {
                console.log("Некорректные данные")
            }
        }
        xhr.send(JSON.stringify({"ID": -1}))
    }
}

let button_withdraw = document.querySelector("#withdraw > button")
let button_deposit = document.querySelector("#deposit > button")

if (button_withdraw) {
    button_withdraw.onclick = (e) => {
        let inputs = document.querySelectorAll("#withdraw > input")
        let block = document.querySelector("#withdraw > .info-id")

        clear_info()


        let data = {}

        for (let i = 0; i < inputs.length; i++) {
            data[inputs[i].name] = Number(inputs[i].value)
        }


        let xhr = new XMLHttpRequest()
        xhr.open("POST", `/accounts/${data.ID}/withdraw`)
        xhr.setRequestHeader("Content-Type", "application/json")

        xhr.onload = function (e) {
            let response = JSON.parse(e.currentTarget.response)
            if (response && "Error" in response) {
                if (response.Error == null) {
                    console.log("Операция снятия денег прошла успешно")
                    block.innerText = `У пользователя с ID=${response.ID} сняли ${response.Balance} у.е.`
                } else {
                    block.innerText = `Возникла ошибка: ${response.Error}`
                }
            } else {
                console.log("Некорректные данные")
            }
        }
        xhr.send(JSON.stringify(data))
    }
}

if (button_deposit) {
    button_deposit.onclick = (e) => {
        let inputs = document.querySelectorAll("#deposit > input")
        let block = document.querySelector("#deposit > .info-id")

        clear_info()


        let data = {}

        for (let i = 0; i < inputs.length; i++) {
            data[inputs[i].name] = Number(inputs[i].value)
        }

        let xhr = new XMLHttpRequest()
        xhr.open("POST", `/accounts/${data.ID}/deposit`)
        xhr.setRequestHeader("Content-Type", "application/json")

        xhr.onload = function (e) {
            let response = JSON.parse(e.currentTarget.response)
            if ("Error" in response) {
                if (response.Error == null) {
                    console.log("Операция внесения денег прошла успешно")
                    block.innerText = `Пользователю с ID=${response.ID} внесли ${response.Balance} у.е.`
                } else {
                    block.innerText = `Возникла ошибка: ${response.Error}`
                }
            } else {
                console.log("Некорректные данные")
            }
        }
        xhr.send(JSON.stringify(data))
    }
}

let check_button = document.querySelector("#check > button")
if (check_button) {
    check_button.onclick = function (e) {
        let inputs = document.querySelectorAll("#check > input")
        let block = document.querySelector("#check > .info-id")

        clear_info()


        let data = {}

        for (let i = 0; i < inputs.length; i++) {
            data[inputs[i].name] = Number(inputs[i].value)
        }

        let xhr = new XMLHttpRequest()
        xhr.open("POST", `/accounts/${data.ID}/balance`)
        xhr.setRequestHeader("Content-Type", "application/json")
        xhr.onload = function (e) {
            let response = JSON.parse(e.currentTarget.response)
            if ("Error" in response) {
                if (response.Error == null) {
                    console.log("Операция прошла успешно")
                    block.innerText = `Баланс аккаунта с ID=${response.ID} - ${response.Balance} у.е.`
                } else {
                    block.innerText = `Возникла ошибка: ${response.Error}`
                }
            } else {
                console.log("Некорректные данные")
            }
        }
        xhr.send(JSON.stringify(data))
    }
}
