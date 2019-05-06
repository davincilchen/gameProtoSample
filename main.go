package main

import (
	"bufio"
	fmt "fmt"
	"log"
	"os"
	"strings"

	// "io"
	// "io/ioutil"
	// "log"
	// "os"
	// "strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/proto"

	// "github.com/golang/protobuf/proto"
	"github.com/slotProtoSample/server"
	sv "github.com/slotProtoSample/server"
)

func main() {

	r := os.Stdin
	var marshal []byte

	for {
		log.Println("Select command")
		log.Println("Set data: 1")
		log.Println("Get data: 2")
		log.Println("Exit: others")
		rd := bufio.NewReader(r)

		//arr, _, err := rd.ReadLine()
		//cmd := string(arr[:])

		rs, err := rd.ReadString('\n')
		cmd := strings.TrimSpace(rs)

		//cmd = strings.Replace(cmd, "\n", "", -1)
		//cmd = strings.Replace(cmd, "\n", "", -1)

		if err != nil {
			log.Println("Read Error", err, cmd)
		}

		switch cmd {
		case "1":
			dump := false

			log.Println("======================================================================")
			fmt.Println("cmd = 1")

			PW := "5566"
			forcedEnd := true
			game := "Happy Game"

			login := sv.LoginRequest_Dev{}

			login.Id = proto.String("YH")
			login.Password = &PW
			login.ForcedEnd = &forcedEnd
			if dump {
				spew.Dump(login)
			}

			loginD := sv.LoginRequest_Dev_{}
			loginD.Dev = &login

			if dump {
				spew.Dump(loginD)
			}

			loginRequest := sv.LoginRequest{}
			loginRequest.User = &loginD
			loginRequest.Game = &game
			if dump {
				spew.Dump(loginRequest)
			}

			ts := &sv.ToServer{
				Message: &sv.ToServer_LoginRequest{
					LoginRequest: &loginRequest,
				},
			}

			spew.Dump(ts)

			log.Println("GetLoginRequest", ts.GetLoginRequest())
			log.Println("GetGameStartRequest", ts.GetGameStartRequest())
			log.Println("GetResultRequest", ts.GetResultRequest())
			log.Println("GetGameEndRequest", ts.GetGameEndRequest())
			log.Println("GetMessage", ts.GetMessage())

			log.Println("======================================================================")

			out, err := proto.Marshal(ts)
			if err != nil {
				log.Fatalln("Failed to encode [toServer]:", err)
			} else {
				//log.Println("after Marshal: ", out)
				marshal = out
				log.Println("after Marshal: ", marshal)
			}

			log.Println("======================================================================")
			log.Println()
		case "2":

			log.Println("======================================================================")
			fmt.Println("cmd = 2")

			ts := &sv.ToServer{}
			if err := proto.Unmarshal(marshal, ts); err != nil {
				log.Fatalln("Failed to parse [toServer]:", err)
			} else {
				spew.Dump(ts)
			}
			log.Println("======================================================================")

			//if ts.GetMessage() != nil {
			switch x := ts.Message.(type) {
			case *server.ToServer_LoginRequest:

				if x.LoginRequest != nil {
					spew.Dump(x.LoginRequest.GetUser())
				}

				log.Println("======================================================================")

				statusCode := server.StatusCode_Ok
				balance := uint64(1000)
				gameData := "GameData"

				rs := sv.LoginResponse{}

				rs.StatusCode = &statusCode
				rs.Balance = &balance
				rs.GameData = &gameData

				tc := &sv.ToClient{
					Message: &sv.ToClient_LoginResponse{
						LoginResponse: &rs,
					},
				}

				spew.Dump(tc)

			case *server.ToServer_GameStartRequest:

			case *server.ToServer_ResultRequest:

			case *server.ToServer_GameEndRequest:

			case nil:
				log.Println("ToServer.Message has unexpected type :", x)
			default:
				//return fmt.Errorf("ToServer.Message has unexpected type %T", x)
				log.Println("ToServer.Message has unexpected type :", x)
			}

			//}
			log.Println()
		default:
			fmt.Println("cmd = other")
			os.Exit(2)
		}

	}
}
