FROM ethereum/solc:stable as solc

FROM golang:1.9.2

COPY --from=solc /usr/bin/solc /usr/bin/solc

RUN go get -u github.com/golang/dep/...
RUN go get -u github.com/ethereum/go-ethereum/...
RUN cd $GOPATH/src/github.com/ethereum/go-ethereum/ && \
    make && make devtools