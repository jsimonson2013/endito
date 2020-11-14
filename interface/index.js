let PAGELINES = []
let PAGENAME = ""
let CTRLDOWN = false

window.onload = () => {
    loadListeners()
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