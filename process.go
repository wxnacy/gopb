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
    progressSymbol string
    waitSymbol string
    arrowSymbol string
    progressColor int
    arrowColor int
    waitColor int
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
        progressSymbol: defaultProgressSymbol,
        arrowSymbol: defaultArrowSymbol,
        waitSymbol: defaultWaitSymbol,
        progressColor: defaultProgressColor,
        arrowColor: defaultArrowColor,
        waitColor: defaultWaitColor,
        done: make(chan bool, 0),
        todo: todo,
    }
    return p
}

func (this *Process) SetProgressSymbol(s string) {
    this.progressSymbol = s
}

func (this *Process) SetWaitSymbol(s string) {
    this.waitSymbol = s
}

func (this *Process) SetArrowSymbol(s string) {
    this.arrowSymbol = s
}

func (this *Process) SetProgressColor(c int) {
    this.progressColor = c
}

func (this *Process) SetWaitColor(c int) {
    this.waitColor = c
}

func (this *Process) SetArrowColor(c int) {
    this.arrowColor = c
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

func (this *Process) toString() string {
    output := fmt.Sprintf(
        "%s %s %s %s",
        this.prefix,
        this.numString(),
        this.progressString(),
        this.percentageString(),
    )
    return output
}

func (this *Process) progressString() string {

    pNum := this.progressNum()
    if pNum > 0 && !this.IsDone(){
        pNum--
    }
    progressStr := SetColor(
        strings.Repeat(this.progressSymbol, pNum),
        0, 0, this.progressColor,
    )

    arrowStr := SetColor(this.arrowSymbol, 0, 0, this.arrowColor)
    if this.IsDone() {
        arrowStr = ""
    }

    waitStr := SetColor(
        strings.Repeat(this.waitSymbol, this.width - this.progressNum()),
        0, 0, this.waitColor,
    )
    output := fmt.Sprintf(
        "%s%s%s%s%s",
        Blue("["),
        progressStr,
        arrowStr,
        waitStr,
        Blue("]"),
    )
    return output
}

func (this *Process) numString() string {
    totalLength := len(strconv.Itoa(this.total))

    totalStringFmt := fmt.Sprintf("%%%dd", totalLength)

    return fmt.Sprintf(
        totalStringFmt + "/" + totalStringFmt,
        this.current,
        this.total,
    )

}

func (this *Process) percentageString() string {
    s := fmt.Sprintf("%6.2f%%", this.Progress() * 100)

    if this.IsDone() {
        return Cyan(s)
    } else {
        return Yellow(s)
    }
}
