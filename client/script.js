let form = document.getElementById("form")
let images = document.getElementById("images")
let formContainer = document.getElementById("form-container")
const response = document.createElement("p")
let fileInfo = document.getElementsByClassName("file-info")[0]

const token = localStorage.getItem("token")
let username = localStorage.getItem("username")

if (token != null || username != null) {
    const loginRegister = document.getElementById("login-register")
    loginRegister.innerHTML = `<p id="welcome">Welcome, ${username}</p>`
}

form.addEventListener("submit", async (e) => {
    e.preventDefault()

    if (token == "" || token == null) {
        window.location.replace("http://localhost:5500/client/gallery")
    }

    const formData = new FormData()

    for (let i = 0; i < images.files.length; i++) {
        formData.append("images", images.files[i])
    }

    try {
        let res = await fetch("http://localhost:8080/upload", {
            method: "POST",
            headers: {
                Authorization: `Bearer ${token}`
            },
            body: formData,
        })

        response.replaceChildren()

        let data = await res.json()

        response.textContent = `${data.message}. Check your Gallery`
        response.setAttribute("class", "notif")
        formContainer.appendChild(response)
        form.before(response)

        fileInfo.textContent = ""

    } catch (err) {
        console.log(err)
        return
    }
})

images.addEventListener("change", (e) => {
    e.preventDefault()

    let n = e.target.files.length
    fileInfo.textContent = `${n} files selected`
    console.log(fileInfo)
})