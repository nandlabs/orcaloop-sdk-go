package service

import (
	"oss.nandlabs.io/golly/lifecycle"
	"oss.nandlabs.io/golly/rest/server"
	"oss.nandlabs.io/orcaloop-sdk/config"
	v1 "oss.nandlabs.io/orcaloop-sdk/service/api/v1"
)

var serviceLifecycleManager = lifecycle.NewSimpleComponentManager()

func Start(c *config.ActionSvcConfig) {
	var srv server.Server
	var err error
	if c.Listener == nil {
		panic("Listener configuration is required")
	}
	options := server.NewOptions()
	options.PathPrefix = "/api/"
	options.ListenHost = c.Listener.ListenHost
	options.ListenPort = c.Listener.ListenPort
	options.ReadTimeout = c.Listener.ReadTimeout
	options.WriteTimeout = c.Listener.WriteTimeout
	options.EnableTLS = c.Listener.EnableTLS
	options.PrivateKeyPath = c.Listener.PrivateKeyPath
	options.CertPath = c.Listener.CertPath

	srv, err = server.New(options)
	if err != nil {
		panic(err)
	}

	// Register all Handlers
	srv.Post("/actions/:actionId", v1.ExecuteAction)
	//register the server with the lifecycle manager
	serviceLifecycleManager.Register(srv)
	//start the server
	serviceLifecycleManager.StartAndWait()
}

func Stop() {
	serviceLifecycleManager.StopAll()
}
