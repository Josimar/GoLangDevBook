$('#login').on('submit', fazerLogin);

function fazerLogin(event){
    event.preventDefault();

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val(),
            password: $('#password').val(),
        }
    }).done(function(){
        window.location = "/home"
    }).error(function(erro){
        alert("Usuário não encontrado!");
    })
}