# go-online
You can run this hello world example inside labs.play-with-docker.com/ just follow this steps: 

#### 1. Dowload build and install hello world file
```sh
docker run golang go get -v github.com/twogg-git/go-online/...
```

#### 2. Built the container into a new image
```sh
docker commit $(docker ps -lq) testgo
```

#### 3. Run the image 
```sh
docker run testgo go-online
```

Based on: https://blog.docker.com/2016/09/docker-golang/
