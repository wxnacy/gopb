package main

import (
    "fmt"
    "github.com/wxnacy/gopb"
    "time"
)

func main() {
    fmt.Println("")

    bar := pb.New()
    bar.Add(111, 33, "进程1", func(progress float64) int {

        time.Sleep(time.Duration(100) * time.Millisecond)
        return 11
    })

    p := pb.NewProcess(110, 33, "进程2", func(progress float64) int {

        time.Sleep(time.Duration(50) * time.Millisecond)
        return 2
    })
    p.SetProgressSymbol("#")
    p.SetArrowSymbol("^")
    p.SetWaitSymbol("*")
    p.SetProgressColor(pb.TextGreen)
    p.SetArrowColor(pb.TextGreen)
    p.SetWaitColor(pb.TextRed)
    bar.AddProcess(p)

    bar.AddDefaultProcess("进程3", func(progress float64) int {
        // fmt.Println(progress)

        time.Sleep(time.Duration(200) * time.Millisecond)
        return 30
    })

    bar.Run()
}
