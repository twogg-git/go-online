# go-online
This go examples are ment to run on labs.play-with-docker.com/ 

First start a session in labs.play-with-docker.com/, then "ADD NEW INSTANCE".

### To run directly each example just use the following commands:

Runing hello world 
```sh
docker run --rm golang sh -c \ "go get github.com/twogg-git/go-online/hello/... && exec hello"
```

OutYet file: This is a more complex example, include rest services and an exposed port, in this case 8080.
```sh
docker run -p "8080:8080" --rm golang sh -c \ "go get github.com/twogg-git/go-online/outyet/... && exec outyet"
```

### Runing just hello world example:

#### 1. Dowload build and install hello world file
```sh
docker run golang go get -v github.com/twogg-git/go-online/hello/...
```

#### 2. Built the container into a new image
```sh
docker commit $(docker ps -lq) testgo
```

#### 3. Run the image 
```sh
docker run testgo hello
```

Based on: https://blog.docker.com/2016/09/docker-golang/
