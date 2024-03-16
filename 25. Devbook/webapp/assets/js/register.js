$('#register-form').on('submit', createUser);

function createUser(event) {
  event.preventDefault();
  if ($('#userpass').val() != $('#userpass-confirm').val()) {
    Swal.fire({
      icon: 'error',
      title: 'Oops...',
      text: 'Password and confirm password are different',
    });
    return;
  }
  $.ajax({
    url: '/users',
    method: 'POST',
    data: {
      username: $('#username').val(),
      email: $('#email').val(),
      nick: $('#nick').val(),
      userpass: $('#userpass').val()
    }
  }).done(function () {
    Swal.fire({
      icon: 'success',
      title: 'Success!',
      text: 'User created with success!',
    }).then(function () {
      $.ajax({
        url: '/login',
        method: 'POST',
        data: {
          email: $('#email').val(),
          userpass: $('#userpass').val()
        }
      }).done(function () {
        window.location = '/home';
      }).fail(function () {
        Swal.fire({
          icon: 'error',
          title: 'Oops...',
          text: 'Error to login',
        });
      });
    });
  }).fail(function () {
    Swal.fire({
      icon: 'error',
      title: 'Oops...',
      text: 'Error creating user',
    });
  });
}