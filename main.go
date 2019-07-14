package main

import (
	_ "github.com/alvin0918/gin_demo/core/commin/db/etcd"
	"encoding/json"
	"github.com/alvin0918/gin_demo/core/commin/log"
	"github.com/alvin0918/gin_demo/core/commin/db/etcd"
	"fmt"
	"github.com/alvin0918/gin_demo/common"
)

func main()  {

	var (
		job        common.Job
		postJob    string
		err        error
		jobList    []*common.Job
		oldJobList *common.Job
	)

	postJob = "{\"name\":\"test\",\"command\":\"ACD\",\"cronExpr\":\"cronExpr\"}"

	if err = json.Unmarshal([]byte(postJob), &job); err != nil {
		log.TracePrintf("Test", err.Error())
	}


	if oldJobList, err = etcd.EtcdWorkerMgr.SaveJob(&job); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(oldJobList)

	if jobList, err = etcd.EtcdWorkerMgr.ListJobs(); err != nil {
		fmt.Println(jobList)
	}

	for index, v := range jobList{
		fmt.Println(index)
		fmt.Println(v.Name)
		fmt.Println(v.Command)
		fmt.Println(v.CronExpr)
	}

	//core.Run()
}



















