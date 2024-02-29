package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"github.com/JIEEEN/grpc-board/client/env"
	"github.com/JIEEEN/grpc-board/client/service"
	"github.com/JIEEEN/grpc-board/client/ui"
	"google.golang.org/grpc"
)

func init() {
	env.InitEnv()
}

func main() {
	addr := os.Getenv("GRPC_SERVER_IP") + os.Getenv("GRPC_SERVER_PORT")
	_, w := ui.AppInit()
	w.Resize(fyne.NewSize(800, 600))

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Can't connect to: %v", err)
	}
	service.InitService(conn)
	defer conn.Close()

	ui.MainPage(w)

	w.ShowAndRun()
}
