// Package models provides the data structures and constants used to define
// workflows in the orcaloop-sdk. This includes the definition of workflows,
// steps, parameters, and various control flow constructs such as parallel
// execution, loops, conditionals, and switches.
package models

const (
	// StepTypeAction represents an action step in the workflow.
	StepTypeAction = "Action"
	// StepTypeParallel represents a parallel execution step in the workflow.
	StepTypeParallel = "Parallel"
	// StepTypeIf represents a conditional step in the workflow.
	StepTypeIf = "If"
	// StepTypeSwitch represents a switch-case step in the workflow.
	StepTypeSwitch = "Switch"
	// StepTypeForLoop represents a loop step in the workflow.
	StepTypeForLoop = "ForLoop"
	// ModeSync represents synchronous execution mode.
	ModeSync = "sync"
	// ModeAsync represents asynchronous execution mode.
	ModeAsync = "async"
)

// Workflow represents a workflow definition.
// Fields:
// - Id: Unique identifier for the workflow.
// - Name: Name of the workflow.
// - Version: Version number of the workflow.
// - Description: Description of the workflow.
// - Steps: List of steps in the workflow.
type Workflow struct {
	Id          string  `yaml:"id" json:"id"`
	Name        string  `yaml:"name" json:"name"`
	Version     int     `yaml:"version" json:"version"`
	Description string  `yaml:"description" json:"description"`
	Steps       []*Step `yaml:"steps" json:"steps"`
}

// Parameter represents a parameter in the workflow.
// Fields:
// - Name: Name of the parameter.
// - Value: Value of the parameter.
// - Var: Variable name of the parameter.
type Parameter struct {
	Name  string `yaml:"name" json:"name"`
	Value any    `yaml:"value" json:"value"`
	Var   string `yaml:"var" json:"var"`
}

// Result represents the result of an action execution.
// Fields:
// - OutputVar:  variable name of the tool output.
// - PipelineVar: Variable Name in pipeline to store the result.
type Result struct {
	OutputVar   string `yaml:"output_var" json:"output_var"`
	PipelineVar string `yaml:"pipeline_var" json:"pipeline_var"`
}

// StepAction represents a substep in the workflow.
// Fields:
// - Id: Unique identifier for the action.
// - Name: Name of the action.
// - Parameters: List of parameters for the action.
// - Output: List of output names for the action.
type StepAction struct {
	Id         string       `yaml:"id" json:"id"`
	Name       string       `yaml:"name" json:"name"`
	Parameters []*Parameter `yaml:"parameters" json:"parameters"`
	Results    []*Result    `yaml:"results" json:"results"`
}

// Step represents a step in the workflow.
// Fields:
// - Id: Unique identifier for the step.
// - Skip: Flag indicating if the step should be skipped.
// - Type: Type of the step (e.g., Action, Parallel, If, Switch, ForLoop).
// - Parallel: Parallel execution step.
// - For: Loop step.
// - If: Conditional step.
// - Switch: Switch-case step.
// - Action: Action step.
type Step struct {
	Id       string      `yaml:"id" json:"id"`
	Skip     bool        `yaml:"skip" json:"skip"`
	Type     string      `yaml:"type" json:"type"`
	Parallel *Parallel   `yaml:"parallel,omitempty" json:"parallel,omitempty"`
	For      *For        `yaml:"for,omitempty" json:"for,omitempty"`
	If       *If         `yaml:"if,omitempty" json:"if,omitempty"`
	Switch   *Switch     `yaml:"switch,omitempty" json:"switch,omitempty"`
	Action   *StepAction `yaml:"action,omitempty" json:"action,omitempty"`
}

// Parallel represents a parallel execution step in the workflow.
// Fields:
// - Steps: List of steps to be executed in parallel.
type Parallel struct {
	Steps []*Step `yaml:"steps" json:"steps"`
}

// For represents a loop step in the workflow.
// Fields:
// - Loopvar: Loop variable name.
// - IndexVar: Index variable name.
// - ItemsVar: Items variable name.
// - ItemsArr: Array of items to loop over.
// - Steps: List of steps to be executed in the loop.
type For struct {
	Loopvar  string  `yaml:"loop_var" json:"loop_var"`
	IndexVar string  `yaml:"index_var" json:"index_var"`
	ItemsVar string  `yaml:"items_var" json:"items_var"`
	ItemsArr []any   `yaml:"items" json:"items"`
	Steps    []*Step `yaml:"steps" json:"steps"`
}

// If represents a conditional step in the workflow.
// Fields:
// - Condition: Condition to evaluate.
// - Steps: List of steps to be executed if the condition is true.
// - ElseIfs: List of else-if branches.
// - Else: Else branch.
type If struct {
	Condition string    `yaml:"condition" json:"condition"`
	Steps     []*Step   `yaml:"steps" json:"steps"`
	ElseIfs   []*ElseIf `yaml:"else_ifs" json:"else_ifs"`
	Else      *Else     `yaml:"else" json:"else"`
}

// ElseIf represents an else-if branch in a conditional step.
// Fields:
// - Condition: Condition to evaluate.
// - Steps: List of steps to be executed if the condition is true.
type ElseIf struct {
	Condition string  `yaml:"condition" json:"condition"`
	Steps     []*Step `yaml:"steps" json:"steps"`
}

// Else represents an else branch in a conditional step.
// Fields:
// - Steps: List of steps to be executed if none of the conditions are true.
type Else struct {
	Steps []*Step `yaml:"steps" json:"steps"`
}

// Switch represents a switch-case step in the workflow.
// Fields:
// - Variable: Variable to switch on.
// - Cases: List of cases to match against.
type Switch struct {
	Variable string  `yaml:"variable" json:"variable"`
	Cases    []*Case `yaml:"cases" json:"cases"`
}

// Case represents a conditional branch in a workflow.
// Fields:
// - Value: Value to match against.
// - Default: Flag indicating if this is the default case.
// - Steps: List of steps to execute if the case matches.
type Case struct {
	Value   any     `yaml:"value" json:"value"`
	Default bool    `yaml:"default" json:"default"`
	Steps   []*Step `yaml:"steps" json:"steps"`
}
