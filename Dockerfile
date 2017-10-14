FROM golang:1.8
WORKDIR /go/src/starter-golang
COPY . /go/src/starter-golang
RUN go install 
EXPOSE 8000
ENTRYPOINT [ "/go/bin/starter-golang"]
