package etcd

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.etcd.io/etcd/client/v3"

	"github.com/acoderup/goserver/core/logger"
)

/*
 etcd常用操作和数据监听
*/

var globalClient = new(Client)

type (
	InitFunc  func(ctx context.Context, res *clientv3.GetResponse)
	WatchFunc func(ctx context.Context, res *clientv3.WatchResponse)
	funcPair  struct {
		key       string
		initFunc  InitFunc
		watchFunc WatchFunc
	}
)

type Client struct {
	cli       *clientv3.Client
	functions []funcPair
	closed    chan struct{}
	wg        sync.WaitGroup
}

func (c *Client) Ctx() context.Context {
	if c.cli != nil {
		return c.cli.Ctx()
	}
	return context.TODO()
}

func (c *Client) Open(etcdUrl []string, userName, passWord string, dialTimeout time.Duration) error {
	var err error

	if len(etcdUrl) == 0 {
		etcdUrl = []string{"localhost:2379"}
	}
	if dialTimeout == 0 {
		dialTimeout = time.Minute
	}

	c.cli, err = clientv3.New(clientv3.Config{
		Endpoints:            etcdUrl,
		Username:             userName,
		Password:             passWord,
		DialTimeout:          dialTimeout,
		DialKeepAliveTime:    5 * time.Second,
		DialKeepAliveTimeout: 30 * time.Second,
	})

	if err != nil {
		logger.Logger.Errorf("EtcdClient.Open(%v) err:%v", etcdUrl, err)
		return err
	}

	logger.Logger.Infof("EtcdClient.Open(%v) success", etcdUrl)

	c.closed = make(chan struct{})
	return err
}

func (c *Client) Close() error {
	logger.Logger.Infof("EtcdClient.Close()")
	select {
	case <-c.closed:
		return nil
	default:
		close(c.closed)
	}

	if c.cli != nil {
		c.cli.Close()
	}

	c.wg.Wait()

	return nil
}

// PutValue 添加键值对
func (c *Client) PutValue(key, value string) (*clientv3.PutResponse, error) {
	resp, err := c.cli.Put(context.TODO(), key, value)
	if err != nil {
		logger.Logger.Warnf("EtcdClient.PutValue(%v,%v) error:%v", key, value, err)
	}
	return resp, err
}

// GetValue 查询
func (c *Client) GetValue(key string) (*clientv3.GetResponse, error) {
	resp, err := c.cli.Get(context.TODO(), key)
	if err != nil {
		logger.Logger.Warnf("EtcdClient.GetValue(%v) error:%v", key, err)
	}
	return resp, err
}

// DelValue 返回删除了几条数据
func (c *Client) DelValue(key string) (*clientv3.DeleteResponse, error) {
	res, err := c.cli.Delete(context.TODO(), key)
	if err != nil {
		logger.Logger.Warnf("EtcdClient.DelValue(%v) error:%v", key, err)
	}
	return res, err
}

// DelValueWithPrefix 按照前缀删除
func (c *Client) DelValueWithPrefix(prefix string) (*clientv3.DeleteResponse, error) {
	res, err := c.cli.Delete(context.TODO(), prefix, clientv3.WithPrefix())
	if err != nil {
		logger.Logger.Warnf("EtcdClient.DelValueWithPrefix(%v) error:%v", prefix, err)
	}
	return res, err
}

// GetValueWithPrefix 获取前缀
func (c *Client) GetValueWithPrefix(prefix string) (*clientv3.GetResponse, error) {
	resp, err := c.cli.Get(context.TODO(), prefix, clientv3.WithPrefix())
	if err != nil {
		logger.Logger.Warnf("EtcdClient.GetValueWIthPrefix(%v) error:%v", prefix, err)
	}
	return resp, err
}

