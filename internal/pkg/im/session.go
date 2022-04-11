package im

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"
)

var once sync.Once

// Sessions 客户端管理实例
var Sessions *session

// session 渠道客户端
type session struct {
	Default *Channel // 默认分组
	Example *Channel // 案例分组

	// 可自行注册其它渠道...
}

func Initialize(ctx context.Context, eg *errgroup.Group) {
	once.Do(func() {
		initialize(ctx, eg)
	})
}

func initialize(ctx context.Context, eg *errgroup.Group) {
	Sessions = &session{
		Default: NewChannel("default", NewNode(10), make(chan *SenderContent, 5<<20)),
		Example: NewChannel("example", NewNode(1), make(chan *SenderContent, 100)),
	}

	eg.Go(func() error {
		return heartbeatManage.Start(ctx)
	})

	eg.Go(func() error {
		return ackManage.Start(ctx)
	})

	eg.Go(func() error {
		return Sessions.Default.Start(ctx)
	})

	eg.Go(func() error {
		return Sessions.Example.Start(ctx)
	})
}
