package internal

import (
	"fmt"
	"strings"
	"encoding/json"
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

func GetCommand(source string, limit int) ([]GetResponse, error) {
	conn, err := Connect();
	if err != nil {
		return nil, err;
	}

	var req strings.Builder;

	req.WriteString("GET ");
	req.WriteString(source);

	if limit != 0 {
		fmt.Fprintf(&req, " %v", limit)
	}
	req.WriteString("\n");

	err = Send(conn, req.String());
	if err != nil {
		return nil, err;
	}

	resp, err := Recv(conn);
	if err != nil {
		return nil, err;
	}

	if resp.Status != "yes" {
		return nil, fmt.Errorf("request failed: %v", resp.Response);
	}

	var res []GetResponse;
	err = json.Unmarshal(resp.Response, &res);
	if err != nil {
		return nil, err;
	}

	return res, nil;
}

