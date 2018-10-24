package main

import (
    "fmt"
    "github.com/wxnacy/gopb"
    "time"
)

func main() {
    fmt.Println("")

    bar := pb.New()
    bar.Add(111, 50, "进程1", func(progress float64) int {

        time.Sleep(time.Duration(200) * time.Millisecond)
        return 2
    })

    p := pb.NewProcess(111, 50, "进程2", func(progress float64) int {

        time.Sleep(time.Duration(200) * time.Millisecond)
        return 2
    })
    bar.AddProcess(p)
    bar.Run()
}
