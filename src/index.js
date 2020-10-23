import * as editor from './actions/editor';
import * as files from './actions/files';
import * as minLevel from './actions/minLevel';

function init() {
  editor.init();
  files.init();
  minLevel.init();
}
init();
