# Emoji Generator

## Usage

`emojigen` command can generate emoji image for Custom Emoji in Chat Application such as Slack, Mattermost etc. 

## Reize some image

```
$ ./emojigen --type resize --image icon.jpg 

$ ./emojigen --type resize --image icon.jpg  --out ./path/to/output
```

## Generate Text Image

```
./emojigen -type text -string "なるほど" -font sample.ttf 

./emojigen -type text -string "忖　　度" -font sample.ttf 
```

## Generate Annimation Gif

```
./emojigen -type animation -image sample.png -out sample.gif
```

## Known Issue(To Do)

* Resize operation supports only square image. Rectangle dose not work properly.
* Generate Text operation does not fit according to string length and font size.
* Animation operation does not support transparent background.