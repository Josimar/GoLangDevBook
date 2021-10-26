$('#desseguir').on('click', pararDeSeguir);
$('#seguir').on('click', seguir);
$('#editar-usuario').on('submit', editar);
$('#atualizar-senha').on('submit', atualizarSenha);

function pararDeSeguir(){
    const usuarioId = $(this).data('usuario-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuario/${usuarioId}/unfollow`,
        method: "POST"
    }).done(function(){
        window.location = `/usuario/${usuarioId}`;
    }).fail(function(){
        Swal.fire("Ops...", "Erro ao parar de seguir o usuário!", "error");
        $(this).prop('disabled', false);
    })
}

function seguir(){
    const usuarioId = $(this).data('usuario-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuario/${usuarioId}/follow`,
        method: "POST"
    }).done(function(){
        window.location = `/usuario/${usuarioId}`;
    }).fail(function(){
        Swal.fire("Ops...", "Erro ao seguir o usuário!", "error");
        $(this).prop('disabled', false);
    })
}

function editar(event){
    event.preventDefault();

    $.ajax({
       url: "/editar-usuario",
       method: "POST",
       data: {
           name: $('#name').val(),
           email: $('#email').val(),
       }
    }).done(function(){
        Swal.fire("Sucesso!", "Usuário atualizado com sucesso", "success")
            .then(function(){
               window.location = "/perfil";
            });
    }).fail(function(){
        Swal.fire("ops...", "Erro ao atualizar dados do usuário", "error");
    });

}

function atualizarSenha(event){
    event.preventDefault();

    if ($('#senha-nova').val() !== $('#confirmar-senha').val()){
        Swal.fire("Ops...", "As senhas não coincidem", "warning");
        return;
    }

    $.ajax({
       url: "/atualizar-senha",
       method: "POST",
       data: {
           atual: $('#senha-atual').val(),
           nova: $('#senha-nova').val(),
       }
    }).done(function(){
        Swal.fire("Sucesso", "Senha alterada com sucesso", "success")
            .then(function(){
                window.location = "/perfil";
            });
    }).fail(function(){
        Swal.fire("Ops...", "Erro ao atualizar a senha", "error");
    });
}