package api

import (
	"oss.nandlabs.io/golly/lifecycle"
	"oss.nandlabs.io/golly/rest"
	"oss.nandlabs.io/orcaloop-sdk/config"
	v1 "oss.nandlabs.io/orcaloop-sdk/service/api/v1"
)

func PrepareServer(serviceLifecycleManager lifecycle.ComponentManager, c *config.ActionSvcConfig) {
	var srv rest.Server
	var err error
	if c.Listener == nil {
		panic("Listener configuration is required")
	}
	options := &rest.SrvOptions{
		PathPrefix:     "/api/",
		ListenHost:     c.Listener.ListenHost,
		ListenPort:     c.Listener.ListenPort,
		ReadTimeout:    c.Listener.ReadTimeout,
		WriteTimeout:   c.Listener.WriteTimeout,
		EnableTLS:      c.Listener.EnableTLS,
		PrivateKeyPath: c.Listener.PrivateKeyPath,
		CertPath:       c.Listener.CertPath,
	}

	srv, err = rest.NewServer(options)
	if err != nil {
		panic(err)
	}

	// Register all Handlers
	srv.Post("v1/actions/:actionId", v1.ExecuteAction)
	//register the server with the lifecycle manager
	serviceLifecycleManager.Register(srv)
}
