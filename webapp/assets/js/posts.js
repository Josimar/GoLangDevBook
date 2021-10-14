$('#nova-publicacao').on('submit', criarPublicacao)
$('#curtir-post').on('submit', curtirPost)

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
        url: `/posts/${postId}/curtir`,
        method: "POST"
    }).done(function (){
        const contadorCurtida = elementClick.next('span')
        const quantityCurtida = parseInt(contadorCurtida.text());

        contadorCurtida.text(quantityCurtida + 1);
    }).fail(function (){
        alert("Erro ao curtido o Post!")
    }).always(function (){
        elementClick.prop('disabled', false);
    })
}