# Custom SBB Screen
A weekend project to play around with SBB apis and Tailscale.
The CI/CD pipeline directly deploys the display-server onto a Raspberry PI through Github-Actions.

## 1. prepare hardware
use Rasperry PI Imager: https://www.raspberrypi.com/software/
```
Raspberry Pi OS
hostname: sbb-screen
user: pi

wlan config: to your network
enable ssh service with key
```
## 2. get access to sbb data
1. register an api on apim-portal with a swisspass login: https://ki-bahnhof-anywhere.app.sbb.ch/#/howto
2. use the api-key in https://ki-bahnhof-anywhere.app.sbb.ch to configure your display
3. configure the GM for your location
4. note down the DISPLAY_API_KEY and the TRAIN_STATION_ID from your url (pattern: https://ki-bahnhof-anywhere.app.sbb.ch/#/stream/{{TRAIN_STATION_ID}}?layout=gm&vmGruppenFilter=FV_RV&api_key={{DISPLAY_API_KEY}}))
5. add these to your local .env file
6. add two secrets into your gihub project

## 3. first boot
1. plugin power and a display to the raspberry pi
2. from another device connect through ssh
```
ssh pi@IP
```
3. install Tailscale for secure and managed access (https://tailscale.com/)
```
curl -fsSL https://tailscale.com/install.sh | sh
sudo tailscale up --ssh
# do not forget to approve and disable key expiry in tailscale
```
After installing tailscale, reconnect through tailscale and disable default ssh
```
ssh pi@sbb-screen
systemctl stop ssh.service
systemctl disable ssh.service
```

4. enable auto-login and switch to X11
```
sudo raspi-config
```
* Go to: System Boot Options > Desktop Autologin
* Go to: Advanced Options > Wayland > X11

X11 is used to support older raspberry-pi versions as well. 

5. update and install dependencies
```
ssh pi@sbb-screen 'bash -s' < install.sh
```

## 4. setup ci/cd with tailscale
1. install tailscale on your computer (tailscale.com)
2. configure your acls to allow ssh: https://tailscale.com/kb/1193/tailscale-ssh
3. setup secrets for the github action (.github/workflows/deploy.yml): https://github.com/tailscale/github-action

## 5. run
1. make sure the env variable are set (see .env_template)
```
go run main.go
```
open http://localhost:3000

2. deploy to raspberry
```
./redeploy.sh
```

