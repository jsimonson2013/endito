let PAGELINES = []
let PAGENAME = ""
let CTRLDOWN = false
let INTERFACE = "./interface/index.html"

window.onload = () => {
    getAllPages()
    loadListeners()
}

function getAllPages() {
    fetch("http://localhost:3333/pages", {
        method: "GET",
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        }
    })
    .then(res => res.text())
    .then(txt => {
        const files = txt.split(",")
        if (files.length < 1) {
            return
        }

        const loc = document.createElement("select")
        loc.id = "page-location"
        loc.style="font-size: large; width: 80%"

        for (let f of files) {
            if (f == INTERFACE) {
                continue
            }
            const opt = document.createElement("option")
            opt.value = f
            opt.innerHTML = f
            loc.appendChild(opt)
        }

        const ploc = document.getElementById("page-location")
        document.getElementById("load-page-form").replaceChild(loc, ploc)
    })
}

function loadListeners() {
    for(let i of document.querySelectorAll(".edit-text")){
        i.addEventListener("click", (e) => editText(e.target))
    }

    document.addEventListener("dblclick", () => clearActive())
}

function clearActive() {
    for(let i of document.querySelectorAll(".active-edit-text")) {
        i.classList.remove("active-edit-text")
    }
}

function editText(tgt) {
    clearActive()
    tgt.classList.add("active-edit-text")   
}

function loadPage(){
    PAGENAME = document.getElementById("page-location").value
    fetch("http://localhost:3333/load", {
        method: "POST",
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "uri": PAGENAME
        })
    })
    .then(res => res.text())
    .then(txt => {
        document.getElementById("page-editor").innerHTML = txt
        PAGELINES = txt.split("\n")
        loadListeners()
    })
}

function buildHTMLFromPageLines(){
    const innerString = document.getElementById("page-editor").innerHTML
    const lines = innerString.split("\n")
    for (let i = 0; i < lines.length; i++){
        if (lines[i].trim().length > 0) {
            PAGELINES[i] = lines[i]
        }
    }
    return PAGELINES.join("\n")
}

function updatePage(){
    clearActive()

    const HTML = buildHTMLFromPageLines()
    const UNAME = document.getElementById("username").value
    const PWORD = document.getElementById("password").value

    fetch("http://localhost:3333/update", {
        method: "POST",
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "uri": PAGENAME,
            "uname": UNAME,
            "pword": PWORD,
            "content": HTML
        })
    })
}