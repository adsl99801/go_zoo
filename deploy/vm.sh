timedatectl set-timezone Asia/Taipei

# /etc/systemd/system/enable-console-blanking.service
sudo cp deploy/enable-console-blanking.service /etc/systemd/system/
#PostgreSQL
# sudo sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
# wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
# sudo apt-get update
# sudo apt-get -y install postgresql-13

#go
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt-get update
sudo apt install golang-go

#nginx
sudo apt install curl gnupg2 ca-certificates lsb-release
echo "deb http://nginx.org/packages/ubuntu $(lsb_release -cs) nginx" |
    sudo tee /etc/apt/sources.list.d/nginx.list
sudo apt update
sudo apt install nginx
sudo cp deploy/go_zoo.conf /etc/nginx/nginx.conf/
sudo systemctl reload nginx
sudo systemctl start nginx

scp -r web/dist/puppysfun keith@192.168.0.104:/home/keith/puppys.fun
sudo cp -r puppys.fun /var/www/

cp deploy/godaddy.sh /etc/cron.daily/godaddy
cp deploy/limits.conf /etc/security/limits.conf
cp deploy/sshd_config /etc/ssh/sshd_config
#git clone https://github.com/adsl99801/go_zoo
#bash go_zoo/deploy/vm.sh
# 環境需要4G
cd /home/keith/go_zoo
rm /etc/nginx/sites-enabled/default
cp deploy/go_zoo.conf /etc/nginx/conf.d
systemctl reload nginx
cp deploy/go_zoo.service /lib/systemd/system/go_zoo.service
env GOOS=linux GOARCH=amd64 go build .
systemctl start go_zoo
systemctl enable go_zoo
