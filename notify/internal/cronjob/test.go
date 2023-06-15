package cronjob

import "fmt"

type TestJob struct {
}

func (a *TestJob) Run() {
	fmt.Println("test")
}
