# Emoji Generator

## Usage

`emojigen` command can generate emoji image for Custom Emoji in Chat Application such asu Slack, Mattermost etc. 

## Reize some image

```
$ ./emojigen --type resize --image icon.jpg 

$ ./emojigen --type resize --image icon.jpg  --out ./path/to/output
```

## Generate Text Image

```
./emojigen -type text -string "疲労困憊" -font sample.ttf 

./emojigen -type text -string "忖　　度" -font sample.ttf 
```

## Generate Annimation Gif

```
./emojigen -type animation -image sample.png -out sample.gif
```