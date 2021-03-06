/* eslint-disable no-console */
import { getData, getSlotIndex, setData } from '../utilities/state';

const names = [
  /* cSpell:disable */
  'Dagger',
  'Mythril Knife',
  'Main Gauche',
  'Air Knife',
  "Thief's Knife",
  "Assassin's Dagger",
  'Man-Eater',
  'Swordbreaker',
  'Gladius',
  'Valiant Knife',
  'Mythril Sword',
  'Great Sword',
  'Rune Blade',
  'Flametounge',
  'Icebrand',
  'Thunder Blade',
  'Bastard Sword',
  'Stoneblade',
  'Blood Sword',
  'Enhancer',
  'Crystal Sword',
  'Falchion',
  'Soul Sabre',
  'Organyx',
  'Excalibur',
  'Zantetsuken',
  'Lightbringer',
  'Ragnarok',
  'Ultima Weapon',
  'Mythril Spear',
  'Trident',
  'Heavy Lance',
  'Partisian',
  'Holy Lance',
  'Golden Spear',
  'Radiant Lance',
  'Impartisian',
  'Kunai',
  'Kodachi',
  'Sakura',
  'Sasuke',
  'Ichigeki',
  'Kagenui',
  'Ashura',
  'Kotetsu',
  'Kiku-ichimonji',
  'Kazekiri',
  'Murasame',
  'Masamune',
  'Murakuma',
  'Mutsunokami',
  'Healing Rod',
  'Mythril Rod',
  'Flame Rod',
  'Ice Rod',
  'Thunder Rod',
  'Poison Rod',
  'Holy Rod',
  'Gravity Rod',
  'Punisher',
  'Magus Rod',
  'Chocobo Brush',
  'DaVinci Brush',
  'Magical Brush',
  'Rainbow Brush',
  'Shuriken',
  'Fuma Shuriken',
  'Pinwheel',
  'Chain Flail',
  'Moonring Blade',
  'Morning Star',
  'Boomerang',
  'Rising Sun',
  'Hawkeye',
  'Bone Club',
  'Sniper',
  'Wing Edge',
  'Cards',
  'Darts',
  'Death Tarot',
  'Viper Darts',
  'Dice',
  'Fixed Dice',
  'Metal Knuckle',
  'Mythril Claw',
  'Kaiser Knuckle',
  'Venom Claws',
  'Burning Fist',
  'Dragon Claws',
  'Tigerfang',
  'Buckler',
  'Heavy Shield',
  'Mythril Shield',
  'Golden Shield',
  'Aegis Shield',
  'Diamond Shield',
  'Flame Shield',
  'Ice Shield',
  'Thunder Shield',
  'Crystal Shield',
  'Genji Shield',
  'Tortoise Shield',
  'Cursed Shield',
  "Paladin's Shield",
  'Force Shield',
  'Leather Hat',
  'Hairband',
  'Plumed Hat',
  'Beret',
  'Magus Hat',
  'Bandana',
  'Iron Helmet',
  'Hypno Crown',
  "Priest's Miter",
  'Green Beret',
  'Twist Headband',
  'Mythril Helm',
  'Tiara',
  'Golden Helmet',
  'Tiger Mask',
  'Red Cap',
  'Mystery Veil',
  'Circlet',
  'Royal Crown',
  'Diamond Helm',
  'Black Cowl',
  'Crystal Helm',
  'Oath Viel',
  'Cat-Ear Hood',
  'Genji Helmet',
  'Thornlet',
  'Saucer',
  'Leather Armor',
  'Cotton Robe',
  'Kenpo Gi',
  'Iron Armor',
  'Silk Robe',
  'Mythril Vest',
  'Ninja Gear',
  'White Dress',
  'Mythril Mail',
  'Gaia Gear',
  'Mirage Dress',
  'Golden Armor',
  'Power Sash',
  'Luminous Robe',
  'Diamond Vest',
  'Red Jacket',
  'Force Armor',
  'Diamond Armor',
  'Black Garb',
  'Magus Robe',
  'Crystal Mail',
  'Regal Gown',
  'Genji Armor',
  'Reed Cloak',
  'Minerva Bustier',
  'Tabby Suit',
  'Chocobo Suit',
  'Moogle Suit',
  'Nutkin Suit',
  'Behemoth Suit',
  'Snow Scarf',
  'Noiseblaster',
  'Bioblaster',
  'Flash',
  'Chainsaw',
  'Debilitator',
  'Drill',
  'Air Anchor',
  'Auto Crossbow',
  'Flame Scroll',
  'Water Scroll',
  'Lightning Scroll',
  'Invisibility Scroll',
  'Shadow Scroll',
  'Silver Spectacles',
  'Star Pendant',
  'Peace Ring',
  'Amulet',
  'White Cape',
  'Jeweled Ring',
  'Fairy Ring',
  'Barrier Ring',
  'Mythril Glove',
  'Protect Ring',
  'Hermes Sandals',
  'Reflect Ring',
  'Angel Wings',
  'Angel Ring',
  'Knights Code',
  'Dragon Boots',
  'Zephyr Cloak',
  'Princess Ring',
  'Cursed Ring',
  'Earrings',
  'Gigas Glove',
  'Blizzard Ring',
  'Berserker Ring',
  "Thief's Bracer",
  'Guard Bracelet',
  "Hero's Ring",
  'Ribbon',
  'Muscle Belt',
  'Crystal Orb',
  'Gold Hairpin',
  'Celestriad',
  "Brigand's Glove",
  'Gauntlet',
  'Genji Glove',
  'Hyper Wrist',
  "Master's Scroll",
  'Prayer Beads',
  'Black Belt',
  "Heji's Jitte",
  'Fake Moustache',
  'Soul Of Tamasa',
  'Dragon Horn',
  'Merit Award',
  'Memento Ring',
  'Safety Bit',
  'Lich Ring',
  "Molulu's Charm",
  'Ward Bangle',
  'Miracle Shoes',
  'Alarm Earring',
  'Gale Hairpin',
  'Sniper Eye',
  'Growth Egg',
  'Tintinabulum',
  'Sprint Shoes',
  'Rename Card',
  'Potion',
  'Hi-Potion',
  'X-Potion',
  'Ether',
  'Hi-Ether',
  'X-Ether',
  'Elixir',
  'Megalixir',
  'Phoenix Down',
  'Holy Water',
  'Antidote',
  'Eye Dropss',
  'Gold Needle',
  'Remedy',
  'Sleeping Bag',
  'Tent',
  'Green Cherry',
  'Magicite Shard',
  'Super Ball',
  'Echo Screen',
  'Smoke Bomb',
  'Teleport Stone',
  'Dried Meat',
  'Apocalypse',
  'Zwill Crossblade',
  'Zanmato',
  'Oborozuki',
  'Longinus',
  'Godhand',
  'Save the Queen',
  'Stardust Rod',
  'Angel Brush',
  'Final Trump',
  'Gungnir',
  'Dueling Mask',
  'Scorpion Tail',
  'Bone Wrist',
  'Excalipoor',
  /* cSpell:enable */
];
const orderKey = 'SortingOrder';
const nothingName = 'Nothing';
const orderLength = 288;

function handler() {
  const data = getData();
  const slotIndex = getSlotIndex();
  const slot = data.Saves[slotIndex];

  const inventory = slot.Iv.Inventory;
  names.forEach((n) => {
    if (inventory[n]) return; // at least one item registered
    inventory[n] = 1;
  });

  inventory[orderKey] = Object.keys(inventory).filter((k) => k !== nothingName && k !== orderKey);
  while (inventory[orderKey].length < orderLength) {
    inventory[orderKey].push(nothingName);
  }

  setData(data);
}

// eslint-disable-next-line import/prefer-default-export
export function init() {
  const button = document.getElementById('all-items');
  button.addEventListener('click', handler);
}
