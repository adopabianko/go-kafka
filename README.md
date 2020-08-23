## How to run

```bash
git clone https://github.com/adopabianko/go-kafka.git

# Add kafka to your hosts
sudo nano /etc/hosts
127.0.0.1 kafka

# go to project
cd go-kafka

# create kafka network
docker network create kafka

# running docker compose
docker-compose up

# build and run consumer with command
go build -o cmd/consumer/consumer cmd/consumer/main.go
./consumer

# buid and run publisher with command
go build -o cmd/publisher/publisher cmd/publisher/main.go
./publisher