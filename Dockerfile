FROM golang:1.11.2

# install bash & curl
RUN apt-get update && apt-get install bash curl -yqq


RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN mkdir -p $GOPATH/src/drone-navigation-service-master
COPY . $GOPATH/src/drone-navigation-service-master
WORKDIR $GOPATH/src/drone-navigation-service-master

RUN dep ensure -vendor-only
RUN rm -rf $GOPATH/src/github.com/cloudqwest/drone-navigation-service-master/testmain

RUN go build

ENV ENV=dev

EXPOSE 5010
ENTRYPOINT ["./drone-navigation-service"]