$(document).ready(function() {
    $('.slider').slick({
        arrows: true,
        dots: false,
        adaptiveHeight: true,
        slidesToShow:1,
        slidesToScrol:1,
        speed:1000,
        easing: 'lianer',
        itFinite: true,
        touchThreshold:8,
        autoplayspeed: 1000,
        autoplay: true,
        pauseondotshover: true,
       // responsive: [{
           // breakpoint: 768,
           // settings: {
           //     slidesToShow:1 
            //}
        //}]
    });
    var s=$('.slider').slick('slickGetOption', 'SlidesToShow');
    console_log=(s);
});