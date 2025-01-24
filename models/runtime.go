package models

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
