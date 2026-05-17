let auth = document.getElementById("auth")
let login = document.getElementById("login")
let register = document.getElementById("register")
let content  = document.getElementById("content")
let authForm = document.getElementById("auth-form")
let authSubmit = document.getElementById("auth-submit")
let text = document.createElement("p")

const token = localStorage.getItem("token")

if (token != null) {
    window.location.replace("http://localhost:5500/client/gallery")
}

let state = 0;

text.setAttribute("class", "regis-info")
authForm.appendChild(text)
authForm.prepend(text)
text.textContent = "You have to log in to access your gallery or upload images"

login.addEventListener("click", async (e) => {
    e.preventDefault()

    state = 0

    login.setAttribute("class", "auth-buttons login-active")
    authForm.style.setProperty("border-top", "3px solid #3885cd")
    register.classList.remove("register-active")
    authSubmit.textContent = "Login"
})

authForm.addEventListener("submit", async (e) => {
    e.preventDefault()

    try {
        const formData = new FormData(authForm)

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

        if (res.status != 200) {
            console.log("wrong username or password")
            
            authForm.appendChild(text)
            authForm.prepend(text)
            text.textContent = "Wrong username or password"
            return
        }

        const data = await res.json()
        console.log(data)

        localStorage.setItem("token", data.token)
        localStorage.setItem("username", data.username)

        text.setAttribute("class", "regis-info")
        authForm.appendChild(text)
        authForm.prepend(text)
        text.textContent = "Logged in"

        window.location.replace("http://localhost:5500/client/gallery")

    } catch(err) {
        console.log(err)
        return
    }
})

register.addEventListener("click", async (e) => {
    e.preventDefault()

    register.setAttribute("class", "auth-buttons register-active")
    authForm.style.setProperty("border-top", "3px solid #cd4438")
    login.classList.remove("login-active")
    authSubmit.textContent = "Register"

    authForm.addEventListener("submit", async (e) => {
        e.preventDefault()

        text.remove()

        try {
            const formData = new FormData(authForm)

            const form = {
                username: formData.get("username"),
                password: formData.get("password"),
            }

            let res = await fetch("http://localhost:8080/register", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(form),
            })

            if (res.status != 201) {
                console.log("username already exists")

                text.setAttribute("class", "regis-info")
                authForm.appendChild(text)
                authForm.prepend(text)
                text.textContent = "Username already exists"
                return
            }

            const data = await res.json()
            console.log(data)
            text.setAttribute("class", "regis-info")
            authForm.appendChild(text)
            authForm.prepend(text)
            text.textContent = "Registered successfully"

            location.reload()
        } catch(err) {
            console.log(err)
            return
        }
    })
})