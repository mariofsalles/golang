$('#unfollow').on('click', unfollowUser);
$('#follow').on('click', followUser);
$("#update-profile").on("submit", updateUser);
$("#update-password").on("submit", updatePassword);
$('#delete-user').on('click', deleteUser);


function unfollowUser() {
  const userId = $(this).data('user-id');
  $(this).prop('disabled', true);

  $.ajax({
    url: `/users/${userId}/unfollow`,
    method: 'POST',
  }).done(function () {
    location.reload();
  }).fail(function () {
    Swal.fire({
      icon: 'error',
      title: 'Oops...',
      text: 'Something went wrong!',
    });
    $(this).prop('disabled', false);
  });
}

function followUser() {
  const userId = $(this).data('user-id');
  $(this).prop('disabled', true);

  $.ajax({
    url: `/users/${userId}/follow`,
    method: 'POST',
  }).done(function () {
    location.reload();
  }).fail(function () {
    Swal.fire({
      icon: 'error',
      title: 'Oops...',
      text: 'Something went wrong!',
    });
    $(this).prop('disabled', false);
  });
}

function updateUser(event) {
  event.preventDefault();

  $.ajax({
    url: "/update-profile",
    method: "PUT",
    data: {
      username: $("#current-user").val(),
      email: $("#email").val(),
      nick: $("#nick").val(),
    },
  }).done(function () {
    Swal.fire({
      icon: "success",
      title: "Success",
      text: "Profile updated!",
    }).then(function () {
      window.location = "/profile";
    })
  }).fail(function () {
    Swal.fire({
      icon: "error",
      title: "Oops...",
      text: "Something went wrong!",
    });
  });
}

function updatePassword(event) {
  event.preventDefault();

  if ($("#new-password").val() != $("#password-confirmation").val()) {
    Swal.fire({
      icon: "error",
      title: "Oops...",
      text: "Passwords do not match!",
    });
    return;
  }

  $.ajax({
    url: "/update-password",
    method: "POST",
    data: {
      current_password: $("#current-password").val(),
      new_password: $("#new-password").val(),
    },
  }).done(function () {
    Swal.fire({
      icon: "success",
      title: "Success",
      text: "Password updated!",
    }).then(function () {
      window.location = "/profile";
    })
  }).fail(function () {
    Swal.fire({
      icon: "error",
      title: "Oops...",
      text: "Something went wrong!",
    });
  });
}

function deleteUser() {
  Swal.fire({
    title: "Are you sure?",
    text: "You won't be able to revert this!",
    icon: "warning",
    showCancelButton: true,
    confirmButtonColor: "#3085d6",
    cancelButtonColor: "#d33",
    confirmButtonText: "Yes, delete it!",
    cancelButtonText: "No, cancel it.",
  }).then((result) => {
    if (result.isConfirmed) {
      $.ajax({
        url: "/delete-user",
        method: "DELETE",
      }).done(function () {
        Swal.fire({
          icon: "success",
          title: "Success",
          text: "User deleted!",
        }).then(function () {
          window.location = "/";
        })
      }).fail(function () {
        Swal.fire({
          icon: "error",
          title: "Oops...",
          text: "Something went wrong!",
        });
      });
    }
  });
}