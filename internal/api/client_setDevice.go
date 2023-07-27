package api

import (
	"fmt"
	"net/http"
)

func (h *ClientAPI) ClientSetDevice(ports string, state string, name string) error {
	r, err := http.NewRequest("POST", "http://localhost:8000/api/devices/"+name+"/"+ports+"&"+state, nil)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}

	client := &http.Client{}
	_, err = client.Do(r)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}

	return nil
}
