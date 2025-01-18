package models

import "oss.nandlabs.io/orcaloop-sdk/data"

const (
	// StatusUnknown represents an unknown status of the workflow
	StatusUnknown Status = iota
	// StatusPending represents a pending status of the workflow
	StatusPending
	// StatusRunning represents a running status of the workflow
	StatusRunning
	// StatusCompleted represents a completed status of the workflow
	StatusCompleted
	// StatusFailed represents a failed status of the workflow
	StatusFailed
	// StatusSkipped represents a skipped status of the workflow
	StatusSkipped

	// String representations of the status constants

	// StatusPendingStr is the string representation of the StatusPending constant
	StatusPendingStr = "Pending"
	// StatusRunningStr is the string representation of the StatusRunning constant
	StatusRunningStr = "Running"
	// StatusCompletedStr is the string representation of the StatusCompleted constant
	StatusCompletedStr = "Completed"
	// StatusFailedStr is the string representation of the StatusFailed constant
	StatusFailedStr = "Failed"
	// StatusSkippedStr is the string representation of the StatusSkipped constant
	StatusSkippedStr = "Skipped"
	// StatusUnkonwnStr is the string representation of the StatusUnknown constant
	StatusUnkonwnStr = "Unknown"
)

// Status represents the execution status of the workflow
type Status int

// String returns the string representation of the Status.
// It maps each Status value to its corresponding string constant.
// If the Status value is not recognized, it returns StatusUnkonwnStr.
// Implements Stringer interface.
func (s Status) String() string {
	switch s {
	case StatusPending:
		return StatusPendingStr
	case StatusRunning:
		return StatusRunningStr
	case StatusCompleted:
		return StatusCompletedStr
	case StatusFailed:
		return StatusFailedStr
	case StatusSkipped:
		return StatusSkippedStr
	default:
		return StatusUnkonwnStr
	}
}

// WorkflowState represents the state of a workflow at a given point in time.
// It includes information such as the workflow's ID, version, status, the workflow itself,
// the states of individual steps within the workflow, and any error that may have occurred.
//
// Fields:
// - Id: The unique identifier of the workflow.
// - Version: The version number of the workflow.
// - Status: The current status of the workflow.
// - Workflow: The workflow object itself.
// - StepStates: A map of step identifiers to their corresponding StepState objects.
// - Error: Any error message associated with the workflow.
type WorkflowState struct {
	Id         string                `json:"id" yaml:"id"`
	Version    int                   `json:"version" yaml:"version"`
	Status     Status                `json:"status" yaml:"status"`
	Workflow   *Workflow             `json:"workflow" yaml:"workflow"`
	StepStates map[string]*StepState `json:"step_states" yaml:"step_states"`
	Error      string                `json:"error" yaml:"error"`
}

// StepState represents the state of a step in a pipeline execution.
// It includes information about the instance, step identifiers, parent-child relationships,
// status, and input/output data of the step.
//
// Fields:
// - InstanceId: The unique identifier of the instance.
// - StepId: The unique identifier of the step.
// - ParentStep: The identifier of the parent step, if any.
// - ChildCount: The number of child steps associated with this step.
// - Status: The current status of the step.
// - Input: The input data for the step, represented as a Pipeline object.
// - Output: The output data from the step, represented as a Pipeline object.
type StepState struct {
	InstanceId string         `json:"instance_id" yaml:"instance_id"`
	StepId     string         `json:"step_id" yaml:"step_id"`
	ParentStep string         `json:"parent_step" yaml:"parent_step"`
	ChildCount int            `json:"child_count" yaml:"child_count"`
	Status     Status         `json:"status" yaml:"status"`
	Input      *data.Pipeline `json:"input" yaml:"input"`
	Output     *data.Pipeline `json:"output" yaml:"output"`
}

// StepChangeEvent represents an event that indicates a change in the status of a step within a pipeline instance.
// It contains the following fields:
// - InstanceId: The unique identifier of the pipeline instance.
// - StepId: The unique identifier of the step within the pipeline instance.
// - Status: The current status of the step.
// - Data: Additional data related to the pipeline, represented by a Pipeline object.
type StepChangeEvent struct {
	InstanceId string         `json:"instance_id" yaml:"instance_id"`
	StepId     string         `json:"step_id" yaml:"step_id"`
	Status     Status         `json:"status" yaml:"status"`
	Data       *data.Pipeline `json:"data" yaml:"data"`
}
