const button = document.getElementById("post-btn");

button.addEventListener("click", async _ => {
    axios({
        method: "post",
        url: "http://localhost:5000/create",
        headers: {
            Authorization: "PASSWORD"
        },
        data: {
            url: document.getElementById("websiteURL").value,
            slug: document.getElementById("slug").value
        }
    });
})