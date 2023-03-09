package server

import (
    "github.com/gorilla/mux"
    "google.golang.org/grpc"
)

func main() error {

    grpc = grpc.NewServer();

    go func () {
        listener, err := net.Listen("tcp", {
            Host: "0.0.0.0",
            port: "8080",
        });
    }
}

