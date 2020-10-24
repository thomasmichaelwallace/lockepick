import { getData, getSlotIndex, setData } from '../utilities/state';

/* eslint-disable no-console */
const input = document.getElementById('value-min-items');

const orderKey = 'SortingOrder';
const nothingName = 'Nothing';

function handler() {
  const count = Number.parseInt(input.value, 10);
  if (Number.isNaN(count) || count < 1 || count > 99) {
    console.error('min items must be between 0-100');
    return;
  }
  const data = getData();
  const slotIndex = getSlotIndex();
  const slot = data.Saves[slotIndex];

  const invetory = slot.Iv.Inventory;
  Object.keys(invetory)
    .filter((k) => k !== orderKey && k !== nothingName)
    .forEach((k) => {
      invetory[k] = Math.max(invetory[k], count);
    });

  setData(data);
}

// eslint-disable-next-line import/prefer-default-export
export function init() {
  const button = document.getElementById('set-min-items');
  button.addEventListener('click', handler);
}
