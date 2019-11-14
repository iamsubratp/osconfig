package test_classes

import (
	"fmt"

	customError "github.com/GoogleCloudPlatform/osconfig/e2etester/error"
	"github.com/GoogleCloudPlatform/osconfig/inventory/packages"
)

type ZypperTest struct {
}

func (z ZypperTest) TestZypperCommands() error {

	ec := new(customError.ErrorCollector)

	// list installed packages
	pkgs, err := packages.ZypperInstalledPatches()
	if err != nil || len(pkgs) == 0 {
		ec.Collect(fmt.Errorf("Error listing patches: %+v\n", err))
	}

	// remove a package which is already installed
	err = packages.RemoveZypperPackages([]string{pkgs[0].Name})
	if err != nil {
		ec.Collect(fmt.Errorf("Error removing package: %+v", err))
	}

	// re-install the same package that we just removed
	err = packages.InstallZypperPackages([]string{pkgs[0].Name})
	if err != nil {
		ec.Collect(fmt.Errorf("error installing package: +%v", err))
	}

	count, errs := ec.Error()
	if count == 0 {
		return nil
	}
	return errs
}
