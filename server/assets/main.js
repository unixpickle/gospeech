(function() {

  var READY_STATE_DONE = 4;
  var HTTP_OK = 200;

  function synthesize() {
    window.app.disableInput();
    var inputType = window.app.inputType();
    var inputText = window.app.inputText();

    var path = '/synthesize_' + inputType + '?' + inputType + '=' +
      encodeURIComponent(inputText);

    var audio = new Audio(path);
    audio.addEventListener('error', function(e) {
      window.app.enableInput();
      alert('Failed to play.');
    });
    audio.addEventListener('canplaythrough', function() {
      window.app.enableInput();
      audio.play();
    });
  }

  window.addEventListener('load', function() {
    var synthButton = document.getElementById('synthesize-button');
    synthButton.addEventListener('click', synthesize);
  });

})();
