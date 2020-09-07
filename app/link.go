package app

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Link struct {
	ID       int
	head     *Collector
	out      chan string
	stopCh   chan struct{}
	lock     *sync.Mutex
	logCache []*LogContainer
}

type LogContainer struct {
	Timestamp string `json:"timestamp"`
	Data      string `json:"data"`
}

func NewLink(id int) *Link {
	return &Link{
		ID:     id,
		out:    make(chan string),
		stopCh: make(chan struct{}),
		lock:   &sync.Mutex{},
	}
}

func (l *Link) StopCh() chan struct{} {
	return l.stopCh
}

func (l *Link) Add(c *Collector) {
	if l.head == nil {
		l.head = c
		return
	}
	cur := l.head
	if cur.next != nil {
		cur = cur.next
	}
	cur.next = c
}

func (l *Link) Run(seed string) {
	ch := make(chan string, 1)
	ch <- seed
	close(ch)

	cur := l.head
	for cur != nil {
		var chs []chan string
		for i := 0; i < cur.runtime; i++ {
			chs = append(chs, cur.collect(ch))
		}
		ch = mergeCH(chs...)
		cur = cur.next
	}
	l.out = ch
	go func() {
		for v := range l.out {
			l.putLog(v)
			fmt.Println(v)
		}
	}()
}

func (l *Link) Abort() {
	close(l.stopCh)
}

func (l *Link) GetLog() []*LogContainer {
	l.lock.Lock()
	defer l.lock.Unlock()

	res := l.logCache
	l.logCache = []*LogContainer{}
	return res
}

func (l *Link) CheckOver() bool {
	cur := l.head
	for cur != nil {
		if atomic.LoadInt32(&cur.over) == 1 {
			return false
		}
		cur = cur.next
	}
	return true
}

func mergeCH(chs ...chan string) chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	merge := func(in chan string) {
		defer wg.Done()
		for v := range in {
			out <- v
		}
	}
	wg.Add(len(chs))
	for _, ch := range chs {
		go merge(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func (l *Link) putLog(data string) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.logCache = append(l.logCache, &LogContainer{
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Data:      data,
	})
}
