Google Go Media Framework is a simple bridge between Golang and ffmpeg.

### HOWTO Install ###
this is only needed when ffmpeg is not installed at default location
```
export CGO_LDFLAGS="-L$FFMPEG_ROOT/lib/ -lavcodec -lavformat -lavutil -lswscale"
export CGO_CFLAGS="-I$FFMPEG_ROOT/include"
```
$FFMPEG\_ROOT must be replaced with the Path to FFMpeg
after setting the environment variables right.
```
goinstall gmf.googlecode.com/hg/gmf
```

### HOWTO Use ###
for a full example see http://code.google.com/p/gmf/source/browse/transcoder.go<br>
or look at the TranscoderTutorial