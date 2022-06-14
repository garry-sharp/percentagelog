package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/garry-sharp/percentagelog"
)

type NetworkLoader struct {
	FileName          string
	Version           string
	totalSize         uint64
	totalAmountLoaded uint64
}

func (n *NetworkLoader) String() string {
	if n.Version != "" {
		return fmt.Sprintf("%s - %s", n.FileName, n.Version)
	} else {
		return n.FileName
	}
}

func (n *NetworkLoader) Percentage() float32 {
	//fmt.Println(n.String(), n.totalAmountLoaded, n.totalSize, float32(n.totalAmountLoaded)/float32(n.totalSize), float32(n.totalAmountLoaded)/float32(n.totalSize)*float32(100))
	return float32(n.totalAmountLoaded) / float32(n.totalSize) * float32(100)
}

func (n *NetworkLoader) loadMore(duration time.Duration) {
	//Select a speed between 50KB/s and 3 MB/s
	fiftykbs := 1024 * 50
	threembs := 1024 * 1024 * 3
	bps := rand.Intn(threembs-fiftykbs) + fiftykbs
	time.Sleep(duration)
	amountLoaded := uint64(float64(bps) * float64(duration.Seconds()))

	if n.totalAmountLoaded+amountLoaded > n.totalSize {
		n.totalAmountLoaded = n.totalSize
	} else {
		n.totalAmountLoaded += amountLoaded
	}
}

func main() {
	loaders := []*NetworkLoader{
		{FileName: "hello_world.txt", totalSize: 100},
		{FileName: "my_cat.jpg", totalSize: 3124826},
		{FileName: "amazing_go_code.go", Version: "v.1.2.2", totalSize: 206811},
		{FileName: "MyPasswords.kdbx", totalSize: 94542},
		{FileName: "Docker.dmg", totalSize: 581860518},
		{FileName: "Application.zip", totalSize: 5655172},
		{FileName: "mysql-workbench-community-8.0.29-macos-x86_64.dmg", totalSize: 118791692},
		{FileName: "offsetexplorer.dmg", totalSize: 61285264},
	}

	var wg sync.WaitGroup
	wg.Add(len(loaders))
	for _, loader := range loaders {
		go func(l *NetworkLoader) {
			for {
				loadInterval := 100
				l.loadMore(time.Duration(loadInterval) * time.Millisecond)
				if l.totalAmountLoaded == l.totalSize {
					wg.Done()
					return
				}
			}
		}(loader)
	}

	var printables []percentagelog.Printable
	for _, l := range loaders {
		printables = append(printables, l)
	}

	fmt.Println("Loading")
	percentagelog.PrintUntilFinished(printables...)

	wg.Wait()
	fmt.Println("Finished")
}
