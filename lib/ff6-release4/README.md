# Usage

	adb pull /data/data/com.square_enix.android_googleplay.FFVI/filesRockDat0.sav filesRockDat0.sav.orig
	adb pull /data/data/com.square_enix.android_googleplay.FFVI/filesRockDat1.sav filesRockDat1.sav.orig
	./ff6 tojson -i filesRockDat0.sav.orig -o filesRockDat0.json
	./ff6 tojson -i filesRockDat1.sav.orig -o filesRockDat1.json

Open up and edit filesRockDat0.json and filesRockDat1.json (or just the one with the highest "FileHeader -> Version" number) as you want, then:

	./ff6 fromjson -i filesRockDat0.json -o filesRockDat0.sav
	./ff6 fromjson -i filesRockDat1.json -o filesRockDat1.sav
	adb push /data/data/com.square_enix.android_googleplay.FFVI/filesRockDat0.sav
	adb push /data/data/com.square_enix.android_googleplay.FFVI/filesRockDat1.sav
	adb shell chmod 777 /data/data/com.square_enix.android_googleplay.FFVI/filesRockDat0.sav
	adb shell chmod 777 /data/data/com.square_enix.android_googleplay.FFVI/filesRockDat1.sav

# Command help:

	10:32 ~/code/go/src/github.com/quarnster/ff6 $ ./ff6 help
	Manage Final Fantasy VI mobile save files

	Usage:
		ff6 command [arguments]

	Available commands:
	fromjson  Convert .json save file to .sav
	help      get more information about a command
	tojson    Convert .sav save file to .json

	Use "ff6 help [command]" for more information about a command.

	10:33 ~/code/go/src/github.com/quarnster/ff6 $ ./ff6 help fromjson
	ff6 fromjson

	options
	  -i="filesRockDat0.json": Input .json file
	  -o="filesRockDat0.sav": Output .sav file
	10:33 ~/code/go/src/github.com/quarnster/ff6 $ ./ff6 help tojson
	ff6 tojson

	options
	  -i="filesRockDat0.sav": Input .sav file
	  -o="filesRockDat0.json": Output .json file
	10:33 ~/code/go/src/github.com/quarnster/ff6 $

There are most likely issues and incorrect mappings, pull requests please.

# License

The license of the project is the [2-clause BSD license](https://github.com/quarnster/ff6/blob/master/LICENSE).