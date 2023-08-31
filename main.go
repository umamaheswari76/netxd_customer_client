package main

import (
	"context"
	"fmt"
	"log"

	cst "github.com/umamaheswari76/netxd_customer_proto/customer"
	tsn "github.com/umamaheswari76/netxd_customer_proto/transaction"
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
	transaction_client := tsn.NewTransactionServiceClient(conn)

	response, err := client.CreateCustomer(context.Background(), &cst.Customer{
		CustomerId: 102,
		FirstName:  "aaaaa",
		SecondName: "m",
		BankId:     "1",
		Balance:    6000,
	})
	if err != nil {
		log.Fatalf("Failed to call CreateCustomer: %v", err)
	}

	//calling transaction
	transaction_response, err := transaction_client.Transfer(context.Background(), &tsn.Transaction{
		Fromaccount: 101,
		Toaccount:   102,
		Amount:      500,
	})
	if err!= nil{
		log.Fatalf("Failed to call Transafer: %v", err)
	}

	fmt.Printf("Response: %v\n", response.CustomerId)
	fmt.Printf("Response: %v\n", transaction_response)
}
