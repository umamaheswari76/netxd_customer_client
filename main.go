package main

import (
	"context"
	"fmt"
	"log"

	cst "github.com/umamaheswari76/netxd_customer_proto/customer"
	
	"google.golang.org/grpc"
)

func main() {

	//Dial creates a client connection to the given target.
	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	client := cst.NewCustomerServiceClient(conn)

	response, err := client.CreateCustomer(context.Background(), &cst.Customer{
		CustomerId: 101,
		FirstName:  "umamaheswari",
		SecondName: "m",
		BankId:     "1",
		Balance:    5000,
	})
	if err != nil {
		log.Fatalf("Failed to call CreateCustomer: %v", err)
	}

	fmt.Printf("Response: %v\n", response.CustomerId)

}
