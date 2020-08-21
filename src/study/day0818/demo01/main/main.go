package main

import (
	"fmt"
	"time"
)

//-------------------------有关Task任务角色的功能---------------------------------
//定义一个任务类型Task
type Task struct {
	f func() error //一个Task里面应该有一个具体的业务，业务名称叫f
}

//创建一个Task任务

func NewTask(arg_f func() error) *Task {
	t := Task{
		f: arg_f,
	}
	return &t
}

//Task也需要一个执行业务的方法
func (t *Task) Execute() {
	t.f() //调用任务中绑定的方法
}

//-------------------------有关协程池Pool角色的功能---------------------------------

//定义一个Pool协程池的类型
type Pool struct {
	//对外的Task入口的EntryChannel
	EntryChannel chan *Task

	//对内的Task队列JobChannel
	JobsChannel chan *Task

	//协程池中最大的worker的数量
	worker_num int
}

//创建Pool的函数
func NewPool(cap int) *Pool {
	p := Pool{
		EntryChannel: make(chan *Task),
		JobsChannel:  make(chan *Task),
		worker_num:   cap,
	}
	//返回Pool
	return &p
}

//协程池创建一个Worker，并且让这个Worker去工作
func (p *Pool) worker(work_ID int) {
	//一个worker的具体工作
	//1 永久的从JobsChannel取任务
	for task := range p.JobsChannel {
		//2 一旦取到任务，就去执行这个任务
		task.Execute()
		fmt.Println("worker ID ", work_ID, "执行完了一个任务")
	}
}

//让协程池开始真正的工作,协程池的一个启动方法
func (p *Pool) run() {
	//1、根据work_num来创建一个worker去工作
	for i := 0; i < p.worker_num; i++ {
		//每个worker都应该是一个goroutine
		go p.worker(i)
	}
	//2、从EntryChannel中去取任务，将取到的任务发送给JobsChannel
	for task := range p.EntryChannel {
		p.JobsChannel <- task
	}
}

//测试
func main() {
	//1、创建一些任务
	t := NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})
	//2、创建一个Pool协程池,4
	p := NewPool(4)

	//3、将这些任务交给协程池Pool
	go func() {
		for {
			p.EntryChannel <- t
		}
	}()
	//4、启动pool
	p.run()

}
