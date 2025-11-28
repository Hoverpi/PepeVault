const URL = "http://localhost:8080"

async function loginUser() {
    const url = URL + "/login";
    try {
        const uuid_v4 = document.getElementById("uuid_v4");
        const password = document.getElementById("password");
        const payload = {"UUID_v4": uuid_v4, "Password": password};

        fetch(url, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(payload)
        })
        .then(response => response.json())
        .then(result => console.log(result));
    } catch (error) {
        console.error(error.message);
    }
}

async function registerUser() {
    const url = URL + "/login";
    try {
        const username = document.getElementById("username");
        const password = document.getElementById("password");
        const confirm_password = document.getElementById("confirm-password");
        const payload = {"username": username, "password": password, "confirm-Password": confirm_password};

        fetch(url, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(payload)
        })
        .then(response => response.json())
        .then(result => console.log(result));
    } catch (error) {
        console.error(error.message);
    }
}