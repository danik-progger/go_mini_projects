package main

import (
	"fmt"
	"net"
	"redis/commands"
	"redis/file"
	"redis/resp"
	"strings"
)

func main() {
	// Create a new server
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("🟢 Listening on port :6379")

	aof, err := file.NewAof("database.aof")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err := aof.Close()
		if err != nil {
			fmt.Println("🔴 Failed to close file")
		}
	}()

	err = aof.Read(func(value resp.Val) {
		command := strings.ToUpper(value.Array[0].Bulk)
		args := value.Array[1:]

		handler, ok := commands.Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			return
		}

		handler(args)
	})
	if err != nil {
		fmt.Println("🔴 Failed to read command")
	}

	// Listen for connections
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("🔴 Failed to close connection")
		}
	}()

	for {
		r := resp.NewResp(conn)
		value, err := r.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		if value.Typ != "Array" {
			fmt.Println("Invalid request, expected array")
			continue
		}

		if len(value.Array) == 0 {
			fmt.Println("Invalid request, expected array length > 0")
			continue
		}

		command := strings.ToUpper(value.Array[0].Bulk)
		args := value.Array[1:]

		writer := resp.NewWriter(conn)
		handler, ok := commands.Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			err = writer.Write(resp.Val{Typ: "String", Str: ""})
			if err != nil {
				fmt.Println("🔴 Failed to write response")
			}
			continue
		}

		if command == "SET" || command == "HSET" {
			err = aof.Write(value)
			if err != nil {
				fmt.Println("🔴 Failed to write to file")
			}
		}

		res := handler(args)
		err = writer.Write(res)
		if err != nil {
			fmt.Println("🔴 Failed to write response")
		}
	}
}
