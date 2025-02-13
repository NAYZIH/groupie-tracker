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

document.getElementById('search-bar').addEventListener('focus', function() {
    document.getElementById('suggestions-list').style.display = 'block';
});

document.getElementById('search-bar').addEventListener('blur', function() {
    setTimeout(() => {
        document.getElementById('suggestions-list').style.display = 'none';
    }, 200);
});
