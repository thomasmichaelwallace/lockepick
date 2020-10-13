# Usage

	adb pull /data/data/com.square_enix.android_googleplay.FFVI/filesRockDat0.sav
	adb pull /data/data/com.square_enix.android_googleplay.FFVI/filesRockDat1.sav
	./ff6 tojson

Open up and edit filesRockDat0.json as you want, then:

	./ff6 fromjson
	adb push /data/data/com.square_enix.android_googleplay.FFVI/filesRockDat0.sav
	adb push /data/data/com.square_enix.android_googleplay.FFVI/filesRockDat1.sav

There are most likely issues and incorrect mappings, pull requests please.

# License

The license of the project is the [2-clause BSD license](https://github.com/quarnster/ff6/blob/master/LICENSE).
