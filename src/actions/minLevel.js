import { getData, getSlotIndex, setData } from '../utilities/state';

/* eslint-disable no-console */
const input = document.getElementById('value-min-level');

// https://gamefaqs.gamespot.com/ps/562865-final-fantasy-vi/faqs/39433
const statIncrements = {
  1: {
    exp: 0, hpp: 0, hp: 0, mpp: 0, mp: 0,
  },
  2: {
    exp: 32, hpp: 11, hp: 11, mpp: 5, mp: 5,
  },
  3: {
    exp: 96, hpp: 12, hp: 23, mpp: 6, mp: 11,
  },
  4: {
    exp: 208, hpp: 14, hp: 37, mpp: 7, mp: 18,
  },
  5: {
    exp: 400, hpp: 17, hp: 54, mpp: 8, mp: 26,
  },
  6: {
    exp: 672, hpp: 20, hp: 74, mpp: 9, mp: 35,
  },
  7: {
    exp: 1056, hpp: 22, hp: 96, mpp: 10, mp: 45,
  },
  8: {
    exp: 1552, hpp: 24, hp: 120, mpp: 11, mp: 56,
  },
  9: {
    exp: 2184, hpp: 26, hp: 146, mpp: 12, mp: 68,
  },
  10: {
    exp: 2976, hpp: 27, hp: 173, mpp: 13, mp: 81,
  },
  11: {
    exp: 3936, hpp: 28, hp: 201, mpp: 14, mp: 95,
  },
  12: {
    exp: 5080, hpp: 30, hp: 231, mpp: 15, mp: 110,
  },
  13: {
    exp: 6432, hpp: 35, hp: 266, mpp: 16, mp: 126,
  },
  14: {
    exp: 7992, hpp: 39, hp: 305, mpp: 17, mp: 143,
  },
  15: {
    exp: 9784, hpp: 44, hp: 349, mpp: 17, mp: 160,
  },
  16: {
    exp: 11840, hpp: 50, hp: 399, mpp: 17, mp: 177,
  },
  17: {
    exp: 14152, hpp: 54, hp: 453, mpp: 17, mp: 194,
  },
  18: {
    exp: 16736, hpp: 57, hp: 510, mpp: 16, mp: 210,
  },
  19: {
    exp: 19616, hpp: 61, hp: 571, mpp: 16, mp: 226,
  },
  20: {
    exp: 22832, hpp: 65, hp: 636, mpp: 16, mp: 242,
  },
  21: {
    exp: 26360, hpp: 67, hp: 703, mpp: 15, mp: 257,
  },
  22: {
    exp: 30232, hpp: 69, hp: 772, mpp: 15, mp: 272,
  },
  23: {
    exp: 34456, hpp: 72, hp: 844, mpp: 15, mp: 287,
  },
  24: {
    exp: 39056, hpp: 76, hp: 920, mpp: 14, mp: 301,
  },
  25: {
    exp: 44072, hpp: 79, hp: 999, mpp: 14, mp: 315,
  },
  26: {
    exp: 49464, hpp: 82, hp: 1081, mpp: 14, mp: 329,
  },
  27: {
    exp: 55288, hpp: 86, hp: 1167, mpp: 14, mp: 343,
  },
  28: {
    exp: 61568, hpp: 90, hp: 1257, mpp: 13, mp: 356,
  },
  29: {
    exp: 68304, hpp: 95, hp: 1352, mpp: 13, mp: 369,
  },
  30: {
    exp: 75496, hpp: 99, hp: 1451, mpp: 13, mp: 382,
  },
  31: {
    exp: 83184, hpp: 100, hp: 1551, mpp: 13, mp: 395,
  },
  32: {
    exp: 91384, hpp: 101, hp: 1652, mpp: 12, mp: 407,
  },
  33: {
    exp: 100088, hpp: 102, hp: 1754, mpp: 12, mp: 419,
  },
  34: {
    exp: 109344, hpp: 102, hp: 1856, mpp: 12, mp: 431,
  },
  35: {
    exp: 119136, hpp: 103, hp: 1959, mpp: 12, mp: 443,
  },
  36: {
    exp: 129504, hpp: 104, hp: 2063, mpp: 11, mp: 454,
  },
  37: {
    exp: 140464, hpp: 106, hp: 2169, mpp: 11, mp: 465,
  },
  38: {
    exp: 152008, hpp: 107, hp: 2276, mpp: 11, mp: 476,
  },
  39: {
    exp: 164184, hpp: 108, hp: 2384, mpp: 11, mp: 487,
  },
  40: {
    exp: 176976, hpp: 110, hp: 2494, mpp: 11, mp: 498,
  },
  41: {
    exp: 190416, hpp: 111, hp: 2605, mpp: 11, mp: 509,
  },
  42: {
    exp: 204520, hpp: 113, hp: 2718, mpp: 10, mp: 519,
  },
  43: {
    exp: 219320, hpp: 114, hp: 2832, mpp: 10, mp: 529,
  },
  44: {
    exp: 234808, hpp: 116, hp: 2948, mpp: 10, mp: 539,
  },
  45: {
    exp: 251000, hpp: 117, hp: 3065, mpp: 10, mp: 549,
  },
  46: {
    exp: 267936, hpp: 119, hp: 3184, mpp: 10, mp: 559,
  },
  47: {
    exp: 285600, hpp: 120, hp: 3304, mpp: 10, mp: 569,
  },
  48: {
    exp: 304040, hpp: 122, hp: 3426, mpp: 10, mp: 579,
  },
  49: {
    exp: 323248, hpp: 125, hp: 3551, mpp: 10, mp: 589,
  },
  50: {
    exp: 343248, hpp: 128, hp: 3679, mpp: 9, mp: 598,
  },
  51: {
    exp: 364064, hpp: 130, hp: 3809, mpp: 9, mp: 607,
  },
  52: {
    exp: 385696, hpp: 131, hp: 3940, mpp: 9, mp: 616,
  },
  53: {
    exp: 408160, hpp: 133, hp: 4073, mpp: 9, mp: 625,
  },
  54: {
    exp: 431488, hpp: 134, hp: 4207, mpp: 9, mp: 634,
  },
  55: {
    exp: 455680, hpp: 136, hp: 4343, mpp: 9, mp: 643,
  },
  56: {
    exp: 480776, hpp: 137, hp: 4480, mpp: 9, mp: 652,
  },
  57: {
    exp: 506760, hpp: 139, hp: 4619, mpp: 9, mp: 661,
  },
  58: {
    exp: 533680, hpp: 142, hp: 4761, mpp: 9, mp: 670,
  },
  59: {
    exp: 561528, hpp: 144, hp: 4905, mpp: 9, mp: 679,
  },
  60: {
    exp: 590320, hpp: 145, hp: 5050, mpp: 10, mp: 689,
  },
  61: {
    exp: 620096, hpp: 147, hp: 5197, mpp: 8, mp: 697,
  },
  62: {
    exp: 650840, hpp: 148, hp: 5345, mpp: 8, mp: 705,
  },
  63: {
    exp: 682600, hpp: 150, hp: 5495, mpp: 8, mp: 713,
  },
  64: {
    exp: 715368, hpp: 152, hp: 5647, mpp: 8, mp: 721,
  },
  65: {
    exp: 749160, hpp: 153, hp: 5800, mpp: 8, mp: 729,
  },
  66: {
    exp: 784016, hpp: 155, hp: 5955, mpp: 8, mp: 737,
  },
  67: {
    exp: 819920, hpp: 156, hp: 6111, mpp: 8, mp: 745,
  },
  68: {
    exp: 856920, hpp: 158, hp: 6269, mpp: 8, mp: 753,
  },
  69: {
    exp: 895016, hpp: 160, hp: 6429, mpp: 8, mp: 761,
  },
  70: {
    exp: 934208, hpp: 162, hp: 6591, mpp: 8, mp: 769,
  },
  71: {
    exp: 974536, hpp: 160, hp: 6751, mpp: 10, mp: 779,
  },
  72: {
    exp: 1016000, hpp: 155, hp: 6906, mpp: 10, mp: 789,
  },
  73: {
    exp: 1058640, hpp: 151, hp: 7057, mpp: 7, mp: 796,
  },
  74: {
    exp: 1102456, hpp: 145, hp: 7202, mpp: 6, mp: 802,
  },
  75: {
    exp: 1147456, hpp: 140, hp: 7342, mpp: 5, mp: 807,
  },
  76: {
    exp: 1193648, hpp: 136, hp: 7478, mpp: 4, mp: 811,
  },
  77: {
    exp: 1241080, hpp: 132, hp: 7610, mpp: 5, mp: 816,
  },
  78: {
    exp: 1289744, hpp: 126, hp: 7736, mpp: 6, mp: 822,
  },
  79: {
    exp: 1339672, hpp: 120, hp: 7856, mpp: 7, mp: 829,
  },
  80: {
    exp: 1390872, hpp: 117, hp: 7973, mpp: 8, mp: 837,
  },
  81: {
    exp: 1443368, hpp: 113, hp: 8086, mpp: 9, mp: 846,
  },
  82: {
    exp: 1497160, hpp: 110, hp: 8196, mpp: 8, mp: 854,
  },
  83: {
    exp: 1552264, hpp: 108, hp: 8304, mpp: 7, mp: 861,
  },
  84: {
    exp: 1608712, hpp: 105, hp: 8409, mpp: 6, mp: 867,
  },
  85: {
    exp: 1666512, hpp: 102, hp: 8511, mpp: 5, mp: 872,
  },
  86: {
    exp: 1725688, hpp: 100, hp: 8611, mpp: 6, mp: 878,
  },
  87: {
    exp: 1786240, hpp: 98, hp: 8709, mpp: 7, mp: 885,
  },
  88: {
    exp: 1848184, hpp: 95, hp: 8804, mpp: 5, mp: 890,
  },
  89: {
    exp: 1911552, hpp: 92, hp: 8896, mpp: 6, mp: 896,
  },
  90: {
    exp: 1976352, hpp: 90, hp: 8986, mpp: 7, mp: 903,
  },
  91: {
    exp: 2042608, hpp: 88, hp: 9074, mpp: 8, mp: 911,
  },
  92: {
    exp: 2110320, hpp: 87, hp: 9161, mpp: 9, mp: 920,
  },
  93: {
    exp: 2179504, hpp: 85, hp: 9246, mpp: 10, mp: 930,
  },
  94: {
    exp: 2250192, hpp: 83, hp: 9329, mpp: 8, mp: 938,
  },
  95: {
    exp: 2322392, hpp: 82, hp: 9411, mpp: 8, mp: 946,
  },
  96: {
    exp: 2396128, hpp: 80, hp: 9491, mpp: 9, mp: 955,
  },
  97: {
    exp: 2471400, hpp: 83, hp: 9574, mpp: 10, mp: 965,
  },
  98: {
    exp: 2548224, hpp: 86, hp: 9660, mpp: 11, mp: 976,
  },
  99: {
    exp: 2637112, hpp: 88, hp: 9748, mpp: 13, mp: 989,
  },
};
const statBase = { // [hp, mp]
  /* cSpell:disable */
  Terra: [40, 16],
  Locke: [48, 7],
  Cyan: [53, 5],
  Shadow: [51, 6],
  Edgar: [49, 6],
  Sabin: [58, 3],
  Celes: [44, 15],
  Strago: [35, 13],
  Relm: [37, 18],
  Setzer: [46, 9],
  Mog: [39, 16],
  Gau: [45, 10],
  Gogo: [36, 12],
  Umaro: [60, 0],
  /* cSpell:enable */
};

function handler() {
  const level = Number.parseInt(input.value, 10);
  if (Number.isNaN(level) || level < 1 || level > 99) {
    console.error('min level must be between 0-100');
    return;
  }
  const data = getData();
  const slotIndex = getSlotIndex();
  const slot = data.Saves[slotIndex];
  slot.Body.Characters.forEach((c) => {
    if (c.Level >= level) return;

    const base = statBase[c.Id];
    if (!base) {
      console.warn('no level stats for character id:', c.Id);
      return;
    }
    const stats = statIncrements[level];
    const hp = base[0] + stats.hp;
    const mp = base[1] + stats.mp;

    /* eslint-disable no-param-reassign */
    c.Level = level;
    c.CurrHp = hp;
    c.TotHp = hp;
    c.CurrMp = mp;
    c.TotMp = mp;
    c.Experience = stats.exp;
    /* eslint-enable no-param-reassign */
  });
  setData(data);
}

// eslint-disable-next-line import/prefer-default-export
export function init() {
  const button = document.getElementById('set-min-level');
  button.addEventListener('click', handler);
}
