package main

import (
	"log"
	"net/http"

	"service_a_grpc/pb" // Update with your module path

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("service-b.goapp.svc.cluster.local:80", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Service B: %v", err)
	}
	defer conn.Close()

	client := pb.NewMyServiceClient(conn)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("handle func")

		// Call Service B using gRPC
		response, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "Service A"})
		if err != nil {
			log.Printf("gRPC call failed: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Println("resp")
		w.Write([]byte(response.Message))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
