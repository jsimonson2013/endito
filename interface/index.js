let PAGELINES = []
let PAGENAME = ""

window.onload = () => {
    loadListeners()
}

function loadListeners() {
    for(let i of document.querySelectorAll(".edit-text")){
        i.addEventListener("dblclick", (e) => {
            editText(e.target)
        })
    }

    document.addEventListener("click", () => {
        document.removeEventListener("keydown", preventDefault, false)
        clearActive()
    })
}

function clearActive() {
    for(let i of document.querySelectorAll(".active-edit-text")) {
        if (i.innerHTML[i.innerHTML.length - 1] == "|") {
            i.innerHTML = i.innerHTML.substring(0, i.innerHTML.length - 1)
        }
        document.removeEventListener("keyup", textKeyHandler, false) 
        i.classList.remove("active-edit-text")
        Window.active = null
    }
}

function editText(tgt) {
    tgt.classList.add("active-edit-text")
    tgt.innerHTML += "|"
    Window.active = tgt
    document.addEventListener("keydown", preventDefault, false)
    document.addEventListener("keyup", textKeyHandler, false)      
}

function preventDefault(e) {
    e.preventDefault()
}

function textKeyHandler(e) {
    if (e.code == "Backspace"){
        Window.active.innerHTML = Window.active.innerHTML.substring(0,Window.active.innerHTML.length - 2) + "|"
    } else if ((e.keyCode <= 90 && e.keyCode >= 48) || 
        (e.keyCode <= 111 && e.keyCode >= 96) ||
        (e.keyCode <= 222 && e.keyCode >= 186) || 
        e.key == " ") {
        Window.active.innerHTML = Window.active.innerHTML.substring(0,Window.active.innerHTML.length - 1) + e.key + "|"
    }
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