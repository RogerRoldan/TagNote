function createUser() {
    var user = {
        username: $("#username").val(),
        password: $("#password").val(),
        email: $("#email").val(),
        first_name: $("#firstName").val(),
        last_name: $("#lastName").val(),
        imagen: $("#image").val()
    };

    fetch("http://localhost:8085/api/user/create", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(user),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error("Error al crear usuario.");
        }
        return response.json();
    })
    .then(data => {
        $("#form-modal-registrer").html("Usuario creado exitosamente.  ID: " + data.id);
    })
    .catch(error => {
        $("#form-modal-registrer").html("Error al crear usuario. " + error.message);
    });
}
