
FROM golang

RUN mkdir /go/src/app1

RUN go get -u github.com/golang/dep/cmd/dep


ADD ./main.go /go/src/app1
ADD ./iam.go /go/src/app1
ADD ./models.go /go/src/app1
ADD ./rating.go /go/src/app1
ADD ./review.go /go/src/app1
ADD ./utils.go /go/src/app1
ADD ./catalogue.go /go/src/app1

COPY ./Gopkg.toml /go/src/app1

WORKDIR /go/src/app1

RUN dep ensure
# RUN go test -v
RUN go build 

EXPOSE 8080

CMD ["./app1"]