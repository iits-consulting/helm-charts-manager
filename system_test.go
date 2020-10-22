package main

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
	cmd := exec.Command("./helm-charts-manager", "lint", "--update", "--config-file", "./test/test-config-apply.yaml", "--charts-path", "./test/")
	result := readByteArray(cmd.CombinedOutput())

	expectedOutput := strings.Split(readByteArray(ioutil.ReadFile("./test/lint-expected-output.txt")), "\n")
	for _, line := range expectedOutput {
		assert.Contains(t, result, line)
	}
}

func TestPlan(t *testing.T) {
	cmd := exec.Command("./helm-charts-manager", "plan", "--update", "--config-file", "./test/test-config-plan.yaml", "--charts-path", "./test/")
	result := readByteArray(cmd.CombinedOutput())

	expectedOutput := strings.Split(readByteArray(ioutil.ReadFile("./test/plan-expected-output.txt")), "\n")
	for _, line := range expectedOutput {
		assert.Contains(t, result, line)
	}
}

func TestApply(t *testing.T) {
	cmd := exec.Command("./helm-charts-manager", "apply", "--auto-approve", "--update", "--config-file", "./test/test-config-apply.yaml", "--charts-path", "./test/")
	result := readByteArray(cmd.CombinedOutput())

	expectedOutput := strings.Split(readByteArray(ioutil.ReadFile("./test/apply-expected-output.txt")), "\n")
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
