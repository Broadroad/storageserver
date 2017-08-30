package main

import (
	//"github.com/broadroad/storageserver/ss"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"runtime"
)

const (
	Name    string = "storageserver"
	Version string = "1.0"
)

func MaxParallelism() int {
	// GOMAXPROCS sets the maximum number of CPUs that can be executing
	// simultaneously and returns the previous setting. If n < 1, it does not
	// change the current setting.
	// The number of logical CPUs on the local machine can be queried with NumCPU..
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	glog.Infof("GOMAXPROCS: %s, NumbCPU : %s", maxProcs. NumCPU)
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
	defer glog.Flush()
	glog.Info("*************************")
	glog.Infof(" name: [%s] version:[%s]", Name, Version)
	
	Start()
	//s := ss.NewSS()
	//s.Setup(2, 2)
}
