package messaging

import (
	"net/url"

	"oss.nandlabs.io/golly/l3"
	"oss.nandlabs.io/golly/lifecycle"
	"oss.nandlabs.io/golly/messaging"
	"oss.nandlabs.io/orcaloop-sdk/data"
	"oss.nandlabs.io/orcaloop-sdk/handlers"
	"oss.nandlabs.io/orcaloop-sdk/service"
)

var logger = l3.Get()

type MsgListener struct {
	*lifecycle.SimpleComponent
	// url is the url of the messaging endpoint
	url *url.URL
}

// NewMsgListener creates a new MsgListener
func NewMsgListener(url *url.URL, id string, client *service.OrcaloopClient) *MsgListener {
	return &MsgListener{
		SimpleComponent: &lifecycle.SimpleComponent{
			CompId: id + "-msg-listener",
			StartFunc: func() (err error) {
				manager := messaging.GetManager()

				go func() {
					//Create a named listener
					options := messaging.NewOptionsBuilder().AddNamedListener(id).Build()
					// Add the listener
					err = manager.AddListener(url, func(msg messaging.Message) {
						var actionId string
						var actionHandler handlers.ActionHandler

						body := make(map[string]any)
						err = msg.ReadJSON(&body)
						if err != nil {
							logger.ErrorF("Failed to decode message body: %v", err)
							return
						}
						pipeline := data.NewPipelineFrom(body)

						actionId = pipeline.GetActionId()
						actionHandler = handlers.ActionRegistry.Get(actionId)
						if actionHandler == nil {
							logger.ErrorF("Action not found: %v", err)
							return
						}

						err = actionHandler.Handle(pipeline)
						if err != nil {
							logger.ErrorF("Failed to handle action: %v", err)
							return
						}
					}, options...)
				}()
				return
			},
			StopFunc: func() (err error) {
				//TODO: Implement stop function
				// Remove the listener
				return
			},
		},
		url: url,
	}
}
