FROM golang:alpine as builder
ARG MS_NAME=backend-challenge-transfeera
WORKDIR /package

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux go build -ldflags="-s -w" -o ${MS_NAME} ./cmd/main.go

FROM alpine
ARG MS_NAME=backend-challenge-transfeera
ENV MS_NAME_BIN=$MS_NAME 

WORKDIR /usr/app/
COPY --from=builder /package/${MS_NAME} .

EXPOSE 8080

RUN chmod +x ${MS_NAME}
CMD  ./${MS_NAME_BIN}
