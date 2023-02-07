FROM golang:1.20.0 as build
COPY . /usr/src/
WORKDIR /usr/src/
RUN go build -o server main.go

FROM scratch
COPY --from=build /usr/src/server /server
ENV SERVER_PORT=80
ENTRYPOINT ["/server"]
