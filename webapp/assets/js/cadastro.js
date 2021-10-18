$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(event){
    event.preventDefault();

    let senha = $('#password').val()
    let confirmarSenha = $('#confirm-password').val()

    if (senha !== confirmarSenha){
        Swal.fire("Ops..", "As senhas não coincidem", "error");
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
        Swal.fire("Sucesso!", "Usuário cadastrado com sucesso!", "success")
            .then(function(){
                $.ajax({
                    url: "/login",
                    method: "POST",
                    data: {
                        email: $('#email').val(),
                        password: $('#password').val()
                    }
                }).done(function(){
                    window.location = "/home";
                }).fail(function(){
                    Swal.fire("Erro!", "Erro ao autenticar usuário!", "error");
                })
            });
    }).fail(function(){
        Swal.fire("Erro!", "Erro ao cadastrar usuário!", "error");
    })
}