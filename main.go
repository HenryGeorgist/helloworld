package main

import (
	"flag"
	"fmt"

	"github.com/usace/wat-go-sdk/plugin"
)

func main() {
	fmt.Println("starting the hello world plugin!")
	var payloadPath string
	flag.StringVar(&payloadPath, "payload", "pathtopayload.yml", "please specify an input file using `-payload pathtopayload.yml`")
	flag.Parse()
	if payloadPath == "" {
		plugin.Log(plugin.Message{
			Status:    plugin.FAILED,
			Progress:  0,
			Level:     plugin.ERROR,
			Message:   "given a blank path...\n\tplease specify an input file using `-payload pathtopayload.yml`",
			Sender:    "helloworldplugin",
			PayloadId: "unknown payloadid because the plugin package could not be properly initalized",
		})
		return
	}
	err := plugin.InitConfigFromEnv()
	if err != nil {
		logError(err, plugin.ModelPayload{Id: "unknownpayloadid"})
		return
	}
	payload, err := plugin.LoadPayload(payloadPath)
	if err != nil {
		logError(err, plugin.ModelPayload{Id: "unknownpayloadid"})
		return
	}
	err = computePayload(payload)
	if err != nil {
		logError(err, payload)
		return
	}
}
func computePayload(payload plugin.ModelPayload) error {
	//payload contains information about resources for running and where to put output
	plugin.Log(plugin.Message{
		Status:    plugin.SUCCEEDED,
		Progress:  100,
		Level:     plugin.INFO,
		Message:   "Hello World!!",
		Sender:    "helloworldplugin",
		PayloadId: payload.Id,
	})
	return nil
}
func logError(err error, payload plugin.ModelPayload) {
	plugin.Log(plugin.Message{
		Status:    plugin.FAILED,
		Progress:  0,
		Level:     plugin.ERROR,
		Message:   err.Error(),
		Sender:    "helloworldplugin",
		PayloadId: payload.Id,
	})
}
