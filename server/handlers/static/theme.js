function toggle_theme() {
    const html = document.getElementById("root");
    if (localStorage.getItem('theme') === 'light') {
        html.classList.add("dark");
        localStorage.setItem('theme', 'dark');
        //htmx.ajax('PUT', '/theme', 'theme-toggle-icon');
    } else {
        html.classList.remove("dark");
        localStorage.setItem('theme', 'light');
    }
}
