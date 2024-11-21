package base

/*
import (
	"encoding/json"
	"fmt"
)

func JsonToMessage(mtype string, data []byte) (MessageI, error) {

	switch mtype {
		case "IPRequest": {
			var iprequest IPRequest
			err := json.Unmarshal(data, &iprequest)
			return iprequest, err
		}

		case "IPResponse": {
			var ipres IPResponse
			err := json.Unmarshal(data, &ipres)
			return ipres, err
		}
	}

	return nil, fmt.Errorf("mtype %s not found", mtype)
}

func MessageToJson(data any) ([]byte, error) {
	return json.Marshal(data)
}
*/
