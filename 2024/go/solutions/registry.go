package solutions

// Registry for day functions
var days = map[string]func(){}

// RegisterDay allows each day to register its function
func RegisterDay(name string, fn func()) {
	days[name] = fn
}

// RunAllDays executes all registered days
func RunAllDays() {
	for _, fn := range days {
		fn()
	}
}
