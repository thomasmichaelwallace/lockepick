/* eslint-disable no-console */
import * as ff6 from '../../lib/ff6-release4-js/ff6-release4';
import saveFile from '../utilities/saveFile';
import { getData, setAdvancedMode, setData } from '../utilities/state';

const uploadButton = document.getElementById('upload-file');
const downloadButton = document.getElementById('download-file');
const advancedInput = document.getElementById('advanced-mode');

async function handleUploadFile() {
  try {
    if (uploadButton.files.length !== 1) {
      console.error('expected exactly one file');
      return;
    }

    console.log('reading uploaded file...');
    const array = await uploadButton.files[0].arrayBuffer();
    const bytes = new Int8Array(array);
    console.log('parsing file...');
    const [jsonBytes] = ff6.toJsonBytes(bytes);
    const json = JSON.parse((new TextDecoder()).decode(jsonBytes));
    setData(json);
    console.log('upload finished successfully.');
  } catch (error) {
    console.error('failed to read file', error);
  }
}

async function handleDownloadFile() {
  try {
    console.log('parsing json...');
    const json = getData();
    const bytes = (new TextEncoder()).encode(JSON.stringify(json));
    const [data] = ff6.fromJsonBytes(bytes);
    console.log('building download...');
    const blob = new Blob([data], { type: 'application/octet-stream' });
    saveFile(blob, 'data');
    console.log('download finished successfully.');
  } catch (error) {
    console.error('failed to build file', error);
  }
}

function handleAdvancedChange(e) {
  const isAdvanced = e.srcElement.checked;
  setAdvancedMode(isAdvanced);
}

// eslint-disable-next-line import/prefer-default-export
export function init() {
  uploadButton.addEventListener('change', handleUploadFile);
  downloadButton.addEventListener('click', handleDownloadFile);
  advancedInput.addEventListener('change', handleAdvancedChange);
}
