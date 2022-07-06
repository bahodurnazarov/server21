package main

import (
	"log"
	"net"
	"os"
)

func main() {
	host := "0.0.0.0"
	port := "9999"

	if err := execute(host, port); err != nil {
		os.Exit(1)
	}
}

func execute(host string, port string) (err error) {
	listner, err := net.Listen("tcp", net.JoinHostPort(host, port))
	if err != nil {
		log.Print(err)
		return err
	}
	defer func ()  {
		if cerr := listner.Close(); cerr != nil {
			if err == nil {
				err = cerr
				return
			}
			log.Print(cerr)
		}
	}()
	return
}
