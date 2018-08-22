FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./app /go/src/github.com/project/app
WORKDIR /go/src/github.com/project/app

RUN go get -v ./
RUN go build -v

CMD if [ -n "$APP_ENV" ] && [ $APP_ENV = "production" ]; then \
	app; else \
	echo APP_ENV: $APP_ENV;\
	go get github.com/pilu/fresh && \
	fresh; \
	fi

EXPOSE 8080