# Picture_perfect


# Backend
Backend server is running on aws and and accesable at ....

Either you can build an image and deploy it using docker
```
cd backend/
docker build ./ -t <image-name>
docker run -d -p 80:8080 <image-name>
```
Now server is running you can check at localhost:8080/

run backend server locally:

Include your directory in $GOPATH
```
cd backend/
go get -u github.com/golang/dep/cmd/dep
dep ensure
go run *.go
```





# Frontend

to start frontend server run
```
cd react-bs
npm install
npm start
open localhost:3000 at browser
```


## Links
[Design Document](https://docs.google.com/document/d/1nzsfwW0Onqwrlb249paajzt2BtgkCGzmbJR4hglbJ3Q/edit?usp=sharing "Design Document")

[API definitions](https://app.swaggerhub.com/apis-docs/pawan475/picture_perfect/1.0.0 "Swaggerhub")



[Docker Images](https://hub.docker.com/repository/docker/pawan47/picture_perfect)



