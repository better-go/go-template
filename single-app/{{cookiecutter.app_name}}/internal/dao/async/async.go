package async

import (
    "github.com/better-go/pkg/log"
    "github.com/better-go/pkg/sync/async"
)

const (
    defaultWorkerNum = 4
    defaultBufSize   = 1024
)

type Dao struct {
    HTTP *async.TaskDispatcher
}

func New() *Dao {
    return &Dao{
        HTTP: async.New("http-task", async.Worker(defaultWorkerNum), async.Buffer(defaultBufSize*10)),
    }
}

// need sync here:
func (m *Dao) Close() {
    // http:
    if err := m.HTTP.Close(); err != nil {
        log.Errorf("async http task close error, err=%v", err)
    }
    return
}
