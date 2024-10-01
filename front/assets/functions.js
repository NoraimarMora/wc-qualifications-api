$(document).ready(function () {
    $("#country-title").on("click", function() {
        if (!$(".countries").hasClass("hidden"))
            $(".countries").addClass("hidden");
        else
            $(".countries").removeClass("hidden");

        if (!$(".leagues").hasClass("hidden"))
            $(".leagues").addClass("hidden");

        if (!$(".matches").hasClass("hidden"))
            $(".matches").addClass("hidden");

        if (!$(".standings").hasClass("hidden"))
            $(".standings").addClass("hidden");
    });

    $("#league-title").on("click", function() {
        if (!$(".leagues").hasClass("hidden"))
            $(".leagues").addClass("hidden");
        else
            $(".leagues").removeClass("hidden");

        if (!$(".countries").hasClass("hidden"))
            $(".countries").addClass("hidden");

        if (!$(".matches").hasClass("hidden"))
            $(".matches").addClass("hidden");

        if (!$(".standings").hasClass("hidden"))
            $(".standings").addClass("hidden");
    });

    $("#match-title").on("click", function() {
        if (!$(".matches").hasClass("hidden"))
            $(".matches").addClass("hidden");
        else
            $(".matches").removeClass("hidden");

        if (!$(".countries").hasClass("hidden"))
            $(".countries").addClass("hidden");

        if (!$(".leagues").hasClass("hidden"))
            $(".leagues").addClass("hidden");

        if (!$(".standings").hasClass("hidden"))
            $(".standings").addClass("hidden");
    });

    $("#standing-title").on("click", function() {
        if (!$(".standings").hasClass("hidden"))
            $(".standings").addClass("hidden");
        else
            $(".standings").removeClass("hidden");
        
        if (!$(".countries").hasClass("hidden"))
            $(".countries").addClass("hidden");

        if (!$(".matches").hasClass("hidden"))
            $(".matches").addClass("hidden");

        if (!$(".leagues").hasClass("hidden"))
            $(".leagues").addClass("hidden");
    });
});