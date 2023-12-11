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