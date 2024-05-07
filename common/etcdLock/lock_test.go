package etcdlock

import (
	"fmt"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	go A()
	go B()
	time.Sleep(4 * time.Second)
	fmt.Println("end")
}
func A() {
	start := time.Now()
	session := NewSession()
	locker := NewLocker(session)
	locker.Lock()
	defer locker.Unlock()
	fmt.Println("aaa", time.Since(start).Milliseconds())
	time.Sleep(1 * time.Second)
	fmt.Println("aaa")
}
func B() {
	start := time.Now()
	session := NewSession()
	locker := NewLocker(session)
	locker.Lock()
	defer locker.Unlock()
	fmt.Println("bbb", time.Since(start).Milliseconds())
	time.Sleep(1 * time.Second)
	fmt.Println("bbb")
}
