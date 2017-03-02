FROM golang

ENV APP_HOME /go/src/goERP

RUN mkdir -p $APP_HOME

WORKDIR $APP_HOME

ADD . $APP_HOME

RUN go get github.com/tools/godep \
	&& go get github.com/beego/bee \
	&& godep restore

CMD go run main.go