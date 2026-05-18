//NEW
let login = document.getElementById("login")
let register = document.getElementById("register")
let content  = document.getElementById("content")
let authForm = document.getElementById("auth-form")
let loginSubmit = document.getElementById("auth-submit-login")
let regisSubmit = document.getElementById("auth-submit-register")
let text = document.createElement("p")

const token = localStorage.getItem("token")

if (token != null) {
    window.location.replace("http://localhost:5500/client/gallery")
}

text.setAttribute("class", "regis-info")
authForm.appendChild(text)
authForm.prepend(text)
text.textContent = "You have to log in to access your gallery or upload images"

login.addEventListener("click", () => {
    regisSubmit.style.display = "none"
    loginSubmit.style.display = "block"
    login.setAttribute("class", "auth-buttons login-active")
    authForm.style.setProperty("border-top", "3px solid #3885cd")
    register.setAttribute("class", "auth-buttons")

})
register.addEventListener("click", () => {
    regisSubmit.style.display = "block"
    loginSubmit.style.display = "none"
    register.setAttribute("class", "auth-buttons register-active")
    authForm.style.setProperty("border-top", "3px solid #cd4438")
    login.setAttribute("class", "auth-buttons")
})

let state = "login"

const response = {
    ok: "Logged in",
    notOk: "Wrong username or password",
}

loginSubmit.addEventListener("click", () => {
    state = "login"
    ok = "Logged in"
    notOk = "Wrong username or password"
})

regisSubmit.addEventListener("click", () => {
    state = "register"
    response.ok = "Registered Successfully"
    response.notOk = "Username already exists"
})

authForm.addEventListener("submit", async (e) => {
    e.preventDefault()

    let uri

    if (state === "login") {
        uri = "http://localhost:8080/login"
    } else if (state === "register") {
        uri = "http://localhost:8080/register"
    } else {
        console.log("error: uri not found")
        return
    }

    const formData = new FormData(authForm)

    const form = {
        username: formData.get("username"),
        password: formData.get("password"),
    }

    const res = await fetch(uri, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(form)
    })

    console.log(form)
    console.log(res)

    if (!res.ok) {
        console.log(response.notOk)
        
        authForm.appendChild(text)
        authForm.prepend(text)
        text.textContent = response.notOk
        return     
    }

    const data = await res.json()

    if (state === "login") {
        localStorage.setItem("token", data.token)
        localStorage.setItem("username", data.username)     
    }

    text.setAttribute("class", "regis-info")
    authForm.appendChild(text)
    authForm.prepend(text)
    text.textContent = response.ok

    if (state === "login") {
        window.location.replace("http://localhost:5500/client/gallery")
    } else if (state === "regis") {
        location.reload()
    }
})