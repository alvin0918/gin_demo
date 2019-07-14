package common

const (

	// 保存任务接口
	JOB_SAVE_DIR = "/cron/jobs/save/"

	// 保存任务事件
	JOB_EVENT_SAVE = 1

	// 删除任务接口
	JOB_DEL_DIR = "/cron/jobs/del/"

	// 删除任务事件
	JOB_EVENT_DEL = 2

	// 强杀任务接口
	JOB_KILLER_DIR = "/cron/jobs/killer/"

	// 强杀任务事件
	JOB_EVENT_KILLER = 3

	// 服务注册接口
	JOB_WORKER_DIR = "/cron/workers"

	// 注册服务事件
	JOB_EVENT_WORKER = 4

)
