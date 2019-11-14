package executor

import "fmt"

type TestOptions struct {
	Name       string
	images     []string
	testHandle interface{}
}

func NewTestOptions(name string, i []string, th interface{}) *TestOptions {
	return &TestOptions{
		Name:       name,
		images:     i,
		testHandle: th,
	}
}

func (t *TestOptions) GetImages() []string {
	return t.images
}

func (t *TestOptions) GetTestHandle() interface{} {
	return t.testHandle
}

type ITestExecutor struct {
	testSets []*TestOptions
}

var executor *ITestExecutor

func GetITestExecutor() *ITestExecutor {
	if executor == nil {
		tOptions := []*TestOptions{}
		executor = &ITestExecutor{testSets: tOptions}
	}
	return executor
}

func (e *ITestExecutor) AddTestOptions(t *TestOptions) error {
	if t == nil {
		return fmt.Errorf("cannot add empty option")
	}
	e.testSets = append(e.testSets, t)
	return nil
}

func (e *ITestExecutor) GetTestOptions() []*TestOptions {
	return e.testSets
}
