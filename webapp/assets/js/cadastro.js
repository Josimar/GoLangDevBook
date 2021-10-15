$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(event){
    event.preventDefault();

    let senha = $('#password').val()
    let confirmarSenha = $('#confirm-password').val()

    if (senha !== confirmarSenha){
        alert("As senhas não coincidem")
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            name: $('#name').val(),
            email: $('#email').val(),
            password: $('#password').val(),
        }
    }).done(function(){
        alert("Usuário cadastrado com sucesso!");
    }).fail(function(){
        alert("Erro as cadastrar usuário!");
    })
}