package internal

import (
	"fmt"
)

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

		case "-t", "--text":
			var flag Flag;

			if len(argv) <= i+1 {
				return nil, fmt.Errorf("failed to parse flags: `--feed` need an argument.");
			}

			flag.Value = argv[i+1];
			flag.Enabled = true;
			flagMap["text"] = flag;
		}
	}

	return flagMap, nil;
}
