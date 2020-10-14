const ff6 = require('./ff6-release4');
const fs = require('fs');

function loop() {
  const data = fs.readFileSync('./verify/data');
  const dataBytes = new Int8Array(data);

  const [jsonBytes] = ff6.tojsonbytes(dataBytes);
  const json = JSON.parse(Buffer.from(jsonBytes).toString());
  fs.writeFileSync('./verify/verify.json', JSON.stringify(json, null, 2));

  const [verify] = ff6.fromJsonBytes(jsonBytes)
  fs.writeFileSync('./verify/verify.sav', verify);
}
loop();