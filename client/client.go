package main

import (
	"io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/ihippik/grpc-test/protocol"
)

const (
	address = "localhost:50051"
)

// createCustomer calls the RPC method CreateCustomer of CustomerServer
func createCustomer(client pb.CustomerClient, customer *pb.CustomerRequest) {
	resp, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatalf("Could not create Customer: %v", err)
	}
	if resp.Success {
		log.Printf("A new Customer has been added with id: %d", resp.Id)
	}
}

// getCustomers calls the RPC method GetCustomers of CustomerServer
func getCustomers(client pb.CustomerClient, filter *pb.CustomerFilter) {
	// calling the streaming API
	stream, err := client.GetCustomers(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get customers: %v", err)
	}
	for {
		customer, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetCustomers(_) = _, %v", client, err)
		}
		log.Printf("%d",customer.Id)
	}
}
func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// Creates a new CustomerClient
	client := pb.NewCustomerClient(conn)

	customer := &pb.CustomerRequest{
		Id:    101,
		Name:  "Иванов Сергей",
		Email: "ivanov@xyz.com",
		Phone: "732-757-2923",
		Addresses: []*pb.CustomerRequest_Address{
			&pb.CustomerRequest_Address{
				Street:            "Брежнева",
				City:              "Смоленск",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: false,
			},
			&pb.CustomerRequest_Address{
				Street:            "Ленина",
				City:              "Калуга",
				State:             "KL",
				Zip:               "68356",
				IsShippingAddress: true,
			},
		},
	}

	// Create a new customer
	createCustomer(client, customer)

	customer = &pb.CustomerRequest{
		Id:    102,
		Name:  "Мазулин Виктор",
		Email: "mazulin@xyz.com",
		Phone: "732-757-2924",
		Addresses: []*pb.CustomerRequest_Address{
			&pb.CustomerRequest_Address{
				Street:            "Юных ленинцев",
				City:              "Ульяновск",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: true,
			},
		},
	}

	// Create a new customer
	createCustomer(client, customer)
	// Filter with an empty Keyword
	filter := &pb.CustomerFilter{Keyword: "Виктор"}
	getCustomers(client, filter)
}