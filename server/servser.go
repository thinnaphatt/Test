package main

import(
	"fmt"
	"net"
)

func main(){
	listener, err := net.Listen("tcp", ":5000")

	if  err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	fmt.Println("GG NOOB")

	for{
		conn, err := listener.Acce
	}

}