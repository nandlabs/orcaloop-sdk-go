package builder

import "oss.nandlabs.io/orcaloop-sdk/models"

// WorkflowBuilder is a builder for constructing Workflow objects.
// It embeds StepsBuilder to provide step-building functionality and
// contains a reference to a Workflow model.
type WorkflowBuilder struct {
	*StepsBuilder
	workflow *models.Workflow
}

// NewWorkflowBuilder creates a new instance of WorkflowBuilder with initialized
// StepsBuilder and an empty Workflow model.
// Returns a pointer to the newly created WorkflowBuilder.
func NewWorkflowBuilder() *WorkflowBuilder {
	return &WorkflowBuilder{
		StepsBuilder: NewStepsBuilder(),
		workflow:     &models.Workflow{},
	}
}

// Id sets the ID of the workflow and returns the updated WorkflowBuilder instance.
//
// Parameters:
//
//	id - a string representing the unique identifier for the workflow.
//
// Returns:
//
//	*WorkflowBuilder - the updated WorkflowBuilder instance with the new ID set.
func (b *WorkflowBuilder) Id(id string) *WorkflowBuilder {
	b.workflow.Id = id
	return b
}

// Name sets the name of the workflow.
// It takes a string parameter 'name' which represents the name to be set.
// It returns a pointer to the WorkflowBuilder to allow for method chaining.
func (b *WorkflowBuilder) Name(name string) *WorkflowBuilder {
	b.workflow.Name = name
	return b
}

// Version sets the version of the workflow.
// It takes an integer representing the version and returns the updated WorkflowBuilder instance.
//
// Parameters:
//   - version: An integer representing the version of the workflow.
//
// Returns:
//   - *WorkflowBuilder: The updated WorkflowBuilder instance.
func (b *WorkflowBuilder) Version(version int) *WorkflowBuilder {
	b.workflow.Version = version
	return b
}

// Description sets the description of the workflow.
// Description sets the description of the workflow and returns the updated WorkflowBuilder.
//
// Parameters:
//
//	description - A string representing the description of the workflow.
//
// Returns:
//
//	*WorkflowBuilder - The updated WorkflowBuilder instance.
func (b *WorkflowBuilder) Description(description string) *WorkflowBuilder {
	b.workflow.Description = description
	return b
}

// Build returns the built workflow.
// Build finalizes the WorkflowBuilder by assigning the accumulated steps to the workflow
// and returns the constructed Workflow object.
func (b *WorkflowBuilder) Build() *models.Workflow {
	b.workflow.Steps = b.steps
	return b.workflow
}
