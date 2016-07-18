#! /bin/bash

find | xargs -i sed -i 's/[ \t]*$//g'  {} 2&> /dev/null
