function openNav() {
    document.getElementById("navbar").style.width = "250px";
}

function closeNav() {
    document.getElementById("navbar").style.width = "0";
}

var fechaActual = new Date();

var opcionesDeFormato = { weekday: 'long', day: 'numeric' , month: 'long', year: 'numeric'};
var fechaFormateada = fechaActual.toLocaleDateString('es-ES', opcionesDeFormato);

document.getElementById('fecha-actual').textContent = fechaFormateada;

var form = document.getElementById('form');
    var input = document.getElementById('input');
    var todosUL = document.getElementById('todos');

    const todos = JSON.parse(localStorage.getItem('todos')) || [];

    todos.forEach(todo => addTodoElement(todo));

    form.addEventListener('submit', (e) => {
        e.preventDefault();
        addTodo();
    });

    function addTodo() {
        let todoText = input.value;

        if (todoText) {
            const todo = {
                text: todoText,
                completed: false
            };

            todos.push(todo);
            addTodoElement(todo);

            input.value = '';
            updateLS();
        }
    }

    function addTodoElement(todo) {
        const todoEl = document.createElement('div');
        todoEl.classList.add('faq');
        if (todo.completed) {
            todoEl.classList.add('completed');
        }

        const titleEl = document.createElement('h3');
        titleEl.classList.add('faq-title');
        titleEl.textContent = todo.text;

        const toggleBtn = document.createElement('button');
        toggleBtn.classList.add('faq-toggle');
        toggleBtn.innerHTML = '<i class="fas fa-chevron-down"></i><i class="fas fa-times"></i>';
        toggleBtn.addEventListener('click', () => {
            todoEl.classList.toggle('completed');
            updateLS();
        });

        todoEl.appendChild(titleEl);
        todoEl.appendChild(toggleBtn);

        todoEl.addEventListener('contextmenu', (e) => {
            e.preventDefault();
            todoEl.remove();
            updateLS();
        });

        todosUL.appendChild(todoEl);
    }

    function updateLS() {
        localStorage.setItem('todos', JSON.stringify(todos));
    }

    const toggles = document.querySelectorAll('.faq-toggle');
    toggles.forEach(toggle => {
        toggle.addEventListener('click', () => {
            toggle.parentNode.classList.toggle('completed');
        });
    });