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
        Swal.fire("Ops..", "Usuário não encontrado!", "error");
    })

    /*
    jqxhr.done(function() {
        alert( "first Done" );
    });

    jqxhr.fail(function(jqxhr, textStatus, error) {
        alert( "first fail" );
    });

    jqxhr.always(function() {
        alert( "first complete" );
    });
    */
}