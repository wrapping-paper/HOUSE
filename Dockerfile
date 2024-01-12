FROM golang:1.21

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o awesome_project ./main.go

CMD ["./awesome_project"]