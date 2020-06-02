package async

import (
    "github.com/better-go/pkg/log"
    "github.com/better-go/pkg/sync/async"
)

const (
    defaultWorkerNum = 4
    defaultBufSize   = 1024
)

// async task handler
type Dao struct {
    DB    *async.TaskDispatcher // db 操作
    Cache *async.TaskDispatcher // 缓存操作
    MQ    *async.TaskDispatcher // 队列操作
    HTTP  *async.TaskDispatcher // http 请求
    RPC   *async.TaskDispatcher // rpc 请求
}

func New() *Dao {
    return &Dao{
        DB:    async.New("db-task", async.Worker(defaultWorkerNum), async.Buffer(defaultBufSize*10)),
        Cache: async.New("cache-task", async.Worker(defaultWorkerNum), async.Buffer(defaultBufSize*10)),
        MQ:    async.New("mq-task", async.Worker(defaultWorkerNum), async.Buffer(defaultBufSize*10)),
        HTTP:  async.New("http-task", async.Worker(defaultWorkerNum), async.Buffer(defaultBufSize*10)),
        RPC:   async.New("rpc-task", async.Worker(defaultWorkerNum), async.Buffer(defaultBufSize*10)),
    }
}

// need sync here:
func (m *Dao) Close() {
    // group:
    taskGroup := []*async.TaskDispatcher{
        m.DB,
        m.Cache,
        m.MQ,
        m.HTTP,
        m.RPC,
    }

    // batch close:
    for _, tg := range taskGroup {
        if err := tg.Close(); err != nil {
            log.Errorf("async task close error, err=%v", err)
        }
    }
    return
}
