package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
	"strings"

	"github.com/Andronovo-bit/go-python-redis-grpc/pkg/proto/user"

	"google.golang.org/grpc"
)

var jsonByteVal []uint8

type Users struct {
	Users []User
}

type User struct {
	ID        int32  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Ip        string `json:"ip_address"`
	Username  string `json:"user_name"`
	Agent     string `json:"agent"`
	Country   string `json:"country"`
}

func jsonToProtobuff(jsonName string) Users {
	
	var users2 Users

	json.Unmarshal(jsonByteVal, &users2)

	jsonFile, err := os.Open("./users_json/" + jsonName)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened json file")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var countKey = len(byteValue)

	fmt.Println(strconv.Itoa((countKey)))

	var users Users

	json.Unmarshal(byteValue, &users)

	return users

}

func clientRun() *grpc.ClientConn {
	log.Println("Client running ...")

	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}

	return conn
}

func main() {

	conn := clientRun()

	client := user.NewSaverClient(conn)

	files,_ := ioutil.ReadDir("./users_json")

	for  i,item := range files{
		if strings.Contains(item.Name(), "json"){
			var users = jsonToProtobuff(item.Name())

			for index, element := range users.Users {

				request := &user.User{
					Id:        element.ID,
					FirstName: element.FirstName,
					LastName:  element.LastName,
					Email:     element.Email,
					Gender:    element.Gender,
					IpAddress: element.Ip,
					UserName:  element.Username,
					Agent:     element.Agent,
					Country:   element.Country,
				}

				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				response, err := client.SaveUser(ctx, request)
				if err != nil {
					log.Fatalln(err)
				}

				log.Println("First-index:", i)
				log.Println("index:", index)
				log.Println("Response:", response)

			}
		}
	}
	defer conn.Close()
}
