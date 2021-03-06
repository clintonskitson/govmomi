/*
Copyright (c) 2015 VMware, Inc. All Rights Reserved.

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

package autostart

import (
	"flag"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/govc/flags"
	"github.com/vmware/govmomi/vim25/types"
)

type configure struct {
	*AutostartFlag

	types.AutoStartDefaults
}

func init() {
	cli.Register("host.autostart.configure", &configure{})
}

func (cmd *configure) Register(f *flag.FlagSet) {
	f.Var(flags.NewOptionalBool(&cmd.Enabled), "enabled", "")
	f.IntVar(&cmd.StartDelay, "start-delay", 0, "")
	f.StringVar(&cmd.StopAction, "stop-action", "", "")
	f.IntVar(&cmd.StopDelay, "stop-delay", 0, "")

	f.Var(flags.NewOptionalBool(&cmd.WaitForHeartbeat), "wait-for-heartbeat", "")
}

func (cmd *configure) Process() error { return nil }

func (cmd *configure) Usage() string {
	return ""
}

func (cmd *configure) Run(f *flag.FlagSet) error {
	return cmd.ReconfigureDefaults(cmd.AutoStartDefaults)
}
