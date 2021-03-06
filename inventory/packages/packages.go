/*
Copyright 2017 Google Inc. All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package packages provides package management functions for Windows and Linux
// systems.
package packages

import (
	"io/ioutil"
	"log"
	"os/exec"
	"time"

	"github.com/GoogleCloudPlatform/osconfig/inventory/osinfo"
)

var (
	// AptExists indicates whether apt is installed.
	AptExists bool
	// YumExists indicates whether yum is installed.
	YumExists bool
	// ZypperExists indicates whether zypper is installed.
	ZypperExists bool
	// GemExists indicates whether gem is installed.
	GemExists bool
	// PipExists indicates whether pip is installed.
	PipExists bool
	// GooGetExists indicates whether googet is installed.
	GooGetExists bool

	noarch = osinfo.Architecture("noarch")

	// DebugLogger is the debug logger to use.
	DebugLogger = log.New(ioutil.Discard, "", 0)
)

// Packages is a selection of packages based on their manager.
type Packages struct {
	Yum           []PkgInfo     `json:"yum,omitempty"`
	Rpm           []PkgInfo     `json:"rpm,omitempty"`
	Apt           []PkgInfo     `json:"apt,omitempty"`
	Deb           []PkgInfo     `json:"deb,omitempty"`
	Zypper        []PkgInfo     `json:"zypper,omitempty"`
	ZypperPatches []ZypperPatch `json:"Zypper_patches,omitempty"`
	Gem           []PkgInfo     `json:"gem,omitempty"`
	Pip           []PkgInfo     `json:"pip,omitempty"`
	GooGet        []PkgInfo     `json:"googet,omitempty"`
	WUA           []WUAPackage  `json:"wua,omitempty"`
	QFE           []QFEPackage  `json:"qfe,omitempty"`
}

// PkgInfo describes a package.
type PkgInfo struct {
	Name, Arch, Version string
}

// ZypperPatch describes a Zypper patch.
type ZypperPatch struct {
	Name, Category, Severity, Summary string
}

// WUAPackage describes a Windows Update Agent package.
type WUAPackage struct {
	Title                    string
	Description              string
	Categories               []string
	CategoryIDs              []string
	KBArticleIDs             []string
	SupportURL               string
	UpdateID                 string
	RevisionNumber           int32
	LastDeploymentChangeTime time.Time
}

// QFEPackage describes a Windows Quick Fix Engineering package.
type QFEPackage struct {
	Caption, Description, HotFixID, InstalledOn string
}

var run = func(cmd *exec.Cmd) ([]byte, error) {
	DebugLogger.Printf("Running %q with args %q\n", cmd.Path, cmd.Args[1:])
	return cmd.CombinedOutput()
}
