$(document).ready(function() {

    $(".show").on('click', function () {
        $(this).closest(".count").siblings(".article").show();
    })

    $(".close").on('click', function () {
        $(this).closest(".article").hide();
    })

});