function toggle_theme(darkMode) {
    const html = document.getElementById("root");
    if (!darkMode) {
        html.classList.add("dark");
    } else {
        html.classList.remove("dark");
    }
}
