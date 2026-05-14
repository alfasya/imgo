let auth = document.getElementById("auth")
let login = document.getElementById("login")
let register = document.getElementById("register")
let content  = document.getElementById("content")

login.addEventListener("click", async (e) => {
    e.preventDefault()

    content.replaceChildren()

    let loginForm = document.createElement("form")
    content.appendChild(loginForm)

    let usernameInput = document.createElement("input")
    usernameInput.setAttribute("type", "text")
    usernameInput.setAttribute("name", "username")
    usernameInput.setAttribute("placeholder", "Username")
    usernameInput.setAttribute("required", "")
    loginForm.appendChild(usernameInput)

    let passwordInput = document.createElement("input")
    passwordInput.setAttribute("type", "password")
    passwordInput.setAttribute("name", "password")
    passwordInput.setAttribute("placeholder", "Password")
    passwordInput.setAttribute("required", "")
    loginForm.appendChild(passwordInput)

    let loginSubmit = document.createElement("button")
    loginSubmit.setAttribute("type", "submit")
    loginSubmit.textContent = "Login"
    loginForm.appendChild(loginSubmit)

    loginForm.addEventListener("submit", async (e) => {
        e.preventDefault()

        try {
            const formData = new FormData(loginForm)

            const form = {
                username: formData.get("username"),
                password: formData.get("password"),
            }

            let res = await fetch("http://localhost:8080/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(form),
            })

            const data = await res.json()
            console.log(data)

            localStorage.setItem("token", data.token)

            console.log(localStorage)

        } catch(err) {
            console.log(err)
            return
        }
    })
})

register.addEventListener("click", async (e) => {
    e.preventDefault()

    content.replaceChildren()

    let registerForm = document.createElement("form")
    content.appendChild(registerForm)

    let usernameInput = document.createElement("input")
    usernameInput.setAttribute("type", "text")
    usernameInput.setAttribute("placeholder", "Username")
    usernameInput.setAttribute("required", "")
    registerForm.appendChild(usernameInput)

    let passwordInput = document.createElement("input")
    passwordInput.setAttribute("type", "password")
    passwordInput.setAttribute("placeholder", "Password")
    passwordInput.setAttribute("required", "")
    registerForm.appendChild(passwordInput)

    let registerSubmit = document.createElement("button")
    registerSubmit.setAttribute("type", "submit")
    registerSubmit.textContent = "Register"
    registerForm.appendChild(registerSubmit)
})