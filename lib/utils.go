package lib

// https://stackoverflow.com/questions/26545883/how-to-do-one-liner-if-else-statement

func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}
