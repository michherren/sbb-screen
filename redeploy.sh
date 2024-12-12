#!/bin/bash
export TARGET_HOST=sbb-screen

GOOS=linux GOARCH=arm go build -ldflags="-X 'main.DisplayApiKey=$DISPLAY_API_KEY' -X 'main.TrainStationId=$TRAIN_STATION_ID'" -o build/display-server main.go

rsync -r build/ pi@${TARGET_HOST}:/home/pi/display

ssh pi@${TARGET_HOST} 'pkill -f -- "chromium-browser"' || true
ssh pi@${TARGET_HOST} 'pkill -f -- "display-server"' || true
ssh pi@${TARGET_HOST} 'nohup /home/pi/display/start.sh > /dev/null 2>&1 &'
