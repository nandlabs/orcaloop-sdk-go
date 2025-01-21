package service

import (
	"oss.nandlabs.io/golly/lifecycle"
	"oss.nandlabs.io/orcaloop-sdk/config"
	"oss.nandlabs.io/orcaloop-sdk/service/api"
)

var serviceLifecycleManager = lifecycle.NewSimpleComponentManager()

func Start(c *config.ActionSvcConfig) {
	//prepare the server
	api.PrepareServer(serviceLifecycleManager, c)
	//start the server
	serviceLifecycleManager.StartAndWait()
}

func Stop() {
	serviceLifecycleManager.StopAll()
}
