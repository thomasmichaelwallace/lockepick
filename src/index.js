import * as allItems from './actions/allItems';
import * as editor from './actions/editor';
import * as files from './actions/files';
import * as fixItems from './actions/fixItems';
import * as minItems from './actions/minItems';
import * as minLevel from './actions/minLevel';
import * as quickSpell from './actions/quickSpell';

function init() {
  // main
  editor.init();
  files.init();
  // utilities
  minLevel.init();
  quickSpell.init();
  allItems.init();
  fixItems.init();
  minItems.init();
}
init();
