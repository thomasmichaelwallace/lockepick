import { getData, getSlotIndex, setData } from '../utilities/state';

const orderKey = 'SortingOrder';
const nothingName = 'Nothing';
const orderLength = 288;

function handler() {
  const data = getData();

  const slotIndex = getSlotIndex();
  const slot = data.Saves[slotIndex];

  const inventory = slot.Iv.Inventory;

  inventory[orderKey] = Object.keys(inventory).filter((k) => k !== nothingName && k !== orderKey);
  while (inventory[orderKey].length < orderLength) {
    inventory[orderKey].push(nothingName);
  }

  setData(data);
}

// eslint-disable-next-line import/prefer-default-export
export function init() {
  const button = document.getElementById('fix-items');
  button.addEventListener('click', handler);
}
