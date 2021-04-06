package gors

import (
	"context"
	"log"
	"sync"
)

type Runnable func(ctx context.Context)

type Task struct {
	f        Runnable
	holder   *Worker
	received bool
}

func (t *Task) Run(ctx context.Context) {
	//before
	log.Printf("holder: %p", t.holder)
	t.f(ctx)
	//after
}

const (
	RUNNING = iota
	WAITING
)

type Worker struct {
	g       *GoPool
	status  int // running, waiting
	ctx     context.Context
	cfunc   context.CancelFunc
	taskCh  <-chan *Task
	running *Task
	sync.Mutex
}

//是否已被取消
func (w *Worker) testCancel() bool {
	select {
	case <-w.ctx.Done():
		//clear
		w.clear()
		return true
	default:
		return false
	}
}

func (w *Worker) clear() {

}

func (w *Worker) Start() {
	log.Printf("worker start %p", w)
	for {
		select {
		case <-w.ctx.Done():
			//clear
			w.clear()
		case t := <-w.taskCh:
			if w.testCancel() {
				return
			}
			w.status = RUNNING
			w.running = t
			t.holder = w
			t.received = true
			//recv task and push it to localq, then try to schedule
			t.Run(w.ctx)
			w.g.complete()
			w.status = WAITING
		}
	}
}

type GoPool struct {
	sync.Mutex
	ctx               context.Context
	cfunc             context.CancelFunc
	active, maxActive int
	rest              int
	wg                sync.WaitGroup
	workers           []*Worker
	taskCh            chan *Task
}

func NewGoPool(maxActive, taskBufSize int) *GoPool {
	g := new(GoPool)
	g.ctx, g.cfunc = context.WithCancel(context.Background())
	g.maxActive = maxActive
	g.taskCh = make(chan *Task, taskBufSize)
	g.workers = make([]*Worker, 0, maxActive)
	return g
}

func (g *GoPool) Push(f Runnable) {
	g.Lock()
	g.rest++
	g.wg.Add(1)
	if g.rest > g.active {
		g.Unlock()
		g.setWorker()
		g.Lock()
	}
	g.Unlock()
	t := new(Task)
	t.f = f
	g.taskCh <- t
}

func (g *GoPool) WaitAll() {
	g.wg.Wait()
}

func newWorker(p *GoPool) *Worker {
	w := new(Worker)
	w.g = p
	w.taskCh = p.taskCh
	w.status = WAITING
	w.ctx, w.cfunc = context.WithCancel(p.ctx)
	return w
}

func (g *GoPool) setWorker() {
	g.Lock()
	defer g.Unlock()
	if g.active >= g.maxActive {
		return
	}
	w := newWorker(g)
	g.workers = append(g.workers, w)
	g.active++
	go w.Start()
}

func (g *GoPool) complete() {
	g.Lock()
	g.rest--
	g.wg.Done()
	g.Unlock()
}

func (g *GoPool) Rest() int {
	g.Lock()
	defer g.Unlock()
	return g.rest
}

func (g *GoPool) StopAll() {
	g.cfunc()
	g.ctx, g.cfunc = context.WithCancel(context.Background())
}
