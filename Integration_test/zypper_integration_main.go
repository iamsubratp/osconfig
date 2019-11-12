package main

import (
	"github.com/GoogleCloudPlatform/compute-image-tools/go/packages"
	"github.com/GoogleCloudPlatform/osconfig/e2etester"
)

type ZypperWrapper struct {
}

func (* ZypperWrapper) ListZypperPatches() {
	packages.ZypperUpdates()
}

func main() {
	zypper := &ZypperWrapper{}
	e2etester.RunTests(zypper)
}
