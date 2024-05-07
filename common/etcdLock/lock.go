package etcdlock

import (
	"log"
	"sync"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

func NewClient() *clientv3.Client {
	//创建客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"192.168.44.132:20000", "192.168.44.132:20002", "192.168.44.132:20004"},
	})
	if err != nil {
		log.Fatal(err)
	}
	return cli
}
func NewSession() *concurrency.Session {
	session, err := concurrency.NewSession(NewClient(), concurrency.WithTTL(5))
	if err != nil {
		log.Fatal(err)
	}
	return session
}
func NewLocker(session *concurrency.Session) sync.Locker {
	//利用会话，指定一个前缀创建锁
	return concurrency.NewLocker(session, "/etcdLock/")
}
