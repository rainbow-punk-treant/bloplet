$( document ).ready(function() {

    var doc = $('container');
    document.addEventListener("drag", function(event) {
      doc.css.color = "red";
    });
    document.addEventListener("dragenter", function(event) {
      doc.css.border = "dotted";
    });
    document.addEventListener("drop", function(event) {
    event.preventDefault();
    if ( event.target.className == "droptarget" ) {
        demo.css.color = "";
        event.target.css.border = "";
        var data = event.dataTransfer.getData("Text");
        event.target.appendChild(document.getElementById(data));
    }
});
    baguetteBox.run('.gallery');
}); //this is a callback! A self contained system of,
//when the document is ready, execute this anonymous
//function! Which happens to be all of the javascript
//required when the page is ready and done loading!
