import * as editor from '../actions/editor';

const saveInput = document.getElementById('save-slot');

export function getSlotIndex() {
  const slot = Number.parseInt(saveInput.value, 10);
  return slot;
}

export function setData(json) {
  const data = json;
  editor.setJson(data);
}

export function getData() {
  const data = editor.getJson();
  return data;
}
