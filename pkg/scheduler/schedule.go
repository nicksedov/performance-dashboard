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
		log.Println("A new scheduler for periodic tasks initialized...")
		scheduler = tasks.New()
	} else {
		log.Println("Scheduler for periodic tasks is stopping...")
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
			ErrFunc:           func(err error) { onTaskError(taskCfg.ID, err) },
			Interval:          taskCfg.Period,
			StartAfter:        time.Now().Add(taskCfg.DelayedStart),
			//RunSingleInstance: true,
		}
		// Schedule future execution
		err := scheduler.AddWithID(taskCfg.ID, &task)
		if err != nil {
			log.Panicf("Error scheduling periodic task '%s'\n", taskCfg.ID)
		}

		// Initial worker execution
		if taskCfg.ExecuteOnStartup {
			log.Printf("Task '%s' is scheduled to execute on startup\n", taskCfg.ID)
			go func(currentTaskCfg profiles.TaskConfig) {
				if currentTaskCfg.DelayedStart > 0 {
					time.Sleep(currentTaskCfg.DelayedStart)
				}
				worker()
				log.Printf("Initial execution complete for task '%s'\n", currentTaskCfg.ID)
			} (taskCfg)
		}
	}
	
	log.Println("Scheduler activated. The wollowing tasks are scheduled to run:")
	for id, task := range scheduler.Tasks() {
		log.Printf("  - %s (runs every %v, starting at %s)\n", id, task.Interval, task.StartAfter.Format(time.RFC1123))
	}
}

func onTaskError(taskId string, err error) {
	log.Printf("Error running task '%s', %s", taskId, err.Error())
}