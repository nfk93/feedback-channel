# syntax=docker/dockerfile:1

FROM golang:1.16-alpine as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /feedback-channel

##
## Deploy
##

FROM scratch

COPY --from=build /feedback-channel /feedback-channel

EXPOSE 8080

ENTRYPOINT ["/feedback-channel"]