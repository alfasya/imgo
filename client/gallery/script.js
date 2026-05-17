let gallery = document.getElementById("gallery")
let body = document.getElementById("body")

let text = document.createElement("p")
gallery.appendChild(text)

let token = localStorage.getItem("token")
let username = localStorage.getItem("username")

if (token != null || token != null) {
    const loginRegister = document.getElementById("login-register")
    loginRegister.innerHTML = `<p id="welcome">Welcome, ${username}</p>`
}

async function getImages() {
    let data
    try {
        let res = await fetch("http://localhost:8080/gallery", {
            headers: {
                "Authorization": `Bearer ${token}`
            }
        })

        if (res.status == 401) {
            window.location.replace("http://localhost:5500/client/auth")
        }

        data = await res.json()

        console.log(data)

        if (data.ImageList == null) {
            let p = document.createElement("p")
            p.textContent = "Your gallery is empty. Upload some images"
            gallery.appendChild(p)
            return
        }

        let ol = document.createElement("ol")
        ol.setAttribute("id", "image-list")

        for (let i = 0; i < data.ImageList.length; i++) {
            let li = document.createElement("li")
            li.setAttribute("class", "list")
            ol.appendChild(li)

            let a = document.createElement("a")
            a.setAttribute("href", `http://localhost:8080/${data.Links[i]}`)
            a.setAttribute("target", "_blank")
            li.appendChild(a)

            let img = document.createElement("img")
            img.setAttribute("class", "image")
            img.setAttribute("width", 300)
            img.setAttribute("src", `http://localhost:8080/${data.Links[i]}`)
            img.setAttribute("alt", data.ImageList[i].Name)
            a.appendChild(img)

            let delBtn = document.createElement("button")
            delBtn.textContent = "del"
            delBtn.setAttribute("class", "delete-button")
            delBtn.setAttribute("type", "button")
            delBtn.setAttribute("id", `${data.Links[i]}`)
            li.appendChild(delBtn)
            li.prepend(delBtn)

            gallery.appendChild(ol)
        }

        //HANDLING DELETE ACTION
        //create element for list, image, option panel, and delete button
        let list = document.querySelectorAll(".list")
        let delBtn = document.querySelectorAll(".delete-button")

        //ADD EVENT addEventListener
        delBtn.forEach(btn => {
            btn.addEventListener("click", async (e) => {
                text.textContent = ""
                let res = await fetch(`http://localhost:8080/${btn.id}`, {
                    method: "DELETE",
                    headers: {
                        "Authorization": `Bearer ${token}`,
                    }
                })

                let msg = await res.json()

                btn.closest(".list").remove()

                text.setAttribute("class", "notif")
                text.textContent = "File deleted"
            })
        });
    } catch (err) {
        console.log(err)
        return
    }
}

getImages()