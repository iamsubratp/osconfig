package main

import (
	"fmt"
	"sync"

	"github.com/GoogleCloudPlatform/osconfig/e2etester/executor"
	"github.com/GoogleCloudPlatform/osconfig/e2etester/test_suite_base"
	"github.com/GoogleCloudPlatform/osconfig/integration_test/test_classes"
)

func main() {
	xcutor := executor.GetITestExecutor()
	xcutor.AddTestOptions(executor.NewTestOptions("zypper",
		[]string{"projects/suse-cloud/global/images/family/sles-12",
			"projects/suse-cloud/global/images/family/sles-15"},
		test_classes.ZypperTest{}))
	var wg sync.WaitGroup
	for _, t := range xcutor.GetTestOptions() {
		fmt.Printf("running test_suite: %s\n", t.Name)
		wg.Add(1)
		go test_suite_base.Start(t, &wg)
		fmt.Printf("finished running test_suite\n")
	}
	fmt.Printf("waiting for tests to complete\n")
	wg.Wait()
}
