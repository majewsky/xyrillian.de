(function() {

  function foreach(parent, selector, callback) {
    Array.prototype.forEach.call(parent.querySelectorAll(selector), callback);
  }

  //onclick handler for the tab bars on podcast episodes
  var onclickHandler = function() {
    var article = this.parentNode.parentNode;
    var tabName = this.getAttribute("data-tab");
    //change activation state of tab bar buttons
    foreach(article, "article > nav > button", function(button) {
      button.classList.toggle("selected", button.getAttribute("data-tab") === tabName);
    });
    //toggle visibility of tab contents
    foreach(article, ".nav-show, .nav-hide", function(node) {
      var match = node.getAttribute("data-tab") === tabName;
      node.classList.toggle("nav-show", match);
      node.classList.toggle("nav-hide", !match);
    });
  };
  foreach(document, "body.noises article > nav > button", function(button) {
    button.onclick = onclickHandler;
  });

})();
