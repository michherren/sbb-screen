sudo apt update -qq && sudo apt upgrade

sudo apt install chromium-browser unclutter -qq

echo "@xset -dpms" | sudo tee -a /etc/xdg/lxsession/LXDE-pi/autostart
echo "@xset s off" | sudo tee -a /etc/xdg/lxsession/LXDE-pi/autostart
echo "@xset s noblank" | sudo tee -a /etc/xdg/lxsession/LXDE-pi/autostart
echo "@unclutter -idle 0" | sudo tee -a /etc/xdg/lxsession/LXDE-pi/autostart

echo "@sh /home/pi/display/start.sh" | sudo tee -a /etc/xdg/lxsession/LXDE-pi/autostart

