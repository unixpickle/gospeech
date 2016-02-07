(function() {

  var inputType = 'text';

  function handleInputTypeChanged() {
    var isIPA = document.getElementById('ipa-input').checked;
    inputType = isIPA ? 'ipa' : 'text';
  }

  function configureInputTypeControls() {
    var labeledRadios = document.getElementsByClassName('labeled-radio');
    for (var i = 0, len = labeledRadios.length; i < len; ++i) {
      var radio = labeledRadios[i];
      var input = radio.getElementsByTagName('input')[0];
      radio.addEventListener('click', function(input) {
        input.checked = true;
        handleInputTypeChanged();
      }.bind(null, input));
      input.addEventListener('change', handleInputTypeChanged);
    }
  }

  function setInputEnabled(flag) {
    var ids = ['synthesize-button', 'text-input'];
    for (var i = 0, len = ids.length; i < len; ++i) {
      document.getElementById(ids[i]).enabled = flag;
    }
  }

  window.app.inputType = function() {
    return inputType;
  };

  window.app.inputText = function() {
    return document.getElementById('text-input').value;
  };

  window.app.disableInput = function() {
    setInputEnabled(false);
  };

  window.app.enableInput = function() {
    setInputEnabled(true);
  };

  window.addEventListener('load', configureInputTypeControls);

})();
