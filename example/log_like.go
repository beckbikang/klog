package main

import (
	"os"
	"io"
	"fmt"
	"log"
	"bytes"
	"strconv"
	"sync"
	"time"
)

func main() {
	//1 基本的写入
	//tlog()
	//2 写一个类似输出的日志的类
	tlog2()
}

func tlog2(){
	logTest := NewLogTest(new(WriterData))
	logTest.WriteStr("abc")

	var wg sync.WaitGroup
	for i:=0;i< 100;i++{
		wg.Add(1)
		go writeLog(logTest, "i="+strconv.Itoa(i)+"\n", &wg)
	}
	time.Sleep(time.Second *1)
	wg.Wait()

	defer logTest.destroy()

}

func writeLog(logTest *LogTest, str string,wg *sync.WaitGroup){
	wg.Done()
	logTest.WriteStr(str)

}

type LogTest struct {
	logdata chan  interface{}
	writer io.Writer
	quiteChan chan struct{}
}

func NewLogTest(w io.Writer) *LogTest{
	logTest := &LogTest{
		writer:w,
		logdata:make(chan  interface{}, 1024),
		quiteChan:make(chan struct{}),
	}
	go logTest.start()
	return logTest
}

func (logTest *LogTest)start(){
	//获取日志数据
LOOP:
	for{
		select {
		case data := <-logTest.logdata:
			byteData,ok := data.([]byte)
			if ok{
				logTest.writer.Write(byteData)
			}
		case <- logTest.quiteChan:
			break LOOP
		}
	}
	//处理剩余的部分
	for{
		if (len(logTest.logdata)) == 0{
			break
		}
		data := <-logTest.logdata
		byteData,ok := data.([]byte)
		if ok{
			logTest.Write(byteData)
		}
	}
}

func (logTest *LogTest) destroy(){
	close(logTest.logdata)
	close(logTest.quiteChan)
}

func (logTest *LogTest) close(){
	logTest.quiteChan <- struct{}{}
}

func (logTest *LogTest)Write(data []byte){
	logTest.logdata <- data
}

func (logTest *LogTest)WriteStr(data string){
	logTest.Write([]byte(data))
}


/**
    定义一个可消费的log

 */

type WriterData struct{
}

func (writerData *WriterData)Write(p []byte) (n int, err error){
	return os.Stdout.Write(p)
}


/**
日志的基本使用
 */
func tlog(){

	var buf bytes.Buffer
	log.New(&buf,"",log.Ltime|log.Ldate)
	log.Print("abc")

	fmt.Println(&buf)
}