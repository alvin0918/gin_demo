package etcd

import (
	"github.com/alvin0918/gin_demo/core/commin/log"
	"go.etcd.io/etcd/clientv3"
	"encoding/json"
	"context"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"github.com/alvin0918/gin_demo/common"
	"gopkg.in/ini.v1"
	"github.com/alvin0918/gin_demo/core/config"
	"fmt"
	"time"
)

// /cron/workers/
type WorkerMgr struct {
	client *clientv3.Client
	kv clientv3.KV
	lease clientv3.Lease
}

var EtcdWorkerMgr *WorkerMgr

func init()  {
	var (
		conf clientv3.Config
		section *ini.Section
		sections *ini.Section
		etcdEndpoints []string
		err error
		num int
		str string
		etcdDialTimeout int64
		client *clientv3.Client
		kv clientv3.KV
		lease clientv3.Lease
	)

	if section, err = config.GetSection("Etcd"); err != nil {
		log.TracePrintf("Etcd", err.Error())
	}

	if num, err = section.Key("num").Int(); err != nil {
		log.TracePrintf("Etcd", err.Error())
	}

	etcdEndpoints = []string{}

	for i := 0; i < num; i++ {
		str = fmt.Sprintf("Etcd_%d", i)

		if sections, err = config.GetSection(str); err != nil {
			log.TracePrintf("Etcd", err.Error())
		}

		etcdEndpoints = append(etcdEndpoints, sections.Key("etcdIPAndPort").String())

	}

	if etcdDialTimeout, err = section.Key("etcdDialTimeout").Int64(); err != nil {
		log.TracePrintf("Etcd", err.Error())
	}

	// 初始化配置
	conf = clientv3.Config{
		Endpoints: etcdEndpoints, // 集群地址
		DialTimeout: time.Duration(etcdDialTimeout) * time.Millisecond, // 连接超时
	}

	// 建立连接
	if client, err = clientv3.New(conf); err != nil {
		log.TracePrintf("Etcd", err.Error())
	}

	// 得到KV和Lease的API子集
	kv = clientv3.NewKV(client)
	lease = clientv3.NewLease(client)

	EtcdWorkerMgr = &WorkerMgr{
		client:client,
		kv:kv,
		lease:lease,
	}

}

// 保存Job
func (jobMgr *WorkerMgr) SaveJob(job *common.Job) (oldJob *common.Job, err error) {

	var (
		jobKey    string
		jobValue  []byte
		putResp   *clientv3.PutResponse
		oldJobObj common.Job
		ctx context.Context
		cancel context.CancelFunc
	)

	// 配置KEY
	jobKey = common.JOB_SAVE_DIR + job.Name

	// 解析任务信息
	if jobValue, err = json.Marshal(job); err != nil {
		log.TracePrintf("Etcd", err.Error())
		return
	}

	// 设置访问ETCD的超时时间
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)

	// 保存到etcd
	if putResp, err = jobMgr.kv.Put(ctx, jobKey, string(jobValue), clientv3.WithPrevKV()); err != nil {
		return
	}

	// 操作完毕，取消ETCD
	cancel()

	// 如果更新成功，则返回旧职
	if putResp.PrevKv != nil {
		// 对旧值反序列化
		if err = json.Unmarshal(putResp.PrevKv.Value, &oldJobObj); err != nil {
			log.TracePrintf("Etcd", err.Error())
			return
		}
		oldJob = &oldJobObj
	}

	return
}

// 列举任务
func (jobMgr *WorkerMgr) ListJobs() (jobList []*common.Job, err error) {
	var (
		dirKey string
		getResp *clientv3.GetResponse
		kvPair *mvccpb.KeyValue
		job *common.Job
		ctx context.Context
		cancel context.CancelFunc
	)

	// 任务保存的目录
	dirKey = common.JOB_SAVE_DIR

	// 设置访问ETCD的超时时间
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)

	// 获取目录下所有任务信息
	if getResp, err = jobMgr.kv.Get(ctx, dirKey, clientv3.WithPrefix()); err != nil {
		return
	}

	// 操作完毕，取消ETCD
	cancel()

	// 初始化数组空间
	jobList = make([]*common.Job, 0)

	// 遍历所有任务, 进行反序列化
	for _, kvPair = range getResp.Kvs {
		job = &common.Job{}
		if err =json.Unmarshal(kvPair.Value, job); err != nil {
			err = nil
			continue
		}
		jobList = append(jobList, job)
	}
	return
}





