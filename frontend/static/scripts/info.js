function searchSuggestions() {
    const query = document.getElementById("search-bar").value;
    const suggestionsList = document.getElementById("suggestions-list");

    if (query.length === 0) {
        suggestionsList.innerHTML = '';
        return;
    }

    fetch(`/search?q=${query}`)
        .then(response => response.json())
        .then(data => {
            suggestionsList.innerHTML = '';
            data.forEach(suggestion => {
                const li = document.createElement("li");
                li.innerHTML = `<a href="${suggestion.url}">${suggestion.label} (${suggestion.type})</a>`;
                suggestionsList.appendChild(li);
            });
        })
        .catch(error => console.error('Error fetching suggestions:', error));
}

const themeToggle = document.getElementById('theme-toggle');

function toggleTheme() {
    if (themeToggle.checked) {
        document.documentElement.setAttribute('data-theme', 'dark');
        localStorage.setItem('theme', 'dark');
    } else {
        document.documentElement.setAttribute('data-theme', 'light');
        localStorage.setItem('theme', 'light');
    }
}

const savedTheme = localStorage.getItem('theme');
if (savedTheme === 'dark') {
    document.documentElement.setAttribute('data-theme', 'dark');
    themeToggle.checked = true;
} else {
    document.documentElement.setAttribute('data-theme', 'light');
    themeToggle.checked = false;
}

themeToggle.addEventListener('change', toggleTheme);
