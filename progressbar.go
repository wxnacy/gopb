package pb

import (
    "fmt"
    "strconv"
    "strings"
)

type ProgressBar struct {
    processes []*Process
    doneNum int
}

func New() *ProgressBar {
    return &ProgressBar{
        processes: make([]*Process, 0),
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


func (this *ProgressBar) Run() {

    // for i := 0; i < this.ProcessNum(); i++ {
        // fmt.Println(this.processes[i])
    // }

    Loop:
    for {
        for i := 0; i < this.ProcessNum(); i++ {
            prog := this.processes[i]
            if prog.IsDone() {
                this.doneNum++
            }
            prog.Print()
        }
        if this.doneNum == this.ProcessNum() {
            break Loop
        }
    }

}

type Process struct {
    index int
    position int
    total int
    current int
    begin int
    incr int
    prefix string
    width int
    todo func(progress float64) int
}

func NewProcess(
    total, begin int,
    prefix string,
    todo func(progress float64) int,
) *Process {
    p := &Process{
        width: 50,
        total: total,
        begin: begin,
        current: begin,
        prefix: prefix,
        todo: todo,
    }
    return p
}

func (this *Process) Progress() float64 {
    return float64(this.current) / float64(this.total)
}

func (this *Process) progressNum() int {
    i, _ := strconv.Atoi(
        fmt.Sprintf("%.0f", this.Progress() * float64(this.width)),
    )
    return i
}

func (this *Process) IsDone() bool {
    return this.current >= this.total
}

func (this *Process) increase() {
    this.current += this.incr
    if this.current > this.total {
        this.current = this.total
    }
}

func (this *Process) IsBegin() bool {
    return this.current > this.begin
}

func (this *Process) Print() {

    if this.IsBegin() && this.index == 0{
        fmt.Printf("\033[%dA\033[K", this.position)
    }

    output := fmt.Sprintf(
        "%s %d/%d [%s%s] %.2f%%",
        this.prefix,
        this.current, this.total,
        strings.Repeat("=", this.progressNum()),
        strings.Repeat("-", this.width - this.progressNum()),
        this.Progress() * 100,
    )

    fmt.Printf("%s \033[K\n", output)

    this.incr = this.todo(this.Progress())
    this.increase()

}
