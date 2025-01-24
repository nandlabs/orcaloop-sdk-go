package events

import (
	"oss.nandlabs.io/orcaloop-sdk/data"
	"oss.nandlabs.io/orcaloop-sdk/models"
)

// StepChangeEvent represents an event that indicates a change in the status of a step within a pipeline instance.
// It contains the following fields:
// - InstanceId: The unique identifier of the pipeline instance.
// - StepId: The unique identifier of the step within the pipeline instance.
// - Status: The current status of the step.
// - Data: Additional data related to the pipeline, represented by a Pipeline object.
type StepChangeEvent struct {
	InstanceId string         `json:"instance_id" yaml:"instance_id"`
	StepId     string         `json:"step_id" yaml:"step_id"`
	Status     models.Status  `json:"status" yaml:"status"`
	Data       *data.Pipeline `json:"data" yaml:"data"`
}
