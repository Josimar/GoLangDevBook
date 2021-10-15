$('#login').on('submit', fazerLogin);

function fazerLogin(event){
    event.preventDefault();

    var jqxhr = $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val(),
            password: $('#password').val(),
        }
    }).done(function() {
        window.location = "/home"
    }).fail(function( jqxhr, textStatus, error ) {
        alert("Usuário não encontrado!");
        alert(textStatus);
        alert(error);
    })

    jqxhr.done(function() {
        alert( "first Done" );
    });

    jqxhr.fail(function(jqxhr, textStatus, error) {
        alert( "first fail" );
        alert(textStatus);
        alert(error);
    });

    jqxhr.always(function() {
        alert( "first complete" );
    });
}