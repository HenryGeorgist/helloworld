package main

import (
	"encoding/json"
	"fmt"

	"github.com/usace/wat-go"
)

func main() {
	fmt.Println("starting the hello world plugin!")
	ws, err := wat.NewS3WatStore()
	if err != nil {
		wat.Log(wat.Message{
			Status:    wat.FAILED,
			Progress:  0,
			Level:     wat.ERROR,
			Message:   err.Error(),
			Sender:    "helloworldplugin",
			PayloadId: "unknown",
		})
	}

	payload, err := ws.GetPayload()
	if err != nil {
		wat.Log(wat.Message{
			Status:    wat.FAILED,
			Progress:  0,
			Level:     wat.ERROR,
			Message:   err.Error(),
			Sender:    "helloworldplugin",
			PayloadId: "unknown",
		})
		return
	}
	err = computePayload(payload)
	if err != nil {
		wat.Log(wat.Message{
			Status:    wat.FAILED,
			Progress:  0,
			Level:     wat.ERROR,
			Message:   err.Error(),
			Sender:    "helloworldplugin",
			PayloadId: "unknown",
		})
		return
	}
}
func computePayload(payload wat.Payload) error {
	//payload contains information about resources for running and where to put output
	data, err := json.Marshal(payload)
	if err != nil {
		wat.Log(wat.Message{
			Status:    wat.FAILED,
			Progress:  0,
			Level:     wat.ERROR,
			Message:   err.Error(),
			Sender:    "helloworldplugin",
			PayloadId: "unknown",
		})
		return err
	}
	wat.Log(wat.Message{
		Status:    wat.SUCCEEDED,
		Progress:  100,
		Level:     wat.INFO,
		Message:   string(data),
		Sender:    "helloworldplugin",
		PayloadId: "unknown",
	})
	return nil
}
