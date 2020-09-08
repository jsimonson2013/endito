window.onload = () => {
    for(let i of document.querySelectorAll(".edit-text")){
        i.addEventListener("dblclick", (e) => {
            editText(e.target)
        })
    }
}

function editText(tgt) {
    tgt.classList.add("active-edit-text")
    tgt.innerHTML += "|"
    document.addEventListener("keyup", (e) => {
        console.log(e.key)
        if (e.code == "Backspace"){
            tgt.innerHTML = tgt.innerHTML.substring(0, tgt.innerHTML.length - 2) + "|"
        } else if ((e.keyCode <= 90 && e.keyCode >= 48) || 
            (e.keyCode <= 111 && e.keyCode >= 96) ||
            (e.keyCode <= 222 && e.keyCode >= 186) || 
            e.key == " ") {
            tgt.innerHTML = tgt.innerHTML.substring(0, tgt.innerHTML.length - 1) + e.key + "|"
        }
      })
}