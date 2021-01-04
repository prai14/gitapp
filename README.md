## Go OAuth GitHub example
This is an example Microservice golang application that implements Githubs OAuth2 API.
All services are running behind resgate Realtime API Gateway.
https://nats.io/blog/resgate_nats/

## In order to run the application:
##### Run the resgate Realtime API Gateway Server + NATS 

$ sudo docker network create res \
$ sudo docker run -d --name nats -p 4222:4222 --net res nats \
$ sudo docker run --name resgate -p 8080:8080 --net res resgateio/resgate --nats nats://nats:4222 


### Run OAuth server + User Service
added both module in a single app. \
$ got get github.com/prai14/gitapp \
$ cd gitapp \
$ go build \
$ ./gitApp OR \
$ go run . 

## open link in browser http://localhost:8090

## Run in Docker -
$ cd gitapp \
$ docker build -t gitapp . \
$ docker run -p 8090:8090 -it my-go-app
