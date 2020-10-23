import JSONEditor from 'jsoneditor';

const element = document.getElementById('json-editor');

let editor;
export function init() {
  const options = {};
  editor = new JSONEditor(element, options);
  editor.set({ todo: 'upload file to start editing' });
}

export function getJson() {
  return editor.get();
}

export function setJson(json) {
  editor.set(json);
}
