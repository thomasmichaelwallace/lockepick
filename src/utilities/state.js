import * as editor from '../actions/editor';

export function setData(json) {
  const data = json;
  editor.setJson(data);
}

export function getData() {
  const data = editor.getJson();
  return data;
}
