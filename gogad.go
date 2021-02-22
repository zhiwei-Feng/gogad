package main

import (
	"encoding/csv"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
	"time"
)

type Machine struct {
	Host     string // machine ip address
	User     string // ssh login username
	Password string // ssh login password
}

func readCSVToMachine(path string) []Machine {
	fs, err := os.Open(path)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()

	machines := make([]Machine, 0, 5)
	r := csv.NewReader(fs)
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		m := Machine{
			Host:     row[0],
			User:     row[1],
			Password: row[2],
		}
		machines = append(machines, m)
	}
	return machines
}

func main() {
	servers := readCSVToMachine("machines.csv")
	port := 22
	for _, server := range servers {
		// ssh login
		config := &ssh.ClientConfig{
			Timeout:         5 * time.Second,
			User:            server.User,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
		}
		config.Auth = []ssh.AuthMethod{ssh.Password(server.Password)}

		// build dial
		addr := fmt.Sprintf("%s:%d", server.Host, port)
		sshClient, err := ssh.Dial("tcp", addr, config)
		if err != nil {
			log.Printf("server %s connect fail\n", server.Host)
			continue
		}
		defer sshClient.Close()

		// create ssh session
		session, err := sshClient.NewSession()
		if err != nil {
			log.Printf("server %s connect fail\n", server.Host)
			continue
		}
		defer session.Close()

		// exec nvidia-smi to get information of GPU
		combo, err := session.CombinedOutput("nvidia-smi --query-gpu=memory.used,memory.total --format=csv,noheader")
		if err != nil {
			log.Printf("server %s exec nvidia-smi fail\n", server.Host)
			continue
		}

		fmt.Printf("machine:[%s]->memory used, memory total\n%s", server.Host, string(combo))
	}
}
