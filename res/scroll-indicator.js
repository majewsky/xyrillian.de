/*******************************************************************************
* Copyright 2017 Stefan Majewsky <majewsky@gmx.net>
* SPDX-License-Identifier: GPL-3.0
*******************************************************************************/

(function() {
  var si = document.querySelector(".scroll-indicator");
  var onscroll = function() {
    var max = document.body.scrollHeight - window.innerHeight;
    si.style.display = "block";
    si.style.width = (window.pageYOffset / max * 100) + "%";
  };

  window.onscroll = onscroll;
  onscroll();
})();
