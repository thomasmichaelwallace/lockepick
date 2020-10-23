import schema from './schema.json';

function matchesPath(match, path) {
  if (!Array.isArray(path) || typeof path[0] !== 'number') return false;
  const bits = match.split('.');
  if (bits.length !== path.length) return false; // no partial
  return bits.every((b, i) => {
    if (path[i] === undefined) return false;
    if (b === '*') return path[i] !== undefined;
    return path[i].toString().toLowerCase() === b.toLowerCase();
  });
}

const equipPaths = [
  'RightHand', 'LeftHand', 'Helmet', 'Body', 'Relic1', 'Relic2',
].map((p) => `*.characters.*.${p}`);

const itemNames = Object.keys(
  schema.definitions.save.properties.Inventory.properties,
).filter((k) => k !== 'Nothing' && k !== 'SortingOrder');

export default function getOptions(text, path, input) {
  if (input === 'value') {
    if (
      matchesPath('*.characters.*.id', path)
      || matchesPath('*.characters.*.portrait', path)
    ) {
      return schema.definitions.characterId.anyOf[0].enum;
    }
    if (matchesPath('*.characters.*.command.*', path)) {
      return schema.definitions.character.properties.Command.items.enum;
    }
    if (matchesPath('*.characters.*.esper', path)) {
      return schema.definitions.character.properties.Esper.enum;
    }
    if (equipPaths.some((p) => matchesPath(p, path))) {
      return schema.definitions.equipment.enum;
    }
  } else if (input === 'field' && matchesPath('*.Inventory.*', path)) {
    return itemNames;
  }
  return null;
}
