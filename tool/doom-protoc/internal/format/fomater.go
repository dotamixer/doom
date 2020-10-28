package format

import (
	"log"
	"os/exec"
)

var (
	protoStyle = `{
Language: Proto,
BasedOnStyle: Google,
AlignConsecutiveAssignments: true,
AlignConsecutiveDeclarations: true,
}`
)

type Formatter struct {
}

func NewFormatter() *Formatter {
	return &Formatter{}
}

func (f *Formatter) Format(absFile []string) {
	args := []string{"-style", protoStyle, "-i"}
	args = append(args, absFile...)
	cmd := exec.Command("clang-format", args...)
	log.Println(cmd.String())
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
		log.Fatal("compile Error:", err)
	}
}
