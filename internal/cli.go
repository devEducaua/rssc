package internal

import "fmt"

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
		}

		resp, err := GetCommand(source, limit);
		if err != nil {
			return err;
		}
		for _,v := range resp {
			fmt.Println(v.Title);
		}

	default:
		return fmt.Errorf("command: `%v` not exists\n", argv[0]);
	}

	return nil;
}

type Flag struct {
	Enabled bool
	Value string
}

func parseFlags(argv []string) (map[string]Flag, error) {
	flagMap := make(map[string]Flag);

	for i,v := range argv {
		switch v {
		case "-i", "--id":
			var flag Flag;

			if len(argv) <= i+1 {
				return nil, fmt.Errorf("failed to parse flags: `--id` need an argument.");
			}

			flag.Value = argv[i+1];
			flag.Enabled = true;
			flagMap["id"] = flag;

		case "-l", "--limit":
			var flag Flag;

			if len(argv) <= i+1 {
				return nil, fmt.Errorf("failed to parse flags: `--limit` need an argument.");
			}

			flag.Value = argv[i+1];
			flag.Enabled = true;
			flagMap["limit"] = flag;

		case "-v", "--verbose":
			flag := Flag{ Enabled: true, Value: "" };
			flagMap["verbose"] = flag;

		case "-u", "--unread":
			flag := Flag{ Enabled: true, Value: "" };
			flagMap["unread"] = flag;

		case "-r", "--read":
			flag := Flag{ Enabled: true, Value: "" };
			flagMap["read"] = flag;

		case "-a", "--all":
			flag := Flag{ Enabled: true, Value: "" };
			flagMap["all"] = flag;

		case "-f", "--feed":
			var flag Flag;

			if len(argv) <= i+1 {
				return nil, fmt.Errorf("failed to parse flags: `--feed` need an argument.");
			}

			flag.Value = argv[i+1];
			flag.Enabled = true;
			flagMap["feed"] = flag;
		}
	}

	return flagMap, nil;
}
