$("#new-post").on("submit", createPost);
$(document).on("click", ".like-post", likePost)
$(document).on("click", ".unlike-post", unlikePost)
$("#update-post").on("click", updatePost);
$(".delete-post").on("click", deletePost);

function createPost(event) {
  event.preventDefault();

  $.ajax({
    url: "/posts",
    method: "POST",
    data: {
      title: $("#title").val(),
      content: $("#content").val(),
    }
  }).done(function () {
    window.location = "/home";
  }).fail(function () {
    Swal.fire({
      icon: 'error',
      title: 'Oops...',
      text: 'Error to create post',
    });
  });
}

function likePost(event) {
  event.preventDefault();
  const clickedHeart = $(event.target);
  const postId = clickedHeart.closest('div').data('post-id');

  clickedHeart.prop('disabled', true);
  $.ajax({
    url: `/posts/${postId}/like`,
    method: "POST",
  }).done(function () {
    const counterLikes = clickedHeart.next('span');
    const quantityLikes = parseInt(counterLikes.text());
    counterLikes.text(quantityLikes + 1);
    clickedHeart.addClass('unlike-post').addClass('text-danger').removeClass('like-post');
  }).fail(function () {
    Swal.fire({
      icon: 'error',
      title: 'Oops...',
      text: 'Error to like post',
    });
  }).always(function () { clickedHeart.prop('disabled', false) });
}

function unlikePost(event) {
  event.preventDefault();
  const clickedHeart = $(event.target);
  const postId = clickedHeart.closest('div').data('post-id');

  clickedHeart.prop('disabled', true);
  $.ajax({
    url: `/posts/${postId}/unlike`,
    method: "POST",
  }).done(function () {
    const counterLikes = clickedHeart.next('span');
    const quantityLikes = parseInt(counterLikes.text());
    counterLikes.text(quantityLikes - 1);
    clickedHeart.removeClass('unlike-post').removeClass('text-danger').addClass('like-post');

  }).fail(function () {
    Swal.fire({
      icon: 'error',
      title: 'Oops...',
      text: 'Error to like post',
    });
  }).always(function () { clickedHeart.prop('disabled', false) });
}

function updatePost() {
  $(this).prop('disabled', true);
  const postId = $(this).data('post-id');
  console.log(postId);

  $.ajax({
    url: `/posts/${postId}`,
    method: "PUT",
    data: {
      title: $("#title").val(),
      content: $("#content").val(),
    }
  }).done(function () {
    Swal.fire({
      icon: 'success',
      title: 'Success!',
      text: 'Post updated with success!',
    }).then(function () {
      window.location = "/home";
    })
  }).fail(function () {
    Swal.fire({
      icon: 'error',
      title: 'Oops...',
      text: 'Something went wrong!',
    })
  }).always(function () {
    $("#update-post").prop('disabled', false)
  });
}

function deletePost(event) {
  event.preventDefault();

  Swal.fire({
    title: 'Are you sure?',
    text: "You won't be able to revert this!",
    icon: 'warning',
    showCancelButton: true,
    confirmButtonColor: '#3085d6',
    cancelButtonColor: '#d33',
    confirmButtonText: 'Yes, delete it!'
  }).then(function (confirmation) {
    if (confirmation.isConfirmed) {
      const clickedTrash = $(event.target);
      const post = clickedTrash.closest('div').parent();
      const postId = clickedTrash.closest('div').data('post-id');
      
      $.ajax({
        url: `/posts/${postId}`,
        method: "DELETE",
      }).done(function () {
        post.fadeOut("slow", function () {
          $(this).remove();
        });
      }).fail(function () {
        Swal.fire({
          icon: 'error',
          title: 'Oops...',
          text: 'Something went wrong to delete post!',
        })
      });
    }
  })


  // clickedTrash.prop('disabled', true);
}
