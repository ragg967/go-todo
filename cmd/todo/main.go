package main

type Task struct {
	name, startDate, endDate string
	isDone                   bool `default:"false"`
}

func NewTask(n string, sd string, ed string) Task {
	return Task{
		name:      n,
		startDate: sd,
		endDate:   ed,
	}
}

func main() {
	var isRunning bool = true

	println("==== Todo ====")

	for isRunning != false {

	}
}
