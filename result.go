package commandr

// BaseResult provides a basic implementation of the Result interface.
// It can be embedded in custom result types to extend functionality.
type BaseResult struct {
	Output string
}

func NewResult(output string) *BaseResult {
	return &BaseResult{
		Output: output,
	}
}

// GetOutput returns the command output.
func (r *BaseResult) GetOutput() string { return r.Output }

// SetOutput sets the command output.
func (r *BaseResult) SetOutput(out string) { r.Output = out }
