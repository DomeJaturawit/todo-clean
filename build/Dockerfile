FROM golang:1.19
# Set environments
ENV APPPATH $GOPATH/src/todo-clean
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

# Set working directory
WORKDIR $APPPATH

COPY ./ .

RUN go mod tidy && go mod download

RUN go build -o /main ./cmd/main.go

EXPOSE 8080

CMD ["/main"]