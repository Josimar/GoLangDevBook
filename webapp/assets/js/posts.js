$('#nova-publicacao').on('submit', criarPublicacao)

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
    }).error(function(){
        alert('Error to create post')
    })
}