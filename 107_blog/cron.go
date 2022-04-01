package main

/*
import (
	"log"
	"time"

	"github.com/robfig/cron"

	"lianxi/107_blog/models"
)

func main() {
	log.Println("Starting...")
	//会根据本地时间创建一个新（空白）的 Cron job runner
	c := cron.New()
	//AddFunc 会向 Cron job runner 添加一个 func ，以按给定的时间表运行
	//会首先解析时间表，如果填写有问题会直接 err,无误则将func添加到Schedule队列中等待执行
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		//清除所有删除的标签
		models.CleanAllTag()
	})
	//向Cron job runner添加一个方法，按照给定的时间运行
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		//清除所有的文章
		models.CleanAllArticle()
	})
	//在当前执行的程序中启动 Cron 调度程序。其实这里的主体是
	//goroutine + for + select + timer 的调度控制哦
	c.Start()

	//创建一个定时器,到达时间后发生一个channel消息
	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			//重置定时器，重新开始定时
			t1.Reset(time.Second * 10)

		}
	}

}
*/
