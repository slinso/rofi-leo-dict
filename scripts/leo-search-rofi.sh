#!/usr/bin/bash

# load clipboard content (for translating highlighted text)
input=$(xclip -o)

# Call menu as long as exit code is 0 (esc is not pressed)
while [ "$?" -eq "0" ]; do
    # write search to history
    input=$(rofi-leo-dict "$input" | rofi -dmenu -l 30 -p "Leo Dict: ")
done
