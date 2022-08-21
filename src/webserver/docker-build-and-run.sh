sudo docker rm -f $(sudo docker ps -aq)
sudo docker rmi -f $(sudo docker images -aq 'go-webserver')
sudo docker build -t go-webserver .
sudo docker run -p 8080:8080 go-webserver
