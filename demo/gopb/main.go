package main

import (
    "fmt"
    "github.com/wxnacy/gopb"
    "time"
)

func main() {
    fmt.Println("")

    bar := pb.New()
    bar.Add(110, 33, "进程1", func(progress float64) int {

        time.Sleep(time.Duration(500) * time.Millisecond)
        return 11
    })


    p := pb.NewProcess(111, 33, "进程2", func(progress float64) int {

        time.Sleep(time.Duration(200) * time.Millisecond)
        return 22
    })
    bar.AddProcess(p)
    bar.AddDefaultProcess("进程3", func(progress float64) int {

        time.Sleep(time.Duration(200) * time.Millisecond)
        return 30
    })
    bar.Run()
}
