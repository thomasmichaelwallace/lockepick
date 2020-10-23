/* eslint-disable no-console */
import { getData, getSlotIndex, setData } from '../utilities/state';

// https://steamcommunity.com/sharedfiles/filedetails/?id=581570122
const spellsByEsper = {
  /* cSpell:disable */
  'Cait Sith': ['Confuse', 'Float', 'Imp'],
  'Zona Seeker': ['Osmose', 'Rasp', 'Shell'],
  Alexandr: ['Dispel', 'Esuna', 'Holy', 'Protect', 'Shell'],
  Bahamut: ['Flare'],
  Bismark: ['Blizzard', 'Fire', 'Raise', 'Thunder'],
  Cactuar: ['Hastega', 'Teleport', 'Vanish'],
  Carbunkl: ['Haste', 'Protect', 'Reflect', 'Shell', 'Teleport'],
  Catoblepas: ['Bio', 'Break', 'Death'],
  Crusader: ['Meltdown', 'Meteor'],
  Diabolos: ['Graviga', 'Gravija'],
  Fenrir: ['Banish', 'Stop', 'Teleport'],
  Gilgamesh: ['Quick', 'Valor'],
  Golem: ['Cura', 'Protect', 'Stop'],
  Ifrit: ['Drain', 'Fira', 'Fire'],
  Kirin: ['Cura', 'Cure', 'Libra', 'Poisona', 'Regen'],
  Lakshmi: ['Cura', 'Curaga', 'Cure', 'Esuna', 'Regen'],
  Leviathan: ['Flood'],
  Maduin: ['Blizzara', 'Fira', 'Thundara'],
  Midgardsormr: ['Graviga', 'Quake', 'Tornado'],
  Odin: ['Meteor'],
  Phantom: ['Berserk', 'Gravity', 'Vanish'],
  Phoenix: ['Arise', 'Curaga', 'Firaga', 'Raise', 'Reraise'],
  Quetzalli: ['Float', 'Haste', 'Hastega', 'Slow', 'Slowga'],
  Ragnarok: ['Ultima'],
  Raiden: ['Quick'],
  Ramuh: ['Poison', 'Thundara', 'Thunder'],
  Seraph: ['Cura', 'Cure', 'Esuna', 'Raise', 'Regen'],
  Shiva: ['Blizzara', 'Blizzard', 'Cure', 'Osmose', 'Rasp'],
  Siren: ['Fire', 'Silence', 'Sleep', 'Slow'],
  Unicorn: ['Cura', 'Dispel', 'Esuna', 'Protect', 'Shell'],
  Valigarmanda: ['Blizzaga', 'Firaga', 'Thundaga'],
  /* cSpell:enable */
};

function handler() {
  const data = getData();
  const slotIndex = getSlotIndex();
  const slot = data.Saves[slotIndex];

  const spells = new Set();
  Object.entries(slot.Iv.EspersOwned).forEach(([e, o]) => {
    if (!o) return; // exclude unowned.
    const esperSpells = spellsByEsper[e];
    if (!esperSpells) {
      console.warn('unknown esper id', e);
      return;
    }
    esperSpells.forEach((s) => spells.add(s));
  });

  spells.forEach((s) => {
    Object.values(slot.Body.SpellMastery).forEach((c) => {
      if (c[s] === undefined) {
        console.warn('unknown spell name', s);
        return;
      }
      // eslint-disable-next-line no-param-reassign
      c[s] = 100;
    });
  });

  setData(data);
}

// eslint-disable-next-line import/prefer-default-export
export function init() {
  const button = document.getElementById('quick-spell');
  button.addEventListener('click', handler);
}
