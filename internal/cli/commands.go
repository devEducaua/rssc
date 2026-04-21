package cli

import (
	"fmt"
	"rssc/internal/requests"
)

func ParseArgs(argv []string) error {
	if len(argv) < 2 {
		return nil; // TODO: print usage or go to man page
	}

	flagMap, err := parseFlags(argv[1:]);
	if err != nil {
		return err;
	}

	switch argv[0] {
	case "get":

		var limit int;
		if !flagMap["limit"].Enabled {
			limit = 0;
		}

		var source string;

		if flagMap["all"].Enabled {
			source = "ALL";
		} else if flagMap["read"].Enabled {
			source = "READ";

		} else if flagMap["unread"].Enabled{
			source = "UNREAD";

		} else if flagMap["feed"].Enabled {
			source = flagMap["feed"].Value;

		}  else if flagMap["id"].Enabled {
			source = flagMap["id"].Value;
		}

		resp, err := requests.GetCommand(source, limit);
		if err != nil {
			return err;
		}
		for _,v := range resp {
			fmt.Println(v.Title);
		}

	case "find":
		var limit int;
		if !flagMap["limit"].Enabled {
			limit = 0;
		}

		var text string;

		if flagMap["text"].Enabled {
			text = flagMap["text"].Value;
		}

		resp, err := requests.FindCommand(text, limit);
		if err != nil {
			return err;
		}
		for _,v := range resp {
			fmt.Println(v);
		}

	case "read":
		var id string;
		if flagMap["id"].Enabled {
			id = flagMap["id"].Value;
		}

		resp, err := requests.ReadCommand(id);
		if err != nil {
			return err;
		}
		fmt.Println(resp);

	case "unread":
		var id string;
		if flagMap["id"].Enabled {
			id = flagMap["id"].Value;
		}

		resp, err := requests.UnreadCommand(id);
		if err != nil {
			return err;
		}
		fmt.Println(resp);

	case "delete":
		var id string;
		if flagMap["id"].Enabled {
			id = flagMap["id"].Value;
		}

		resp, err := requests.DeleteCommand(id);
		if err != nil {
			return err;
		}
		fmt.Println(resp);

	case "open":
		var id string;
		if flagMap["id"].Enabled {
			id = flagMap["id"].Value;
		}

		resp, err := requests.OpenCommand(id);
		if err != nil {
			return err;
		}
		fmt.Println(resp);

	case "update":
		var source string;
		if flagMap["source"].Enabled {
			source = flagMap["source"].Value;
		}

		resp, err := requests.UpdateCommand(source);
		if err != nil {
			return err;
		}
		fmt.Println(resp);

	default:
		return fmt.Errorf("command: `%v` not exists\n", argv[0]);
	}

	return nil;
}

