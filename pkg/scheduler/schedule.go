package scheduler

import (
	"log"
	"time"

	"github.com/madflojo/tasks"
)

var scheduler *tasks.Scheduler

func initScheduler() *tasks.Scheduler {
	if scheduler == nil {
		scheduler = tasks.New()
	} else {
		scheduler.Stop()
	}
	return scheduler
}

func Schedule() {

	scheduler := initScheduler()

	duration := time.Duration(15 * time.Minute)
	timeGap := time.Duration(10 * time.Second)

	projectTask := tasks.Task{
		TaskFunc:          jiraCoreWorker,
		Interval:          duration,
		RunSingleInstance: true,
	}

	sprintTask := tasks.Task{
		TaskFunc:          jiraAgileWorker,
		StartAfter:        time.Now().Add(timeGap),
		Interval:          duration,
		RunSingleInstance: true,
	}

	// Schedule future execution
	scheduler.AddWithID("Project update task", &projectTask)
	scheduler.AddWithID("Active sprint update task", &sprintTask)

	log.Println("Scheduler activated. The wollowing tasks are scheduled to run:")
	for id, task := range scheduler.Tasks() {
		log.Printf("  - %s (runs every %v)\n", id, task.Interval)
	}

	// Initial workers execution
	go jiraCoreWorker()
	time.Sleep(timeGap)
	go jiraAgileWorker()
}
