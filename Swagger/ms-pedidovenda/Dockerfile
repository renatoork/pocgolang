FROM golang
MAINTAINER Renato Arruda <renato@mega.com.br>
LABEL description="Exposição da documentação da API."

RUN apt-get update
RUN apt-get install -y pkg-config
RUN apt-get install -y libaio1

RUN go get -u github.com/mgutz/logxi/v1
RUN go get -u github.com/fvbock/endless

ADD /PDV-API /usr/PDV-API

WORKDIR /usr/PDV-API

ENTRYPOINT go run web.go docs.go

EXPOSE 80
