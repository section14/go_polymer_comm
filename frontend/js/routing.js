window.addEventListener("WebComponentsReady", function() {

    page.base("");

    page('/', function () {
        app.route = "home";
        app.pageContent = "/app/general/home.html";
    });

    page('/admin', function () {
        app.route = "admin";
        app.pageContent = "/app/admin/dashboard.html";
    });

    /*
    page('/users/:name', function (data) {
      app.route = 'user-info';
      app.params = data.params;
    });
    */

    page('/product', function () {
        app.route = "product";
        app.pageContent = "/app/general/product.html";
    });

    page('/contact', function () {
        app.route = "contact";
        app.pageContent = "/app/general/contact.html";
    });

    // add #! before urls
    page({
        hashbang: true
    });

});
