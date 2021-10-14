$('#nova-publicacao').on('submit', criarPublicacao)

$(document).on('click', '.curtir-post', curtirPost);
$(document).on('click', '.descurtir-post', descurtirPost);

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
        alert('Error to create post')
    })
}

function curtirPost(event){
    event.preventDefault()

    const elementClick = $(event.target);

    const postId = elementClick.closest('div').data('data-post-id')

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
        alert("Erro ao curtido o Post!")
    }).always(function (){
        elementClick.prop('disabled', false);
    })
}

function descurtirPost(event){
    event.preventDefault()

    const elementClick = $(event.target);

    const postId = elementClick.closest('div').data('data-post-id')

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
        alert("Erro ao curtido o Post!")
    }).always(function (){
        elementClick.prop('disabled', false);
    })
}