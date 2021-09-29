package main

import (
	"log"

	"github.com/robfig/cron"
)

type Job01 struct{}
type Job02 struct{}

func (j Job01) Run() {
	log.Printf("Job01 已经执行任务")
}
func (j Job02) Run() {
	log.Printf("Job02 已经执行任务 ")
}
func main() {
	//案例一
	// c := cron.New()
	// spec := "*/5 * * * * ?" //分别是 秒 分 时 日 月 年  //当前案例表示每隔5秒执行一次
	// c.AddFunc(spec, func() {
	// 	log.Println("cron running:", time.Now().Format("2006-01-02 15:04:05"))
	// })
	// c.Start()

	// select {}
	//案例二
	C := cron.New()
	var spec string = "0 */5 * * * ?" //5分钟执行一次
	// var spec string="0 3-59/15 * * * ?"  //表示每小时的第3分中开始每隔15分钟执行一次
	// var spec string="0 0 */1 * * ?"    	//表示每小时执行一次
	C.AddJob(spec, Job01{})
	C.AddJob(spec, Job02{})
	C.Start()
	select {}
}
