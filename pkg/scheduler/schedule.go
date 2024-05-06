package scheduler

import (
	"log"
	"time"

	"performance-dashboard/pkg/profiles"

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
	config := profiles.GetSettings()

	for	_, taskCfg := range config.Schedule.Task {
		var worker func()error
		switch taskCfg.ID {
		case "jira_project":
			worker = jiraProjectWorker
		case "jira_sprint":
			worker = jiraSprintWorker
		}
		task := tasks.Task{
			TaskFunc:          worker,
			Interval:          taskCfg.Period,
			StartAfter:        time.Now().Add(taskCfg.DelayedStart),
			RunSingleInstance: true,
		}
		// Schedule future execution
		scheduler.AddWithID(taskCfg.ID, &task)

		// Initial worker execution
		if taskCfg.ExecuteOnStartup {
			go func(cfg profiles.TaskConfig) {
				if cfg.DelayedStart > 0 {
					time.Sleep(cfg.DelayedStart)
				}
				worker()
			} (taskCfg)
		}
	}
	
	log.Println("Scheduler activated. The wollowing tasks are scheduled to run:")
	for id, task := range scheduler.Tasks() {
		log.Printf("  - %s (runs every %v, starting at %s)\n", id, task.Interval, task.StartAfter.Format(time.RFC1123))
	}
}
