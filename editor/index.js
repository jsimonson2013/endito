let PAGELINES = []
let PAGENAME = ""
let CTRLDOWN = false


// ------------------------- INIT LOGIC ------------------------------------ //

/*
    onload initial event listeners are attached

    getAllPages builds the select by requesting all editable pages from the 
    endito tool
*/

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
    .then(txt => buildPageSelect(txt))
}

// ------------------------- END INIT -------------------------------------- //


// ------------------------- EVENT HANDLERS -------------------------------- //

/*
    loadPage is responsible for handling requests to load an html page from
    the endito tool given the selected option

    updatePage sends a request to save the page with the edits applied in the
    editor to the endito tool
*/

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
    .then(txt => populateEditor(txt))
}

function updatePage(){
    clearActive()
    clearContentEditable()

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
    }).then(setEditable())
}

// ------------------------- END HANDLERS ---------------------------------- //


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

function buildPageSelect(txt) {
    const files = txt.split(",") // files come through as stringified array
    if (files.length < 1) {
        return
    }

    const loc = document.createElement("select")
    loc.id = "page-location"
    loc.style="font-size: large; width: 80%"

    for (let f of files) { // create option for each file
        const opt = document.createElement("option")
        opt.value = f
        opt.innerHTML = f
        loc.appendChild(opt)
    }

    const ploc = document.getElementById("page-location")
    document.getElementById("load-page-form").replaceChild(loc, ploc)
}


function clearActive() {
    for(let i of document.querySelectorAll(".active-edit-text")) {
        i.classList.remove("active-edit-text")
    }
}

function clearContentEditable() {
    for(let i of document.querySelectorAll(".edit-text")){
        i.removeAttribute("contenteditable", true)
    }
}

function editText(tgt) {
    clearActive()
    tgt.classList.add("active-edit-text")   
}

function loadListeners() {
    for(let i of document.querySelectorAll(".edit-text")){
        i.addEventListener("click", (e) => editText(e.target))
    }

    document.addEventListener("dblclick", () => clearActive())
}

function populateEditor(txt) {
    document.getElementById("page-editor").innerHTML = txt
    PAGELINES = txt.split("\n")
    setEditable()
    loadListeners()
}

function setEditable() {
    for(let i of document.querySelectorAll(".edit-text")){
        i.setAttribute("contenteditable", true)
    }
}