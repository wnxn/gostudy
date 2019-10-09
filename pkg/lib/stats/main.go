package main

import (
	"fmt"
	"io/ioutil"
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/volume"
	"k8s.io/kubernetes/pkg/volume/util/fs"
	"os"
)

func main() {
	klog.InitFlags(nil)
	fullpath, err := ioutil.TempDir("", "gostudy-stats")
	if err != nil{
		klog.Info(err)
	}
	defer os.Remove(fullpath)
	fmt.Println(fullpath)
	metrics,err := volume.NewMetricsStatFS(fullpath).GetMetrics()
	if err != nil{
		klog.Info(err)
	}
	klog.Info(metrics.Available.Value(), metrics.Capacity.Value(), metrics.Used.Value())
	klog.Info(metrics.Inodes.Value(), metrics.InodesFree.Value(), metrics.InodesUsed.Value())

	bAvail, bCap, bUsage, iTotal, iFree, iUsage, err :=fs.FsInfo(fullpath)
	if err != nil{
		klog.Info(err)
	}
	klog.Info(bAvail, bCap, bUsage, iTotal, iFree, iUsage)
}
