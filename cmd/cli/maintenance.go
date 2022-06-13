package main

import (
	"fmt"
	"net/rpc"
	"os"

	"github.com/fatih/color"
)

func maintenanceMode(enable bool) error {
	client, err := rpc.Dial("tcp", fmt.Sprintf("127.0.0.1:%s", os.Getenv("RPC_PORT")))
	if err != nil {
		return err
	}
	var result string
	if err := client.Call("RPCServer.MaintenanceMode", enable, &result); err != nil {
		return err
	}
	color.Yellow(result)
	return nil
}
