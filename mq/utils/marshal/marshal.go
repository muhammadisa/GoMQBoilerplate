package marshal

import (
	"encoding/json"

	"github.com/muhammadisa/go-mq-boilerplate/mq/utils/errhandler"
)

// U unmarshaling data
func U(body []byte, data interface{}) error {
	err := json.Unmarshal(body, &data)
	return errhandler.HandleErrorThenReturn(err)
}

// M marshalling data
func M(data interface{}) []byte {
	byteArray, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return byteArray
}
