# Locke Pick

An online FF6 PC (Steam) save editor.

## Usage

1. Open the web application: https://www.thomasmichaelwallace.com/lockepick/
1. On the side bar use the 'Upload file' button to load your FF6 save file (probably at: `C:/Program Files (x86)/steam/userdata/{userid}/382900/local`)
1. Using the JSON editor make the changes you want to the file:
    1. By editing the JSON objects directly
    1. By using the utilities on the selected save slot
    1. By switching to advanced mode and editing the unknown bytes (uint8)
1. One finished, use the 'Save file' button to download your altered save file.
1. **Make a backup of your original file before replacing it with the altered version!** (nothing is perfect, and this approach definately isn't!)

At the moment there is no validation, so it is up to you to alter the options in a sensible way. If you're looking for some guidance, check out this [save file documentation.](https://www.ff6hacking.com/wiki/doku.php?id=ff6a:doc:savefile)

## Contributing

Contributions are welcome, however there are two parts to this application:
 * `index.html` and `./src` / `./css` form the interface, with [bulma](https://bulma.io/) and [jsoneditor](https://github.com/josdejong/jsoneditor) doing most of the heavy lifting.
 * `./lib` contains a [gopherjs](https://github.com/gopherjs/gopherjs) transpiled version of a modified (bug-fixed) version of [quarnster's original FF6 save editor](https://github.com/quarnster/ff6). To make modifications to the editing capability, it's easiest to make them directly in `go` by updating `lib/ff6-release4`, transpiling, and then making the ui changes required to surface them.


## Prior Art

This project was made possible by the work of the following (and the people they, in-turn, were helped by):
 * [Quarnster's FF6 save editor](https://github.com/quarnster/ff6) for the initial work making a working to-json editor (especially with checksum).
 * [Experience Guide by mog255](https://gamefaqs.gamespot.com/ps/562865-final-fantasy-vi/faqs/39433) for the level/hp/mp/exp evolution curves.
 * [Esper Guide By Nerd House](https://steamcommunity.com/sharedfiles/filedetails/?id=581570122) for the list of spells by esper.
