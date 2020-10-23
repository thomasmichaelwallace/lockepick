/* eslint-disable no-console */
import * as editor from '../actions/editor';

const saveInput = document.getElementById('save-slot');

let data = { todo: 'upload valid ff6 save file' };
let isSimple = true;

function withoutBits(obj) {
  return Object.entries(obj).reduce((o, [k, v]) => {
    if (k.startsWith('Unknown') || k.startsWith('Something')) return o;
    return { ...o, [k]: v };
  }, {});
}

function toSimpleView(json) {
  try {
    const view = json.Saves.map((s) => {
      const save = {
        ...withoutBits(s.Body),
        Characters: s.Body.Characters.map(withoutBits),
        ...withoutBits(s.Iv),
        ...withoutBits(s.Body2),
      };
      return save;
    });
    return view;
  } catch (error) {
    console.error('error simplifying view', error);
    return json;
  }
}

function mergeIfKey(whole, partial) {
  const o = {};
  Object.keys(whole).forEach((k) => {
    if (partial[k]) {
      o[k] = partial[k];
    } else {
      o[k] = whole[k];
    }
  });
  return o;
}

function fromSimpleView(view, json) {
  const full = json;
  try {
    view.forEach((v, i) => {
      const save = full.Saves[i];
      const merged = {
        Body: {
          ...mergeIfKey(save.Body, v),
          Characters: save.Body.Characters.map((c, j) => mergeIfKey(c, v.Characters[j])),
        },
        Iv: mergeIfKey(save.Iv, v),
        Body2: mergeIfKey(save.Body2, v),
      };
      full.Saves[i] = merged;
    });
    return full;
  } catch (error) {
    console.error('error merging view', error);
    return data;
  }
}

export function getSlotIndex() {
  const slot = Number.parseInt(saveInput.value, 10);
  return slot;
}

export function setData(json) {
  data = json;
  const view = isSimple ? toSimpleView(data) : data;
  editor.setJson(view);
}

export function getData() {
  const view = editor.getJson();
  data = isSimple ? fromSimpleView(view, data) : view;
  return JSON.parse(JSON.stringify(data)); // clone to prevent mutation.
}

export function setAdvancedMode(isAdvanced) {
  if (isAdvanced !== isSimple) return; // no-op.
  console.log('switching mode', isAdvanced ? 'advanced' : 'simple');
  getData(); // save editor data
  isSimple = !isAdvanced;
  setData(data); // reload editor state
  editor.enableSchema(isSimple);
}
