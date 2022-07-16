sudo cp deploy/go_zoo.conf /etc/nginx/conf.d/
sudo nginx -t
sudo nginx -s reload
sudo systemctl stop go_zoo
go build
