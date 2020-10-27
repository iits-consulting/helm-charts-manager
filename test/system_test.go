package test

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestLint(t *testing.T) {
	cmd := exec.Command("../helm-charts-manager", "lint", "--update", "--config-file", "./test-config.yaml", "--charts-path", "./")
	result := readByteArray(cmd.CombinedOutput())

	expectedOutput := strings.Split(readByteArray(ioutil.ReadFile("./expected-outputs/lint-expected-output")), "\n")
	for _, line := range expectedOutput {
		assert.Contains(t, result, line)
	}
}

func TestPlan(t *testing.T) {
	cmd := exec.Command("../helm-charts-manager", "plan", "--update", "--config-file", "./test-config.yaml", "--charts-path", "./")
	result := readByteArray(cmd.CombinedOutput())

	expectedOutput := strings.Split(readByteArray(ioutil.ReadFile("./expected-outputs/plan-expected-output")), "\n")
	for _, line := range expectedOutput {
		assert.Contains(t, result, line)
	}
}

func TestApply(t *testing.T) {
	cmd := exec.Command("../helm-charts-manager", "apply", "--auto-approve", "--update", "--config-file", "./test-config.yaml", "--charts-path", "./")
	result := readByteArray(cmd.CombinedOutput())

	expectedOutput := strings.Split(readByteArray(ioutil.ReadFile("./expected-outputs/apply-expected-output")), "\n")
	for _, line := range expectedOutput {
		assert.Contains(t, result, line)
	}
}

func readByteArray(byteArray []byte, err error) string {
	if err != nil {
		panic(err)
	}
	return string(byteArray)
}
