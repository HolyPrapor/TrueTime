#!/bin/bash

begin=$(date --date="0:00" +%s)
end=$(date --date="12:00" +%s)

now=$(date +%s)

AM_PLAYLIST="https://www.youtube.com/watch?v=Tf3pHuoDaEc&list=PL6DfM86Tqa164MVESueWmX4tgdItsIJUc&t="
PM_PLAYLIST="https://www.youtube.com/watch?v=1MuREToGCbY&list=PL6DfM86Tqa164MVESueWmX4tgdItsIJUc&t="

if [ "$begin" -le "$now" -a "$now" -lt "$end" ]; then
    t=$(($(date '+(%H*60+%M)*60+%S')))
    echo "$AM_PLAYLIST""$t"
else
    t=$(($(date '+((%H-12)*60+%M)*60+%S')))
    echo "$PM_PLAYLIST""$t"
fi
