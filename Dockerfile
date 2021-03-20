FROM golang:1.16-buster as go
RUN wget https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz 
RUN tar -xf migrate.linux-amd64.tar.gz
RUN mv migrate.linux-amd64 /usr/local/bin/migrate 
RUN rm migrate.linux-amd64.tar.gz
RUN go get github.com/kyleconroy/sqlc/cmd/sqlc
COPY . ./src/aptitude_bulb
RUN cd src/aptitude_bulb && go get && go test && go build .

FROM gcr.io/distroless/base-debian10:latest
ENV PATH=/usr/local/go/bin:/go/bin:$PATH
COPY --from=go /usr/local/go /usr/local/go
COPY --from=go /go /go
COPY --from=go /usr/local/bin /usr/local/bin
COPY --from=go /go/src/aptitude_bulb/aptitude_bulb .
CMD ["./aptitude_bulb"]
