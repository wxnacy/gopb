package pb

import (
    "fmt"
    "strconv"
    "strings"
)

type Process struct {
    index int
    position int
    total int
    current int
    begin int
    incr int
    prefix string
    width int
    done chan bool
    todo func(progress float64) int
}

func NewProcess(
    total, begin int,
    prefix string,
    todo func(progress float64) int,
) *Process {
    p := &Process{
        width: defaultWidth,
        total: total,
        begin: begin,
        current: begin,
        prefix: prefix,
        done: make(chan bool, 0),
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

func (this *Process) run() {
    // this.printProgress()
    go func(){
        Loop:
        for {
            this.increase()
            if this.IsDone() {
                break Loop
            }
        }
    }()
}

func (this *Process) increase() {
    this.incr = this.todo(this.Progress())
    this.current += this.incr
    if this.current > this.total {
        this.current = this.total
    }
}

func (this *Process) IsBegin() bool {
    return this.current > this.begin
}

func (this *Process) printProgress() {

    if this.IsBegin() && this.index == 0{
        fmt.Printf("\033[%dA\033[K", this.position)
    }
    output := this.progressString()

    fmt.Printf("%s \033[K\n", output)
}

func (this *Process) progressString() string {
    output := fmt.Sprintf(
        "%s %d/%d %s%s%s%s %s",
        this.prefix,
        this.current, this.total,
        Blue("["),
        Cyan(strings.Repeat("=", this.progressNum())),
        Yellow(strings.Repeat("-", this.width - this.progressNum())),
        Blue("]"),
        this.percentageString(),
    )
    return output
}

func (this *Process) percentageString() string {
    s := fmt.Sprintf("%6.2f%%", this.Progress() * 100)

    if this.IsDone() {
        return Cyan(s)
    } else {
        return Yellow(s)
    }
}
