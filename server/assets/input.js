(function() {

  var inputType = 'text';
  var ipaSpecialChars = 'əɛʌæʊɔŋθðʃʒɹʔɾ';

  function handleInputTypeChanged() {
    var isIPA = document.getElementById('ipa-input-type').checked;
    inputType = isIPA ? 'ipa' : 'text';

    var ipaInput = document.getElementById('ipa-input');
    var ipaInputBreak = document.getElementById('ipa-input-break');
    if (isIPA) {
      ipaInput.style.display = 'inline-block';
      ipaInputBreak.style.display = 'block';
    } else {
      ipaInput.style.display = 'none';
      ipaInputBreak.style.display = 'none';
    }
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

  function configureSpecialCharacters() {
    var ipaInput = document.getElementById('ipa-input');
    for (var i = 0, len = ipaSpecialChars.length; i < len; ++i) {
      var char = ipaSpecialChars[i];
      var button = document.createElement('button');
      button.innerText = char;
      button.addEventListener('click', function(char) {
        var textInput = document.getElementById('text-input');
        textInput.value += char;
        if ('number' === typeof textInput.selectionStart) {
          textInput.selectionStart = textInput.selectionEnd = textInput.value.length;
        } else if ('undefined' !== typeof textInput.createTextRange) {
          textInput.focus();
          var range = textInput.createTextRange();
          e.collapse(false);
          range.select();
        }
      }.bind(this, char));
      button.className = 'ipa-input-symbol';
      ipaInput.appendChild(button);
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

  window.addEventListener('load', function() {
    configureInputTypeControls();
    configureSpecialCharacters();
    handleInputTypeChanged();
  });

})();
