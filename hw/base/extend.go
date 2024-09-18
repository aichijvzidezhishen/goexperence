package base

import "time"

type Task struct {
	TaskExec
	Status    int
	Name      string
	StartTime int64
	EndTime   int64
}

func (t Task) start() {
	t.StartTime = time.Now().Unix()

}

func (t Task) end() {
	t.EndTime = time.Now().Unix()
}

type TaskExec interface {
	exec()
}

type TaskA struct {
	Task
	ParamA string
}

type TaskB struct {
	Task
	ParamA string
}

func (ta TaskA) exec() {
	// 执行任务A的逻辑
	ta.start()
	ta.end()
}

func (TaskB) exec() {
	// 执行任务B的逻辑
}

// 执行任务，问题点在入参这里，这样写入参就只能是Task，不能传TaskA和TaskB，
// 用go有什么办法可以实现像java一样传子类做为入参
func Exec(task TaskExec) {
	// // 修改任务状态为执行中
	// task.Status = 1
	// // 记录任务开始执行时间
	// task.StartTime = time.Now().Unix()
	// // 这里还会有很共性逻辑处理被省略掉了

	// // 执行具体任务的逻辑
	// task.exec()
	// // 修改任务状态为执行结束
	// task.Status = 2
	// // 记录任务执行结束时间
	// task.EndTime = time.Now().Unix()
	// // 这里也是还会有很共性逻辑处理被省略掉了
	task.exec()
}

// func main() {
// 	ta := TaskA{}
// 	tb := TaskB{}

// 	// 这里要怎么才能传子类到函数里面执行？？？？
// 	Exec(&ta)
// 	Exec(&tb)
// }
