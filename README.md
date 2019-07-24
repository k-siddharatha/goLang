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
Creating network "soa_golang_default" with the default driver
Building goserver
Step 1/30 : FROM golang:latest
 ---> be63d15101cb
Step 2/30 : RUN mkdir -p /go/src/app/
 ---> Using cache
 ---> 6c23f1673231
Step 3/30 : COPY builderService /go/src/app/builderService
 ---> Using cache
 ---> ea5821e94328
Step 4/30 : COPY factoryMethod /go/src/app/factoryMethod
 ---> Using cache
 ---> 902ea43fa0bd
Step 5/30 : COPY postService /go/src/app/postService
 ---> Using cache
 ---> a86b674d892b
Step 6/30 : COPY singleton /go/src/app/singleton
 ---> Using cache
 ---> ef55be8f852a
Step 7/30 : COPY disDataDict /go/src/app/disDataDict
 ---> Using cache
 ---> 6078700865d1
Step 8/30 : COPY main /go/src/app/main
 ---> Using cache
 ---> 3d4d6d812d0f
Step 9/30 : EXPOSE 8080
 ---> Using cache
 ---> 659e95c983e6
Step 10/30 : WORKDIR /go/src/app/
 ---> Using cache
 ---> 4737fb7a7e3d
Step 11/30 : RUN go get -d -v go.etcd.io/etcd/client
 ---> Using cache
 ---> d78066090146
Step 12/30 : RUN go install -v go.etcd.io/etcd/client
 ---> Using cache
 ---> 2bb3bd206979
Step 13/30 : RUN go get -d -v ./...
 ---> Using cache
 ---> 5fee467bbefb
Step 14/30 : RUN go install -v ./...
 ---> Using cache
 ---> 9844ccec9dbd
Step 15/30 : RUN mkdir -p /go/src/etcd/
 ---> Using cache
 ---> 08f59f2180ac
Step 16/30 : RUN apt-get update && apt-get -y upgrade
 ---> Using cache
 ---> 1f333d940ef5
Step 17/30 : RUN apt-get install -y build-essential curl
 ---> Using cache
 ---> 477bcd28becc
Step 18/30 : WORKDIR /go/src/
 ---> Using cache
 ---> c10ff1a60a0c
Step 19/30 : ENV ETCD_VER v3.3.13
 ---> Using cache
 ---> a6d18fea1d64
Step 20/30 : ENV GITHUB_URL https://github.com/etcd-io/etcd/releases/download
 ---> Using cache
 ---> 9f23c92b7830
Step 21/30 : ENV DOWNLOAD_URL ${GITHUB_URL}
 ---> Using cache
 ---> 766e02a2b297
Step 22/30 : RUN mkdir -p $GOPATH/bin/
 ---> Using cache
 ---> 5c1f9010a823
Step 23/30 : RUN curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
 ---> Using cache
 ---> 8ea08bc63e85
Step 24/30 : RUN tar xzvf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz -C $GOPATH/bin/ --strip-components=1
 ---> Using cache
 ---> 5d6d53fe0432
Step 25/30 : RUN rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
 ---> Using cache
 ---> 465ea3fb7f8b
Step 26/30 : RUN $GOPATH/bin/etcd --version
 ---> Using cache
 ---> 41b71ff3416f
Step 27/30 : ENV ETCDCTL_API 3
 ---> Using cache
 ---> 5cc525062f68
Step 28/30 : RUN $GOPATH/bin/etcdctl version
 ---> Using cache
 ---> 6f57691b0bd3
Step 29/30 : RUN mkdir -p /var/etcd_data/
 ---> Using cache
 ---> 30f35bcc06d1
Step 30/30 : CMD ["main"]
 ---> Using cache
 ---> 9a259d1da7e9
Successfully built 9a259d1da7e9
Successfully tagged soa_golang_goserver:latest
Building goserver1
.....................<Similar output as above>
Successfully built 9a259d1da7e9
Successfully tagged soa_golang_goserver1:latest
Building goserver2
.....................<Similar output as above>
Successfully built 9a259d1da7e9
Successfully tagged soa_golang_goserver2:latest
Creating soa_golang_goserver2_1 ... done
Creating soa_golang_goserver_1  ... done
Creating soa_golang_goserver1_1 ... done


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

For this purpose I used an opensource project known by the name of "etcd". This is a distributed key-value store that maintains eventual consistency by using RAFT algorithm under the hood. The following test case exemplifies the use of etcd:

```

$ winpty docker exec soa_golang_goserver_1 bash -c "cd /go/src/app/disDataDict/; go test"
2019/07/24 19:22:22 Set is done, metadata is :&{set {Key: /foo, CreatedIndex: 8, ModifiedIndex: 8, TTL: 0} <nil> 8 bd16ad9b711e9502}
2019/07/24 19:22:22 Get is done, metadata is :&{get {Key: /foo, CreatedIndex: 8, ModifiedIndex: 8, TTL: 0} <nil> 8 bd16ad9b711e9502}
2019/07/24 19:22:22 "/foo" key has "bar" value
PASS
ok      app/disDataDict 0.060s


```
