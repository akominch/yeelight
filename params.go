package yeelight

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

type YeelightParams struct {
	Model     string   `json:"model"`
	Support   []string `json:"support"`
	Power     string   `json:"power"`
	Bright    int      `json:"bright"`
	ColorMode int      `json:"color_mode"`
	RGB       int      `json:"rgb"`
	Hue       int      `json:"hue"`
	Sat       int      `json:"sat"`
	Name      string   `json:"name"`
}

func parseAnswer(msg string) *YeelightParams {
	dict := make(map[string]interface{})

	arr := strings.Split(msg, crlf)
	for _, line := range arr {
		if strings.Contains(line, ":") {
			//fmt.Println(line)
			lineArr := strings.Split(line, ": ")
			key := lineArr[0]
			value := lineArr[1]

			switch key {
			case "support":
				valueArr := strings.Split(value, " ")
				dict[key] = valueArr
			case "fw_ver", "bright", "color_mode", "rgb", "hue", "sat":
				intValue, err := strconv.Atoi(value)
				if err != nil {
					log.Println("Error convert to int", key)
				}
				dict[key] = intValue
			default:
				dict[key] = value
			}
		}
	}
	j, err := json.Marshal(dict)
	if err != nil {
		log.Println("Error convert params dict to JSON")
	}

	params := new(YeelightParams)
	err = json.Unmarshal(j, &params)
	if err != nil {
		log.Println("Error convert JSON to param struct")
	}

	return params
}
