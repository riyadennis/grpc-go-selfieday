FROM golang
ADD . github.com/riyadennis/grpc-go-selfieday
RUN go install github.com/riyadennis/grpc-go-selfieday
ENTRYPOINT ["/go/bin/server"]
EXPOSE 5052