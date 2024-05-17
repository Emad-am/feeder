package feederservice

import (
	"bufio"
	"dde/config"
	"dde/internal/redis"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func Start() {
	var conf = config.GetConfig()

	host := conf.DDE.Host
	port := conf.DDE.Port
	user := conf.DDE.UserName
	password := conf.DDE.Password

	connString := fmt.Sprintf("%s:%s", host, port)

	conn, err := net.Dial("tcp", connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to server: %v\n", err)
		return
	}

	fmt.Printf("Connected to %s\n", connString)

	reader := bufio.NewReader(conn)

	initialResponse, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading initial response: %v\n", err)
		return
	}
	fmt.Printf("Server: %s", initialResponse)

	_, err = fmt.Fprintf(conn, user+"\n")

	secondResponse, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading initial response: %v\n", err)
		return
	}
	fmt.Printf("Server: %s", secondResponse)

	loginPrompt, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading initial response: %v\n", err)
		return
	}
	fmt.Printf("Server: %s", loginPrompt)

	_, err = fmt.Fprintf(conn, user+"\n")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sending username: %v\n", err)
		return
	}
	fmt.Printf("Sent username: %s\n", user)

	passwordPrompt, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading password prompt: %v\n", err)
		return
	}
	fmt.Printf("Server: %s", passwordPrompt)

	_, err = fmt.Fprintf(conn, password+"\n")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sending password: %v\n", err)
		return
	}
	fmt.Printf("Sent password: %s\n", password)

	authResponse, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading authentication response: %v\n", err)
		return
	}

	fmt.Printf("Authentication response: %s", authResponse)

	go func() {
		for {
			time.Sleep(30 * time.Second) // Adjust the interval as needed
			_, err := fmt.Fprintf(conn, "KEEP_ALIVE\n")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error sending keep-alive message: %v\n", err)
				return
			}
			fmt.Println("Sent keep-alive message")
		}
	}()

	for {
		res, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading authentication response: %v\n", err)
			return
		}

		if res != "" {
			str := trim(res, []string{"\n", "\r"})
			splitted := strings.Split(str, " ")
			redis.Lpush(splitted[0], splitted[1]+" "+splitted[2])
			fmt.Println(res)
		}
	}
}

func trim(s string, chars []string) string {
	for i := 0; i < len(chars); i++ {
		s = strings.ReplaceAll(s, chars[i], "")
	}
	return s
}
