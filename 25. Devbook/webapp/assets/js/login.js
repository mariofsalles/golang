$('#login').on('submit', loginUser);

function loginUser(event) {
  event.preventDefault();
  $.ajax({
    url: '/login',
    method: 'POST',
    data: {
      email: $('#email').val(),
      userpass: $('#userpass').val()
    }
  }).done(function () {
    window.location = "/home";
  }).fail(function () {
    Swal.fire({
      icon: 'error',
      title: 'Oops...',
      text: 'Invalid email or password',
    });
  });
}