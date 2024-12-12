#!/bin/bash

$(dirname "$0")/display-server &

sleep 2

DISPLAY=:0 chromium-browser --kiosk --no-first-run "http://localhost:3000" &
