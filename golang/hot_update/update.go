package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

type confModel struct {
	HTTPPort    string `json:"http_port"`
	Ver         string `json:"ver"`
	DelaySecond int    `json:"delay_second"`
	BinName     string `json:"bing_name"`
}

var restartChan = make(chan bool)
var watchPath = "watch.conf"
var logFile = "log.txt"
var wait sync.WaitGroup

var curVer string = "init"
var curPid string

func main() {
	simpleLog("启动 main")

	wait.Add(3)
	go watch()
	go restartWeb()
	go web()
	go func() {
		for i := 0; i < 50; i++ {
			simpleLog("======> echo ------>> ", i, " ver: ", curVer)
			time.Sleep(time.Second * 1)
		}
	}()
	wait.Wait()
}

func watch() {
	simpleLog("启动watch")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		simpleLog(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				simpleLog("监听到事件event: ", event, time.Now())
				if event.Op&fsnotify.Write == fsnotify.Write {
					simpleLog("检测到文件被修改")
					restartChan <- true
				}
				restartChan <- true
				//if event.Op&fsnotify.Remove == fsnotify.Remove {
				//simpleLog("检测到文件被删除")
				//restartChan <- true
				//}
			case <-watcher.Errors:
				simpleLog("unknown err: ", watcher.Errors)
			}
		}
	}()

	if err := watcher.Add(watchPath); err != nil {
		simpleLog(err)
		log.Fatal(err)
	}
	<-make(chan bool)
}

func restartWeb() {
	simpleLog("启动 restartWeb")
	for {
		<-restartChan
		simpleLog("收到通知和准备 restartWeb")

		if m := loadConf(); m == nil {
			simpleLog("conf 解析失败")
		} else {
			curVer = m.Ver
			binName := m.BinName

			time.Sleep(time.Duration(m.DelaySecond) * time.Second)

			curPid = strconv.Itoa(syscall.Getpid())
			simpleLog("当前进程pid: ", curPid)

			simpleLog("chmod 777", binName)
			exec.Command("chmod", "777", binName).Run()

			simpleLog("kill -1", curPid)
			exec.Command("kill", "-1", curPid).Run()
			simpleLog("新进程pid：", syscall.Getpid())

			simpleLog("更新完成", curPid, binName)
		}
	}
}

func web() {
	g := gin.New()

	g.GET("/hello", func(c *gin.Context) {
		pid := strconv.Itoa(syscall.Getpid())
		simpleLog("/hello", pid)
		c.JSON(200, gin.H{"message": "Hello gin new2!", "ver": curVer, "pid": pid})
	})

	g.GET("/filelog", func(c *gin.Context) {
		bytes, err := ioutil.ReadFile(logFile)
		if err != nil {
			c.String(200, "打开失败")
		} else {
			var text string = string(bytes)
			c.String(200, text)
		}
	})

	m := loadConf()
	curVer = m.Ver
	s := endless.NewServer(":"+m.HTTPPort, g)
	err := s.ListenAndServe()
	if err != nil {
		simpleLog("server err:", err)
	}

	simpleLog("Server on " + m.HTTPPort + " stopped")
	os.Exit(0)
}

func loadConf() *confModel {
	c, err := ioutil.ReadFile(watchPath)
	if err != nil {
		simpleLog(watchPath, "读取失败 error: ", err)
		return nil
	}
	var conf confModel
	if err := json.Unmarshal(c, &conf); err != nil {
		simpleLog("conf 解析失败", err)
		return nil
	}
	return &conf
}

func simpleLog(a ...interface{}) {
	fmt.Println(a...)
	msg := fmt.Sprintln(a...)

	msg = time.Now().Format("2006-01-02 15:04:05") + "|" + msg

	f3, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("打开文件错误", err)
	} else {
		defer f3.Close() // 延迟到最后关闭
		s1 := make([]byte, 1024)
		f3.Read(s1)

		f3.WriteString(msg)
	}
}
