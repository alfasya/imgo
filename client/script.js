let form = document.getElementById("form")
let images = document.getElementById("images")
let formContainer = document.getElementById("form-container")
const response = document.createElement("p")

form.addEventListener("submit", async (e) => {
    e.preventDefault()

    const token = localStorage.getItem("token")
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

        response.textContent = data.message
        formContainer.appendChild(response)

    } catch (err) {
        console.log(err)
        return
    }
})