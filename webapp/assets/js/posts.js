$('#nova-publicacao').on('submit', criarPublicacao);

$(document).on('click', '.curtir-post', curtirPost);
$(document).on('click', '.descurtir-post', descurtirPost);

$('#atualizar-post').on('click', atualizarPost);
$('.deletar-post').on('click', deletarPost);

// $('#curtir-post').on('submit', curtirPost)

function criarPublicacao(event){
    event.preventDefault()

    $.ajax({
        url: "/posts",
        method: "POST",
        data: {
            title: $('#title').val(),
            content: $('#content').val(),
            description: $('#description').val(),
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function(){
        Swal.fire('Error!', 'Error to create post', 'error')
    });
}

function curtirPost(event){
    event.preventDefault()

    const elementClick = $(event.target);

    const postId = elementClick.closest('div').data('post-id')

    elementClick.prop('disabled', true);

    $.ajax({
        url: `/posts/${postId}/like`,
        method: "POST"
    }).done(function (){
        const contadorCurtida = elementClick.next('span')
        const quantityCurtida = parseInt(contadorCurtida.text());

        contadorCurtida.text(quantityCurtida + 1);

        elementClick.addClass('descurtir-post');
        elementClick.addClass('text-danger');
        elementClick.removeClass('curtir-post');
    }).fail(function (){
        Swal.fire('Error!', 'Error to like post', 'error')
    }).always(function (){
        elementClick.prop('disabled', false);
    });
}

function descurtirPost(event){
    event.preventDefault()

    const elementClick = $(event.target);

    const postId = elementClick.closest('div').data('post-id')

    elementClick.prop('disabled', true);

    $.ajax({
        url: `/posts/${postId}/dislike`,
        method: "POST"
    }).done(function (){
        const contadorCurtida = elementClick.next('span')
        const quantityCurtida = parseInt(contadorCurtida.text());

        contadorCurtida.text(quantityCurtida - 1);

        elementClick.removeClass('descurtir-post');
        elementClick.removeClass('text-danger');
        elementClick.addClass('curtir-post');
    }).fail(function (){
        Swal.fire('Error!', 'Error to dislike a post', 'error')
    }).always(function (){
        elementClick.prop('disabled', false);
    });
}

function atualizarPost(event){
    $(this).prop('disabled', true);

    const postId = $(this).data('post-id');

    $.ajax({
        url: `/posts/${postId}`,
        method: "PUT",
        data: {
            title: $('#title').val(),
            content: $('#content').val(),
            description: $('#description').val(),
        }
    }).done(function(){
        Swal.fire("Sucesso!", "Post editada com sucesso", "success")
            .then(function(){
               window.location = "/home";
            });
    }).fail(function(){
        Swal.fire('Error!', 'Error to edit a post', 'error')
    }).always(function(){
        $('#atualizar-post').prop('disabled', true);
    });
}

function deletarPost(event){
    event.preventDefault()

    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja excluir? essa ação é irreversível.",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirm){
        if (!confirm.value) return;

        const elementClick = $(event.target);
        const post = elementClick.closest('div');
        const postId = post.data('post-id')

        elementClick.prop('disabled', true);

        $.ajax({
            url: `/posts/${postId}`,
            method: "DELETE",
        }).done(function(){
            post.fadeOut("slow", function(){
                $(this).remove();
            });
        }).fail(function(){
            Swal.fire('Error!', 'Error to delete a post', 'error')
        });
    });
}