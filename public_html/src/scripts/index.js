// Switching Light Mode to Dark Mode using document replacement

let switchMode = document.getElementById("switchMode");

switchMode.onclick = function () {
    let theme = document.getElementById("theme");

    if (theme.getAttribute("href") == "styles/light-mode.css") {
        theme.href = "styles/dark-mode.css";
    } else {
        theme.href = "styles/light-mode.css";
    }
}
