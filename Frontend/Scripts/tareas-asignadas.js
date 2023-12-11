// function createUser() {
//     var user = {
//         username: $("#username").val(),
//         password: $("#password").val(),
//         email: $("#email").val(),
//         first_name: $("#firstName").val(),
//         last_name: $("#lastName").val(),
//         imagen: $("#image").val()
//     };

//     fetch("http://localhost:8085/api/user/create", {
//         method: "POST",
//         headers: {
//             "Content-Type": "application/json",
//         },
//         body: JSON.stringify(user),
//     })
//     .then(response => {
//         if (!response.ok) {
//             throw new Error("Error al crear usuario.");
//         }
//         return response.json();
//     })
//     .then(data => {
//         $("#form-modal-registrer").html("Usuario creado exitosamente.  ID: " + data.id);
//     })
//     .catch(error => {
//         $("#form-modal-registrer").html("Error al crear usuario. " + error.message);
//     });
// }

// Hacer la solicitud Fetch
fetch('https://localhost:8085/api/task-user/get-all')
.then(response => response.json())
    .then(data => {
        // Manipular el DOM y llenar la tabla
        const tableBody = document.getElementById('table-body');

        data.forEach(group => {
            const row = document.createElement('tr');

            const titleCell = document.createElement('td');
            titleCell.textContent = group.title;

            const descriptionCell = document.createElement('td');
            descriptionCell.textContent = group.description;

            const actionsCell = document.createElement('td');
            actionsCell.innerHTML = `
                <button class="btn btn-primary" type="button" style="margin: 0px;margin-left: 5px;background: rgb(238,153,27);">Editar</button>
                <button class="btn btn-primary" type="button" style="margin: 0px;margin-left: 5px;background: rgb(194,51,51);">Eliminar</button>
            `;

            row.appendChild(titleCell);
            row.appendChild(descriptionCell);
            row.appendChild(actionsCell);

            tableBody.appendChild(row);
        });
    })
    .catch(error => console.error('Error fetching data:', error));