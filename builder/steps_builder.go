package builder

import "oss.nandlabs.io/orcaloop-sdk/models"

// StepsBuilder is a builder for constructing a sequence of steps.
// It maintains a list of steps and a tracker to keep track of steps by their names.
type StepsBuilder struct {
	steps        []*models.Step
	stepsTracker map[string]*models.Step
}

// NewStepsBuilder creates a new instance of StepsBuilder with initialized steps slice and stepsTracker map.
// It returns a pointer to the newly created StepsBuilder.
func NewStepsBuilder() *StepsBuilder {
	return &StepsBuilder{
		steps:        []*models.Step{},
		stepsTracker: make(map[string]*models.Step),
	}
}

// trackSteps adds the provided steps to the stepsTracker map using their Id as the key.
// It accepts a variadic number of pointers to models.Step.
//
// Parameters:
//
//	steps - A variadic number of pointers to models.Step to be tracked.
func (b *StepsBuilder) trackSteps(steps ...*models.Step) {
	for _, s := range steps {
		b.stepsTracker[s.Id] = s
	}
}

// AddStep adds a new step to the StepsBuilder and tracks the step.
// It appends the provided step to the steps slice and calls the trackSteps method.
//
// Parameters:
//   - step: A pointer to the Step model to be added.
//
// Returns:
//   - A pointer to the updated StepsBuilder instance.
func (b *StepsBuilder) AddStep(step *models.Step) *StepsBuilder {
	b.steps = append(b.steps, step)
	b.trackSteps(step)
	return b
}

// AddActionStep adds an action step to the StepsBuilder with the given id, name, parameters, and output.
// It creates a new Step of type StepTypeAction and sets its Action field with the provided details.
// The step is then tracked and added to the StepsBuilder.
//
// Parameters:
//   - id: The unique identifier for the step.
//   - name: The name of the action step.
//   - parameters: A slice of Parameter pointers representing the parameters for the action.
//   - output: A slice of strings representing the output of the action.
//
// Returns:
//   - *StepsBuilder: The updated StepsBuilder instance with the new action step added.
func (b *StepsBuilder) AddActionStep(id, name string, parameters []*models.Parameter, output []string) *StepsBuilder {
	step := &models.Step{
		Id:   id,
		Type: models.StepTypeAction,
		Action: &models.StepAction{
			Id:         id,
			Name:       name,
			Parameters: parameters,
			Output:     output,
		},
	}
	b.trackSteps(step)
	return b.AddStep(step)
}

// AddParallelSteps adds multiple steps to be executed in parallel to the StepsBuilder.
// It accepts a variadic number of pointers to models.Step and returns the updated StepsBuilder.
//
// Parameters:
//
//	steps - A variadic number of pointers to models.Step that will be executed in parallel.
//
// Returns:
//
//	*StepsBuilder - The updated StepsBuilder with the added parallel steps.
func (b *StepsBuilder) AddFor(loopVar, indexVar, itemsVar string, items []interface{}, steps ...*models.Step) *StepsBuilder {
	step := &models.Step{
		Type: models.StepTypeForLoop,
		For: &models.For{
			Loopvar:  loopVar,
			IndexVar: indexVar,
			ItemsVar: itemsVar,
			ItemsArr: items,
			Steps:    steps,
		},
	}
	b.trackSteps(step)
	return b.AddStep(step)
}

// AddIf adds a conditional step to the StepsBuilder. The step will only be executed
// if the specified condition is met.
//
// Parameters:
//   - condition: A string representing the condition to be evaluated.
//   - steps: A variadic parameter representing the steps to be executed if the condition is true.
//
// Returns:
//   - *StepsBuilder: The updated StepsBuilder instance.
func (b *StepsBuilder) AddIf(condition string, steps ...*models.Step) *StepsBuilder {
	step := &models.Step{
		Type: models.StepTypeIf,
		If: &models.If{
			Condition: condition,
			Steps:     steps,
		},
	}
	b.trackSteps(step)
	return b.AddStep(step)
}

// AddSwitch adds a switch step to the StepsBuilder. A switch step allows for
// conditional branching based on the value of a variable. The variable is
// evaluated, and the corresponding case is executed.
//
// Parameters:
//   - variable: The variable to be evaluated in the switch step.
//   - cases: A slice of Case objects representing the possible branches.
//
// Returns:
//   - *StepsBuilder: The updated StepsBuilder instance with the new switch step added.
func (b *StepsBuilder) AddSwitch(variable string, cases []*models.Case) *StepsBuilder {
	step := &models.Step{
		Type: models.StepTypeSwitch,
		Switch: &models.Switch{
			Variable: variable,
			Cases:    cases,
		},
	}
	b.trackSteps(step)
	return b.AddStep(step)
}

// AddStepToFor adds a new step to an existing For loop.
// AddStepsToFor adds new steps to the "For" field of an existing step identified by forStepId.
// If the step with forStepId exists and has a non-nil "For" field, the new steps are appended to the existing steps.
// The method also tracks the new steps using the trackSteps method.
//
// Parameters:
//   - forStepId: The ID of the step to which new steps will be added.
//   - newSteps: A variadic parameter representing the new steps to be added.
//
// Returns:
//   - *StepsBuilder: The updated StepsBuilder instance.
func (b *StepsBuilder) AddStepToSwitchCase(switchCase string, caseVal string, newSteps ...*models.Step) *StepsBuilder {
	if switchStep, exists := b.stepsTracker[switchCase]; exists && switchStep.Switch != nil {
		for _, c := range switchStep.Switch.Cases {
			if c.Value == caseVal {
				c.Steps = append(c.Steps, newSteps...)
				b.trackSteps(newSteps...)
			}
		}
	}
	return b
}

// AddStepToSwitchDefault adds new steps to the default case of a switch statement
// identified by the given switchCase string. If the switch case exists and has a
// default case, the new steps are appended to the existing steps of the default case.
// The steps are also tracked by the StepsBuilder.
//
// Parameters:
//   - switchCase: A string representing the identifier of the switch case.
//   - newSteps: A variadic parameter of pointers to models.Step representing the new steps to be added.
//
// Returns:
//   - *StepsBuilder: The StepsBuilder instance to allow for method chaining.
func (b *StepsBuilder) AddStepToSwitchDefault(switchCase string, newSteps ...*models.Step) *StepsBuilder {
	if switchStep, exists := b.stepsTracker[switchCase]; exists && switchStep.Switch != nil {
		for _, c := range switchStep.Switch.Cases {
			if c.Default {
				c.Steps = append(c.Steps, newSteps...)
				b.trackSteps(newSteps...)
			}
		}
	}
	return b
}

// Build constructs and returns a slice of pointers to models.Step
// that have been accumulated in the StepsBuilder.
//
// Returns:
//
//	[]*models.Step: A slice of pointers to models.Step.
func (b *StepsBuilder) Build() []*models.Step {
	return b.steps
}
