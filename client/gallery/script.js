let gallery = document.getElementById("gallery")

async function getImages() {
    let token = localStorage.getItem("token")
    try {
        let res = await fetch("http://localhost:8080/gallery", {
            headers: {
                "Authorization": `Bearer ${token}`
            }
        })

        let data = await res.json()

        if (data.ImageList == null) {
            let p = document.createElement("p")
            p.textContent = "Your gallery is empty. Upload some images"
            gallery.appendChild(p)
            return
        }

        let ol = document.createElement("ol")

        for (let i = 0; i < data.ImageList.length; i++) {
            let li = document.createElement("li")
            let img = document.createElement("img")
            img.setAttribute("src", `${data.ImageList[i].Path}`)
            img.setAttribute("width", 300)
            img.setAttribute("alt", data.ImageList[i].Name)

            li.appendChild(img)
            ol.appendChild(li)
            gallery.appendChild(ol)
        }

    } catch(err) {
        console.log(err)
        return
    }
}

getImages()