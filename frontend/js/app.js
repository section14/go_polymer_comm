(function(document) {
  'use strict';

  // Grab a reference to our auto-binding template
  // and give it some initial binding values
  var app = document.querySelector("#app");

  app.displayInstalledToast = function() {
    document.querySelector("#caching-complete").show();
  };

  app.addEventListener("dom-change", function() {
    console.log("Our app is ready to rock!");
  });

  window.addEventListener("WebComponentsReady", function() {
    // imports are loaded and elements have been registered
  });

})(document);
