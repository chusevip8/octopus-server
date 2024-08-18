package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/service"
	"go.uber.org/zap"
	"strconv"
	"time"
)

func MonitorIntervalTask() {
	go func() {
		global.GVA_LOG.Info("interval task monitor start")
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				runIntervalTask()
			}
		}
	}()
}

func runIntervalTask() {
	tasks, err := service.TaskService.FindIntervalTasks()
	if err != nil {
		global.GVA_LOG.Error("find interval task fail", zap.Error(err))
		return
	}
	for _, task := range tasks {
		if intervalCheck(task) {
			global.GVA_LOG.Info("interval task can run", zap.Any("task", task))
			TryPushIntervalTask(task)
		}
	}
}
func intervalCheck(task octopus.Task) bool {
	if task.FinishAt == nil {
		return true
	}

	now := time.Now()
	duration := now.Sub(*task.FinishAt).Minutes()

	intervalSetupID := strconv.Itoa(int(task.TaskParams.TaskSetupId))
	intervalTaskSetup, err := service.IntervalTaskSetupService.GetIntervalTaskSetup(intervalSetupID)
	if err != nil {
		global.GVA_LOG.Error("Failed to check interval task setup",
			zap.String("Task ID", strconv.Itoa(int(task.ID))),
			zap.Error(err))
		return false
	}
	return uint(duration) >= intervalTaskSetup.IntervalMin
}
