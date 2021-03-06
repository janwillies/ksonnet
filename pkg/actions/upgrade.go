// Copyright 2018 The ksonnet authors
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package actions

import (
	"io"
	"os"

	"github.com/ksonnet/ksonnet/pkg/app"
	"github.com/ksonnet/ksonnet/pkg/registry"
	"github.com/ksonnet/ksonnet/pkg/upgrade"
)

// RunUpgrade runs `upgrade`.
func RunUpgrade(m map[string]interface{}) error {
	a, err := newUpgrade(m)
	if err != nil {
		return err
	}

	return a.run()
}

// Upgrade upgrades an application.
type Upgrade struct {
	app       app.App
	pm        registry.PackageManager
	upgradeFn func(a app.App, out io.Writer, pl upgrade.PackageLister, dryRun bool) error
	dryRun    bool
}

func newUpgrade(m map[string]interface{}) (*Upgrade, error) {
	ol := newOptionLoader(m)

	a := ol.LoadApp()
	if ol.err != nil {
		return nil, ol.err
	}
	pm := registry.NewPackageManager(a)

	u := &Upgrade{
		app:       a,
		pm:        pm,
		upgradeFn: upgrade.Upgrade,
		dryRun:    ol.LoadBool(OptionDryRun),
	}

	if ol.err != nil {
		return nil, ol.err
	}

	return u, nil
}

// Upgrade upgrades a ksonnet application.
func (u *Upgrade) run() error {
	return u.upgradeFn(u.app, os.Stdout, u.pm, u.dryRun)
}
