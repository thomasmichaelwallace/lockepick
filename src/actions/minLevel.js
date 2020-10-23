/* eslint-disable no-console */
const input = document.getElementById('value-min-level');

function handler() {
  const level = Number.parseInt(input.value, 10);
  if (Number.isNaN(level) || level < 1 || level > 99) {
    console.error('min level must be between 0-100');
    return;
  }
  console.log('set level', level);
}

// eslint-disable-next-line import/prefer-default-export
export function init() {
  const button = document.getElementById('set-min-level');
  button.addEventListener('click', handler);
}
