function confirmSignOut() {
    // Mostrar el modal de confirmación con SweetAlert2
    Swal.fire({
        title: '¿Seguro de que desea salir?',
        text: '¡esta a punto de salir de la sesión!',
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#d33',
        confirmButtonText: 'Sí, salir'
    }).then((result) => {
        if (result.isConfirmed) {
            localStorage.setItem('idUser', "");
            window.location.href = 'index.html';
        }
    });
}

function signOut(){

}