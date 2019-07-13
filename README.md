# Introduction

Description: The code demonstrates design patterns developed in GoLang and also demonstrates skills in developing
REST APIs in GoLang by following Service Oriented Architecture principles.

# Use

The repository is to practise GoLang development for production use. 

## Abstract

With containerized REST APIs distributed services can be built to provide data persistence, reliability and high availability. In this particular repository we are using Docker, Docker-Compose and Go containers built out of one image.

### Steps

1. Clone this repository
2. Run docker-compose cluster
3. run curl on the only service that is available on <container-ip-address>:8080
4. endpoints are
    1. /singleton
    2. /factorymethod
    3. /post/add
    
    For rest of the interfaces I am building test routines to test respective service. For instance for testing Builder Pattern do the following.
    
    1. run ```docker-compose up --build -d```
    2. run ```winpty docker exec soa_golang_goserver_1 bash -c "cd /go/src/app/builderService; go test"```
    
    Note: in the second step "winpty" is needed because the command was run on Git Bash shell that doesn't provide appropriate terminal byy default

Endpoints demonstrate respective design patterns.

# Output

```

$ docker-compose up --build -d
Building goserver
Step 1/8 : FROM golang:latest
 ---> f50db16df5da
Step 2/8 : RUN mkdir -p /go/src/app/
 ---> Using cache
 ---> eac81f734e94
Step 3/8 : COPY . /go/src/app/
 ---> 58f1bbc8bcc7
Step 4/8 : WORKDIR /go/src/app/
 ---> Running in 5e30e532afcb
Removing intermediate container 5e30e532afcb
 ---> 28d9ade1084b
Step 5/8 : EXPOSE 8080
 ---> Running in 6c533c340028
Removing intermediate container 6c533c340028
 ---> ccccaea13546
Step 6/8 : RUN go get -d -v ./...
 ---> Running in e6a2dad76374
Removing intermediate container e6a2dad76374
 ---> ddb671ae5282
Step 7/8 : RUN go install -v ./...
 ---> Running in 93ee30fc0a80
app/factoryMethod
app/builderService
app/singleton
app/postService
app/main
Removing intermediate container 93ee30fc0a80
 ---> 70fb0260a367
Step 8/8 : CMD ["main"]
 ---> Running in f517de74bc00
Removing intermediate container f517de74bc00
 ---> c1c2b23b812f
Successfully built c1c2b23b812f
Successfully tagged soa_golang_goserver:latest
Creating soa_golang_goserver_1 ... done


```

## Singleton

```
$ docker exec soa_golang_goserver_1 curl 192.168.16.2:8080/singleton
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    45  100    45    0     0  24271      0 --:--:-- --:--:-- --:--:-- 45000

 Received a request singleton2['Amit']=Kumar


```

## Factory Method

```

$ docker exec soa_golang_goserver_1 curl 192.168.16.2:8080/factorymethod
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    60  100    60    0     0  40540      0 --:--:-- --:--:-- --:--:-- 60000

 Default File System
 Main File System
 Block File System



docker exec soa_golang_goserver_1 curl  -X POST -H "Content-Type: application/json" -d @data.json 172.26.0.2:8080/post/add


```

## Builder Pattern

```

$ winpty docker exec soa_golang_goserver_1 bash -c "cd /go/src/app/builderService; go test"
{map[top:Home Page Top] map[bottom:Home Page Bottom] map[left:Home Page Left] map[right:Home Page Right]}PASS
ok      _/go/src/app/builderService     0.002s


```

## Concurrency Patterns

### Maintain local object consensus protocol

In order to understand and practice concurrency patterns, I can build a membership service for an FSM to coordinate among themselves to keep a consensus state in their local memory in a redundant manner.

Algorithm: Paxos algorithm.

How to build consensus on a document data structure by multiple collaborating parties.
