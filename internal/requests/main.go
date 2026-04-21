package requests

import (
	"fmt"
	"strings"
	"encoding/json"

	"rssc/internal/rssd"
)

type GetResponse struct {
	Id int
	Title string
	Url string
	Content string
	Updated string
	Read bool
	FeedId int
}

func sendCommand[T any](req string) (T, error) {
	var result T;

	conn, err := rssd.Connect();
	if err != nil {
		return result, err;
	}

	err = rssd.Send(conn, req);
	if err != nil {
		return result, err;
	}

	resp, err := rssd.Recv(conn);
	if err != nil {
		return result, err;
	}

	if resp.Status != "yes" {
		return result, fmt.Errorf("request failed: %v", resp.Response);
	}

	err = json.Unmarshal(resp.Response, &result);
	if err != nil {
		return result, err;
	}

	return result, nil;

}

func GetCommand(source string, limit int) ([]GetResponse, error) {
	var req strings.Builder;

	req.WriteString("GET ");
	req.WriteString(source);

	if limit != 0 {
		fmt.Fprintf(&req, " %v", limit)
	}
	req.WriteString("\n");

	res, err := sendCommand[[]GetResponse](req.String());
	if err != nil {
		return nil, err;
	}
	return res, nil;
}

func FindCommand(text string, limit int) ([]int, error) {
	var req strings.Builder;

	req.WriteString("FIND ");
	req.WriteString(text);

	if limit != 0 {
		fmt.Fprintf(&req, " %v", limit)
	}
	req.WriteString("\n");

	res, err := sendCommand[[]int](req.String());
	if err != nil {
		return nil, err;
	}
	return res, nil;
}

func ReadCommand(id string) (string, error) {
	req := fmt.Sprintf("READ %v\n", id);

	res, err := sendCommand[string](req);
	if err != nil {
		return "", err;
	}
	return res, nil;
}

func UnreadCommand(id string) (string, error) {
	req := fmt.Sprintf("UNREAD %v\n", id);

	res, err := sendCommand[string](req);
	if err != nil {
		return "", err;
	}
	return res, nil;
}

func DeleteCommand(id string) (string, error) {
	req := fmt.Sprintf("DELETE %v\n", id);

	res, err := sendCommand[string](req);
	if err != nil {
		return "", err;
	}
	return res, nil;
}

func OpenCommand(id string) (string, error) {
	req := fmt.Sprintf("OPEN %v\n", id);

	res, err := sendCommand[string](req);
	if err != nil {
		return "", err;
	}
	return res, nil;
}

func UpdateCommand(source string) (string, error) {
	req := fmt.Sprintf("UPDATE %v\n", source);

	res, err := sendCommand[string](req);
	if err != nil {
		return "", err;
	}
	return res, nil;
}