// WatchWithPrefix 监测前缀创建事件
func (c *Client) WatchWithPrefix(prefix string, revision int64) clientv3.WatchChan {
	if c.cli != nil {
		opts := []clientv3.OpOption{clientv3.WithPrefix(), clientv3.WithCreatedNotify()}
		if revision > 0 {
			opts = append(opts, clientv3.WithRev(revision))
		}
		return c.cli.Watch(clientv3.WithRequireLeader(context.Background()), prefix, opts...)
	}
	return nil
}

// AddFunc 添加监听函数
func (c *Client) AddFunc(key string, initFunc InitFunc, watchFunc WatchFunc) {
	logger.Logger.Infof("EtcdClient.AddFunc(%v)", key)
	fs := funcPair{
		key:       key,
		initFunc:  initFunc,
		watchFunc: watchFunc,
	}
	c.functions = append(c.functions, fs)
}

// Start 重新监听
func (c *Client) Start() {
	select {
	case <-c.closed:
		return
	default:
	}

	logger.Logger.Infof("EtcdClient.Start")
	for _, v := range c.functions {
		c.initAndWatch(v.key, v.initFunc, v.watchFunc)
	}
}

func (c *Client) Restart() {
	go func() {
		c.Close()
		cli := new(Client)
		err := cli.Open(Config.Url, Config.UserName, Config.Password, time.Duration(Config.DialTimeout)*time.Second)
		if err != nil {
			logger.Logger.Errorf("EtcdClient.Restart error:%v", err)
			return
		}
		c.cli = cli.cli
		c.wg = sync.WaitGroup{}
		c.closed = make(chan struct{})
		c.Start()
	}()
}

// initAndWatch 开始监听
func (c *Client) initAndWatch(key string, initFunc InitFunc, watchFunc WatchFunc) {
	res, err := c.GetValueWithPrefix(key)
	if err != nil {
		panic(fmt.Sprintf("initAndWatch WithPrefix(%v) err:%v", key, err))
	}

	logger.Logger.Infof("etcd initFunc WithPrefix(%v) start", key)
	initFunc(c.Ctx(), res)

	ctx, _ := context.WithCancel(c.cli.Ctx())
	vision := int64(-1)
	if res.Header != nil {
		vision = res.Header.Revision
	}
	c.goWatch(ctx, vision+1, key, watchFunc)
}

// goWatch 异步监听
func (c *Client) goWatch(ctx context.Context, revision int64, prefix string, f WatchFunc) {
	c.wg.Add(1)
	go func() {
		defer func() {
			c.wg.Done()
		}()
		defer func() {
			if err := recover(); err != nil {
				logger.Logger.Errorf("etcd watch WithPrefix(%v) panic:%v", prefix, err)
			}
			logger.Logger.Warnf("etcd watch WithPrefix(%v) quit!!!", prefix)
		}()
		x, _ := context.WithCancel(ctx)
		var times int64
		for {
			times++
			logger.Logger.Warnf("etcd watch WithPrefix(%v) base revision %v start[%v]!!!", prefix, revision, times)
			rch := c.WatchWithPrefix(prefix, revision)
			for {
				skip := false
				select {
				case <-c.closed:
					return
				case _, ok := <-ctx.Done():
					if !ok {
						return
					}
				case resp, ok := <-rch:
					if !ok {
						logger.Logger.Warnf("etcd watch WithPrefix(%v) be closed", prefix)
						skip = true
						break
					}
					if resp.Header.Revision > revision {
						revision = resp.Header.Revision
					}
					if resp.Canceled {
						logger.Logger.Warnf("etcd watch WithPrefix(%v) be closed, reason:%v", prefix, resp.Err())
						skip = true
						break
					}
					if err := resp.Err(); err != nil {
						logger.Logger.Warnf("etcd watch WithPrefix(%v) err:%v", prefix, resp.Err())
						continue
					}
					if len(resp.Events) == 0 {
						continue
					}

					logger.Logger.Tracef("@goWatch %v changed, header:%#v", prefix, resp.Header)
					f(x, &resp)
				}

				if skip {
					break
				}
			}
			time.Sleep(time.Second)
		}
	}()
}
