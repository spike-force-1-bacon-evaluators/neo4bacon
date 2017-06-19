package api

import (
	"log"
	"net"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

type server struct{}

// List calls the backend server to retrieve a new restaurant raw list
func (server) List(context.Context, *Empty) (*RestaurantList, error) {

	// restaurants, err := Run()
	// if err != nil {
	// 	return nil, fmt.Errorf("could not list restaurants: %s", err)
	// }

	// r := &RestaurantList{
	// 	Restaurants: restaurants,
	// }
	// return r, nil

	// Mock
	restaurant := &Restaurant{
		ID:      "id",
		Name:    "name",
		Lastpos: int32(2),
		Newpos:  int32(1),
		Change:  "change",
	}

	r := &RestaurantList{
		Restaurants: []*Restaurant{restaurant},
	}
	return r, nil
}

// Serve tcp connection
func Serve() {

	log.Println("serving...")

	port := ":50051"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
	}

	s := grpc.NewServer()
	server := server{}
	RegisterNeo4BaconServer(s, server)

	if err := s.Serve(lis); err != nil {
		log.Fatal("could not serve: ", err)
	}
}
