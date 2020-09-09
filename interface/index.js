window.onload = () => {
    for(let i of document.querySelectorAll(".edit-text")){
        i.addEventListener("dblclick", (e) => {
            editText(e.target)
        })
    }

    document.addEventListener("click", () => {
        clearActive()
    })
}

function clearActive() {
    for(let i of document.querySelectorAll(".active-edit-text")) {
        if (i.innerHTML[i.innerHTML.length - 1] == "|") {
            i.innerHTML = i.innerHTML.substring(0, i.innerHTML.length - 1)
        }
        i.removeEventListener("keyup", textKeyHandler, false)
        i.classList.remove("active-edit-text")
        Window.active = null
    }
}

function editText(tgt) {
    tgt.classList.add("active-edit-text")
    tgt.innerHTML += "|"
    Window.active = tgt
    document.addEventListener("keyup", textKeyHandler, false)      
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