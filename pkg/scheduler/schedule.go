package scheduler

import (
	"log"
	"time"

	"github.com/madflojo/tasks"
	"golang.org/x/exp/maps"
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

	delay := time.Duration(1000 * time.Millisecond)
	duration := time.Duration(5000 * time.Millisecond)
	
	projectTask := tasks.Task{
		TaskFunc: updateProject,
		Interval: duration,
		RunSingleInstance: true,
	}

	sprintTask := tasks.Task{
		TaskFunc: updateSprint,
		StartAfter: time.Now().Add(delay),
		Interval: duration,
		RunSingleInstance: true,
	}

	scheduler.AddWithID("Project update task", &projectTask)
	scheduler.AddWithID("Active sprint update task", &sprintTask)

	log.Println("Scheduler activated. Scheduled tasks list:")
	for i, taskId := range maps.Keys(scheduler.Tasks()) {
		log.Printf("[%d] %s\n", i + 1, taskId)
	}
}
