FROM golang as build-stage
WORKDIR /go/
COPY main.go /go/
RUN cd /go/ && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o logz

FROM alpine
COPY --from=build-stage /go/logz /
#RUN mkdir -p /mnt/logs
CMD ["/logz"]