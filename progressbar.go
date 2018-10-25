package pb


import (
    // "fmt"
    // "strconv"
    // "strings"
    "time"
    "fmt"
)


type ProgressBar struct {
    processes []*Process
    doneNum int
    done chan bool
}

func New() *ProgressBar {
    return &ProgressBar{
        processes: make([]*Process, 0),
        done: make(chan bool, 1),
    }
}

func (this *ProgressBar) Add(
    total,
    begin int,
    prefix string,
    todo func(progress float64) int,
) {
    p := NewProcess(total, begin, prefix, todo)
    this.AddProcess(p)
}

func (this *ProgressBar) AddDefaultProcess(
    prefix string,
    todo func(progress float64) int,
) {
    p := NewProcess(defaultTotal, defaultBegin, prefix, todo)
    this.AddProcess(p)
}

func (this *ProgressBar) AddProcess(p *Process) {

    this.processes = append(this.processes, p)
    for i := 0; i < this.ProcessNum(); i++ {
        this.processes[i].position++
        this.processes[i].index = i
    }
}

func (this *ProgressBar) ProcessNum() int {
    return len(this.processes)
}

func (this *ProgressBar) IsDone() bool {
    for i := 0; i < this.ProcessNum(); i++ {
        if ! this.processes[i].IsDone() {
            return false
        }
    }
    return true
}

func (this *ProgressBar) progressString() string {
    out := ""
    for i := 0; i < this.ProcessNum(); i++ {
        prog := this.processes[i]
        out += prog.toString()
        if i + 1 < this.ProcessNum() {
            out += "\n"
        }
    }
    return out
}

func (this *ProgressBar) Run() {

    // for i := 0; i < this.ProcessNum(); i++ {
        // fmt.Println(this.processes[i])
    // }
    out := this.progressString()
    fmt.Printf("%s \033[K\n", out)

    for i := 0; i < this.ProcessNum(); i++ {
        this.processes[i].run()
    }

    t := time.NewTicker(200 * time.Millisecond)

    Loop:
    for {
        select {
            case <- t.C: {

                fmt.Printf("\033[%dA\033[K", this.ProcessNum())
                out := this.progressString()
                fmt.Printf("%s \033[K\n", out)

                if this.IsDone() {
                    this.done <- true
                }

            }
            case <- this.done: {
                break Loop
            }
        }
    }

}

