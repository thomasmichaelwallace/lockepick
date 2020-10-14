const ff6 = require('../lib/ff6-release4-js/ff6-release4');

const jsonView = document.getElementById('jsonView');
const inputFile = document.getElementById('inputFile');
const downloadFile = document.getElementById('downloadFile');

// Start file download.

export function saveFile(blob, filename) {
  if (window.navigator.msSaveOrOpenBlob) {
    window.navigator.msSaveOrOpenBlob(blob, filename);
  } else {
    const a = document.createElement('a');
    document.body.appendChild(a);
    const url = window.URL.createObjectURL(blob);
    a.href = url;
    a.download = filename;
    a.click();
    setTimeout(() => {
      window.URL.revokeObjectURL(url);
      document.body.removeChild(a);
    }, 0);
  }
}

function init() {
  inputFile.addEventListener('change', (e) => {
    if (inputFile.files.length === 1) {
      console.log('reading...');
      inputFile.files[0].arrayBuffer()
        .then((a) => {
          console.log('read');
          const bytes = new Int8Array(a);
          console.log('to json bytes...')
          const [jsonBytes] = ff6.toJsonBytes(bytes);
          console.log('got bytes; parsing...');
          console.log(jsonBytes);
          const json = JSON.parse((new TextDecoder()).decode(jsonBytes));
          console.log('writing');
          jsonView.value = JSON.stringify(json, null, 2);
        })
        .catch((e) => console.error(e));
    }
    console.log('event', inputFile, e);
  })
  downloadFile.addEventListener('click', (e) => {
    const bytes = (new TextEncoder()).encode(jsonView.value);
    let [data] = ff6.fromJsonBytes(bytes)
    const blob = new Blob([data], { type: 'application/octet-stream' });
    saveFile(blob, 'data_two');
  });
}

init();
