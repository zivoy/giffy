#!/bin/bash

org=$1

if [ -z $org ]
then
	echo "Usage:"
	echo "pathToBlendFile fps"
	exit
fi

file=$(wslpath -a $1)
dir=$(dirname $file)
loc=$(wslpath -m $dir)
name=$(basename $file)
striped=${name%.*}

shift
delay=$1
if [ -z $delay ]
then
    delay=4
else
	delay=$((100/$delay))
fi

cd $dir
mkdir "$striped-frames"
echo rendering frames
"/mnt/c/Program Files/Blender Foundation/Blender 2.91/blender.exe" -b $org -o "$loc/$striped-frames/frame_####" -F PNG -a &>/dev/null
#blender.exe -b $org -o "$loc/$striped-frames/frame_####" -F PNG -a &>/dev/null
cd "$striped-frames"
echo converting files to gifs
mogrify -format gif *.png
echo compositing final gif
gifsicle -O3 --delay=$delay --colors 256 --loop *.gif > "../$striped.gif"
echo cleaning up
cd ..
rm -rf "$striped-frames"
#read -p "Press [Enter] key"