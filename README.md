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

$ docker-compose up --build
Building goserver
Step 1/9 : FROM golang:latest
 ---> 9fe4cdc1f173
Step 2/9 : RUN mkdir -p /go/src/app/
 ---> Using cache
 ---> 9126518a1750
Step 3/9 : COPY . /go/src/app/
 ---> bcdab0ca8604
Step 4/9 : WORKDIR /go/src/app/
 ---> Running in 4e08458553a0
Removing intermediate container 4e08458553a0
 ---> 732d4cb1c491
Step 5/9 : EXPOSE 8080
 ---> Running in 4279164e57db
Removing intermediate container 4279164e57db
 ---> 91e7d240a774
Step 6/9 : EXPOSE 8081
 ---> Running in f7b080a7d748
Removing intermediate container f7b080a7d748
 ---> 8d7550992197
Step 7/9 : RUN go get -d -v ./...
 ---> Running in a0764d3dba5e
Removing intermediate container a0764d3dba5e
 ---> ce928669cead
Step 8/9 : RUN go install -v ./...
 ---> Running in 7fb491fe3c81
app/factoryMethod
app/singleton
app/main
Removing intermediate container 7fb491fe3c81
 ---> 73293078bbeb
Step 9/9 : CMD ["main"]
 ---> Running in 10d004776482
Removing intermediate container 10d004776482
 ---> ce8ce981e613
Successfully built ce8ce981e613
Successfully tagged soa_golang_goserver:latest
Recreating soa_golang_goserver_1 ... done


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
