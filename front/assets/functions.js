$(document).ready(function () {
    $("#country-title").on("click", function() {
        if (!$(".countries").hasClass("hidden")) {
            $(".countries").addClass("hidden");
            $("#country-title i").css("transform", "rotate(0deg)")
        } else {
            $(".countries").removeClass("hidden");
            $("#country-title i").css("transform", "rotate(180deg)")
        }

        if (!$(".leagues").hasClass("hidden"))
            $(".leagues").addClass("hidden");

        if (!$(".matches").hasClass("hidden"))
            $(".matches").addClass("hidden");

        if (!$(".standings").hasClass("hidden"))
            $(".standings").addClass("hidden");

        if (!$(".news").hasClass("hidden"))
            $(".news").addClass("hidden");

        if (!$(".ranking").hasClass("hidden"))
            $(".ranking").addClass("hidden");
    });

    $("#league-title").on("click", function() {
        if (!$(".leagues").hasClass("hidden")) {
            $(".leagues").addClass("hidden");
            $("#league-title i").css("transform", "rotate(0deg)")
        } else {
            $(".leagues").removeClass("hidden");
            $("#league-title i").css("transform", "rotate(180deg)")
        }

        if (!$(".countries").hasClass("hidden"))
            $(".countries").addClass("hidden");

        if (!$(".matches").hasClass("hidden"))
            $(".matches").addClass("hidden");

        if (!$(".standings").hasClass("hidden"))
            $(".standings").addClass("hidden");

        if (!$(".news").hasClass("hidden"))
            $(".news").addClass("hidden");

        if (!$(".ranking").hasClass("hidden"))
            $(".ranking").addClass("hidden");
    });

    $("#match-title").on("click", function() {
        if (!$(".matches").hasClass("hidden")) {
            $(".matches").addClass("hidden");
            $("#match-title i").css("transform", "rotate(0deg)")
        } else {
            $(".matches").removeClass("hidden");
            $("#match-title i").css("transform", "rotate(180deg)")
        }    

        if (!$(".countries").hasClass("hidden"))
            $(".countries").addClass("hidden");

        if (!$(".leagues").hasClass("hidden"))
            $(".leagues").addClass("hidden");

        if (!$(".standings").hasClass("hidden"))
            $(".standings").addClass("hidden");

        if (!$(".news").hasClass("hidden"))
            $(".news").addClass("hidden");

        if (!$(".ranking").hasClass("hidden"))
            $(".ranking").addClass("hidden");
    });

    $("#standing-title").on("click", function() {
        if (!$(".standings").hasClass("hidden")) {
            $(".standings").addClass("hidden");
            $("#standing-title i").css("transform", "rotate(0deg)")
        } else {
            $(".standings").removeClass("hidden");
            $("#standing-title i").css("transform", "rotate(180deg)")
        }

        if (!$(".countries").hasClass("hidden"))
            $(".countries").addClass("hidden");

        if (!$(".matches").hasClass("hidden"))
            $(".matches").addClass("hidden");

        if (!$(".leagues").hasClass("hidden"))
            $(".leagues").addClass("hidden");

        if (!$(".news").hasClass("hidden"))
            $(".news").addClass("hidden");

        if (!$(".ranking").hasClass("hidden"))
            $(".ranking").addClass("hidden");
    });

    $("#news-title").on("click", function() {
        if (!$(".news").hasClass("hidden")) {
            $(".news").addClass("hidden");
            $("#news-title i").css("transform", "rotate(0deg)")
        } else {
            $(".news").removeClass("hidden");
            $("#news-title i").css("transform", "rotate(180deg)")
        }

        if (!$(".countries").hasClass("hidden"))
            $(".countries").addClass("hidden");

        if (!$(".matches").hasClass("hidden"))
            $(".matches").addClass("hidden");

        if (!$(".leagues").hasClass("hidden"))
            $(".leagues").addClass("hidden");

        if (!$(".standings").hasClass("hidden"))
            $(".standings").addClass("hidden");

        if (!$(".ranking").hasClass("hidden"))
            $(".ranking").addClass("hidden");
    });

    $("#ranking-title").on("click", function() {
        if (!$(".ranking").hasClass("hidden")) {
            $(".ranking").addClass("hidden");
            $("#ranking-title i").css("transform", "rotate(0deg)")
        } else {
            $(".ranking").removeClass("hidden");
            $("#ranking-title i").css("transform", "rotate(180deg)")
        }

        if (!$(".countries").hasClass("hidden"))
            $(".countries").addClass("hidden");

        if (!$(".matches").hasClass("hidden"))
            $(".matches").addClass("hidden");

        if (!$(".leagues").hasClass("hidden"))
            $(".leagues").addClass("hidden");

        if (!$(".news").hasClass("hidden"))
            $(".news").addClass("hidden");

        if (!$(".standings").hasClass("hidden"))
            $(".standings").addClass("hidden");
    });
});