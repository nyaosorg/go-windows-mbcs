package mbcs_test

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/nyaosorg/go-windows-mbcs"
)

func TestConsoleCP(t *testing.T) {
	actual := mbcs.ConsoleCP()

	outputBytes, err := exec.Command("chcp").Output()
	if err != nil {
		t.Fatalf("chcp: %s", err.Error())
		return
	}
	outputString, err := mbcs.AnsiToUtf8(outputBytes, actual)
	if err != nil {
		t.Fatalf("mbcs.AtoU: %s", err.Error())
		return
	}
	fields := strings.Fields(outputString)

	if fmt.Sprintf("%d", actual) != fields[len(fields)-1] {
		t.Fatalf("ConsoleCP()==%d, chcp->%s", actual, outputString)
	}
}
