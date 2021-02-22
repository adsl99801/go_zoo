#PostgreSQL
# Create the file repository configuration:
sudo sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
# Import the repository signing key:
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
sudo apt-get update
# Install the latest version of PostgreSQL.
# If you want a specific version, use 'postgresql-12' or similar instead of 'postgresql':
sudo apt-get -y install postgresql-13

#go
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt-get update
sudo apt install golang-go

#nginx
sudo apt install curl gnupg2 ca-certificates lsb-release
echo "deb http://nginx.org/packages/ubuntu `lsb_release -cs` nginx" \
    | sudo tee /etc/apt/sources.list.d/nginx.list
sudo apt update
sudo apt install nginx

#git clone https://github.com/adsl99801/go_zoo
#bash go_zoo/deploy/vm.sh
cd /root/go_zoo
rm /etc/nginx/sites-enabled/default
cp deploy/go_zoo.conf /etc/nginx/conf.d
systemctl reload nginx
cp deploy/go_zoo.service /lib/systemd/system/go_zoo.service
env GOOS=linux GOARCH=amd64 go build .
systemctl start go_zoo
systemctl enable go_zoo
