package main
 
import (
	"context"
	"log"
	"time"
	
	"github.com/Andronovo-bit/go-python-redis-grpc/pkg/proto/user"
	"google.golang.org/grpc"
)
 
func main() {
	log.Println("Client running ...")
 
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
 
	client := user.NewSaverClient(conn)
 
	request := &user.User{
		Id:0, 
		FirstName:"Hasan", 
		LastName:"Kemal",
		Email:"hasn@gmail.com",
		Gender:"erkek",
		IpAddress:"432543543",
		UserName:"hsnkmel",
		Agent:"firefox",
		Country: "Turkey",
		}			  
										

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
 
	response, err := client.SaveUser(ctx, request)
	if err != nil {
		log.Fatalln(err)
	}
 
	log.Println("Response:", response)
}