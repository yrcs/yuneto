package util

// IfThenElse ternary operator
func IfThenElse(condition bool, tV, fV any) any {
	if condition {
		return tV
	} else {
		return fV
	}
}
