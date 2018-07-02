package main

import (
	//"github.com/broadroad/storageserver/ss"
	"flag"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

// ContextHook is a hook for logrus.
type ContextHook struct{}

// Levels returns the whole levels.
func (hook ContextHook) Levels() []log.Level {
	return log.AllLevels
}

// Fire helps logrus record the related file, function and line.
func (hook ContextHook) Fire(entry *log.Entry) error {
	pc := make([]uintptr, 3, 3)
	cnt := runtime.Callers(1, pc)

	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i] - 1)
		fmt.Println(fu.Name())
		name := fu.Name()
		if !strings.Contains(name, "github.com/Sirupsen/log") {
			file, line := fu.FileLine(pc[i] - 1)
			entry.Data["file"] = path.Base(file)
			entry.Data["func"] = path.Base(name)
			entry.Data["line"] = line
			break
		}
	}
	return nil
}

const (
	Name    string = "storageserver"
	Version string = "1.0"
	logFile string = "storageserver.log"
)

func init() {
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	log.SetFormatter(&log.TextFormatter{})
	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	log.SetOutput(os.Stdout)
	//设置最低loglevel
	log.SetLevel(log.InfoLevel)
	if err != nil {
		// Cannot open log file. Logging to stderr
		fmt.Println(err)
	} else {
		log.SetOutput(f)
	}

	log.AddHook(ContextHook{})
	log.SetLevel(log.DebugLevel)
}

func MaxParallelism() int {
	// GOMAXPROCS sets the maximum number of CPUs that can be executing
	// simultaneously and returns the previous setting. If n < 1, it does not
	// change the current setting.
	// The number of logical CPUs on the local machine can be queried with NumCPU..
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	log.Info("GOMAXPROCS: %d, NumbCPU : %d", maxProcs, numCPU)
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

func Start() {
	//TODO Start Storage Server
}

func main() {
	flag.Parse()
	MaxParallelism()
	// ensure log is writen before server quit
	log.Info("*************************")
	log.WithFields(log.Fields{
		"Name":    Name,
		"Version": Version,
	}).Info("test")

	Start()
	//s := ss.NewSS()
	//s.Setup(2, 2)
}
