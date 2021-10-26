$('#desseguir').on('click', pararDeSeguir);
$('#seguir').on('click', seguir);

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