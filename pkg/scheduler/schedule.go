package scheduler

import (
	"log"
	"time"

	"github.com/madflojo/tasks"
	"performance-dashboard/pkg/profiles"
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
	config := profiles.GetSettings()

	duration := config.Schedule.CoreTask.Period
	timeGap := config.Schedule.SecondaryTasks.DelayedStart

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
	if config.Schedule.CoreTask.ExecuteOnStartup {
		go jiraCoreWorker()
		time.Sleep(timeGap)
		go jiraAgileWorker()
	}
}
