package internal

import (
	"net"
	"fmt"
	"encoding/json"
)

type Response struct {
	Status string
	Response json.RawMessage
}

func Connect() (net.Conn, error) {
	conn, err := net.Dial("unix", "/tmp/rssd.sock");
	if err != nil {
        return nil, fmt.Errorf("failed to create connection: %v", err);
	}
    return conn, nil;
}

func Send(conn net.Conn, command string) error {
	_, err := conn.Write([]byte(command));
	if err != nil {
		return fmt.Errorf("failed to send the command: %v\n", err);
	}
	return nil;
}

func Recv(conn net.Conn) (Response, error) {

	var r Response;

	dec := json.NewDecoder(conn);

	if err := dec.Decode(&r); err != nil {
		return Response{}, err;
	}

	return r, nil;
}


