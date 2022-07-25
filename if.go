package goutils

// If Return trueVal if the condition holds, otherwise falseVal
func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}
