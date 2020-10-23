import JSONEditor from 'jsoneditor';
import getOptions from '../utilities/getOptions';
import schema from '../utilities/schema.json';

const element = document.getElementById('json-editor');

let editor;

export function getJson() {
  return editor.get();
}

export function setJson(json) {
  editor.set(json);
}

export function enableSchema(isEnabled) {
  if (isEnabled) {
    editor.setSchema(schema);
  } else {
    editor.setSchema(null);
  }
}

export function init() {
  const autocomplete = {
    filter: 'contain',
    caseSensitive: false,
    getOptions,
  };
  const options = { autocomplete };
  editor = new JSONEditor(element, options);
  editor.set({ todo: 'upload file to start editing' });
  enableSchema(true);
}
