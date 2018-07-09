# Emoji Generator

`emojigen` command can generate emoji image for Custom Emoji in Chat Application such as Slack, Mattermost etc. 

## Build

1. Get Libraries to build

 ```bash
 go get github.com/nfnt/resize
 go get github.com/golang/freetype
 go get github.com/soniakeys/quant/median
 go get github.com/BurntSushi/graphics-go/graphics
```
2. Build Command

```
go build -o emojigen main.go
```

## Usage

### Reize some image

```
$ ./emojigen -type resize -image sample.png 

$ ./emojigen -type resize -image sample.png  -out ./path/to/output.png
```

### Generate Text Image

```
$ ./emojigen -type text -string "なるほど" -font sample.ttf 

$ ./emojigen -type text -string "忖　　度" -font sample.ttf -out ./path/to/output.png
```

### Generate Annimation Gif

```
$ ./emojigen -type animation -image sample.png -out sample.gif
```

## Known Issue(To Do)

* emojine support only "png" file.
* Resize operation supports only square image. Rectangle dose not work properly.
* Generate Text operation does not fit size according to string length and font size.
* Animation operation support only rotation.
* Animation operation does not support transparent background.