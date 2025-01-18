package models

import (
	"oss.nandlabs.io/golly/clients"
	"oss.nandlabs.io/golly/codec"
	"oss.nandlabs.io/golly/managers"
	"oss.nandlabs.io/orcaloop-sdk/data"
)

// Action represents an interface for defining actions in the system.
// Each action is expected to have a unique identifier, name, description,
// input and output schemas, and an execution method. Additionally, actions
// can specify whether they are asynchronous and provide a specification
// and provider information.
type Action interface {
	// Id returns the id of the action. This is expected to be unique.
	Id() string
	// Name returns the name of the action
	Name() string
	// Description returns the description of the action
	Description() string
	// Inputs returns the inputs of the action
	Inputs() []*Schema
	// Outputs returns the outputs of the action
	Outputs() []*Schema
	// Execute function executes the action
	Execute(*data.Pipeline) error
	// Spec returns the spec of the action
	Spec() *ActionSpec
	// IsAsync returns true if the action is asynchronous
	IsAsync() bool
	// Provider returns the provider of the action
	Provider() string
}

// ActionSpec represents the specification of an action.
// It contains metadata and configuration for the action.
//
// Fields:
//   - Id: The unique identifier of the action.
//   - Name: The name of the action.
//   - Description: A brief description of the action.
//   - Parameters: A list of schemas representing the parameters of the action.
//   - Returns: A list of schemas representing the return values of the action.
//   - Async: A flag indicating whether the action is asynchronous.
//   - Endpoint: The endpoint configuration for the action.
type ActionSpec struct {
	// Id is the id of the action
	Id string `json:"id" yaml:"id"`
	// Name is the name of the action
	Name string `json:"name" yaml:"name"`
	// Description is the description of the action
	Description string `json:"description" yaml:"description"`
	//Parameters is the parameters of the action
	Parameters []*Schema `json:"parameters" yaml:"parameters"`
	// Returns is the returns of the action
	Returns []*Schema `json:"returns" yaml:"returns"`
	// Async  is the async flag of the action
	Async bool `json:"async" yaml:"async"`
	// Endpoint is the endpoint of the action
	Endpoint *Endpoint `json:"endpoint" yaml:"endpoint"`
}

// Endpoint represents a network endpoint configuration.
//
// Fields:
//   - Type: Specifies the type of the Endpoint.
//   - Local: Represents the local endpoint configuration.
//   - Rest: Represents the REST endpoint configuration.
//   - Messaging: Represents the messaging endpoint configuration.
//   - Grpc: Represents the gRPC endpoint configuration.
//   - Qos: Represents the quality of service configuration.
type Endpoint struct {

	// type is the type of the Endpoint
	Type string `json:"type" yaml:"type"`
	//Local is the local endpoint
	Builtin *Builtin `json:"builtin" yaml:"builtin"`
	//Rest is the rest endpoint
	Rest *RestEndpoint `json:"rest" yaml:"rest"`
	//Messaging is the messaging endpoint
	Messaging *MessagingEndpoint `json:"messaging" yaml:"messaging"`
	//Qos is the quality of service
	Qos *Qos `json:"qos" yaml:"qos"`
}

// Qos represents the Quality of Service settings for a particular action.
// It includes the number of retries, timeout duration, and circuit breaker information.
type Qos struct {
	// Retries is the number of retries
	Retries int
	// Timeout is the timeout
	Timeout int
	// CircuitBreakerInfo is the circuit breaker info
	BreakerInfo *clients.BreakerInfo
}

type Builtin struct {
}

type RestEndpoint struct {
	// Url is the url of the rest endpoint
	// The format is http[s]://host:[port]/[pathname|$pathParam]*/?[name=value&]*
	// The $pathParam is the path param name from the pipeline
	Url string `json:"url" yaml:"url"`
	// AuthProvider
	AuthProvider string `json:"authProvider" yaml:"authProvider"`
}

// MessagingEndpoint represents the configuration for a messaging endpoint.
// It contains the URL, request headers, body, and body MIME type for the messaging provider.
// See https://pkg.go.dev/oss.nandlabs.io/golly/messaging#Provider for more information for the messaging provider.
type MessagingEndpoint struct {
	// Url is the URL of the messaging provider.
	// The format is <provider_scheme>://destination.
	// The provider_scheme is the scheme of the messaging provider.
	// The destination is the destination of the messaging provider.
	Url string `json:"url" yaml:"url"`
}

// ActionSpecs retrieves the specifications of all actions managed by the given action manager.
// It iterates over the items in the action manager, calls the Specification method on each action,
// and appends the resulting ActionSpec to a slice, which is then returned.
//
// Parameters:
//   - actionManager: An instance of ItemManager that manages Action items.
//
// Returns:
//   - A slice of pointers to ActionSpec, each representing the specification of an action.
func ActionSpecs(actionManager managers.ItemManager[Action]) []*ActionSpec {
	var actions []*ActionSpec

	for _, action := range actionManager.Items() {
		actions = append(actions, action.Spec())
	}
	return actions
}

// DescribeActions encodes the specifications of actions managed by the given action manager
// into a string format specified by the provided format string.
//
// Parameters:
//   - format: A string specifying the desired encoding format.
//   - actionManager: An ItemManager instance managing Action items.
//
// Returns:
//   - val: A string containing the encoded action specifications.
//   - err: An error if the encoding process fails, otherwise nil.
func DescribeActions(format string, actionManager managers.ItemManager[Action]) (val string, err error) {
	var c codec.Codec

	c, err = codec.GetDefault(format)
	if err != nil {
		return
	}
	val, err = c.EncodeToString(ActionSpecs(actionManager))
	return
}
