package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go-grpc-k8s-starter-client/proto"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
	"time"
)

var client proto.AddServiceClient

func main() {
	// 1. create a Conn with grpc.Dial
	// 2. create a client which uses the gRPC Conn: proto.NewAddServiceClient(conn)

	conn, err := grpc.Dial("add-service:3000", grpc.WithInsecure())
	//conn ,err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to connect to the add service", err.Error())
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	//defer cancel()
	//
	//a := uint64(21)
	//b := uint64(54)
	//
	//request := &proto.AddRequest{
	//			A: a,
	//			B: b,
	//		}
	//
	//response, err := client.Compute(ctx, request)
	//	if err != nil {
	//		log.Fatal("Encountered an error when attempting to compute addition")
	//		return
	//	}
	//
	//	result := response.GetResult()
	//
	//	fmt.Printf("The result is\n", result)
	routes := mux.NewRouter()
	routes.HandleFunc("/add/{a}/{b}", handleAdd).Methods(http.MethodGet)

	client = proto.NewAddServiceClient(conn)

	fmt.Println("Starting client")
	err = http.ListenAndServe(":8080", routes)
	if err != nil {
		log.Fatal(err)
	}
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")
	vars := mux.Vars(r)
	a, err := strconv.ParseUint(vars["a"], 10, 64)
	if err != nil {
		log.Println("Failed to make a an int")
	}
	b, err := strconv.ParseUint(vars["b"], 10, 64)
	if err != nil {
		log.Println("Failed to make b an int")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := &proto.AddRequest{
		A: a,
		B: b,
	}

	response, err := client.Compute(ctx, request)
	if err != nil {
		log.Fatal("Encountered an error when attempting to compute addition")
		return
	}

	result := response.GetResult()

	fmt.Printf("The result is %d\n", result)

}
