package main

import (
	"io/ioutil"
	"os/exec"
)

func main() {
	info, err := ioutil.ReadDir("pb")
	if err != nil {
		panic(err)
	}

	for _, a := range info {
		if a.IsDir() {
			path := "pb/" + a.Name() + "/" + a.Name() + ".proto"
			cmd := exec.Command("protoc", "--go_out=.", "--go_opt=paths=source_relative", "--go-grpc_out=.", "--go-grpc_opt=paths=source_relative", path)
			cmd.Run()
		}
	}
}
