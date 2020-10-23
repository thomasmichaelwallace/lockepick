import * as editor from './actions/editor';
import * as files from './actions/files';

function init() {
  editor.init();
  files.init();
}
init();
