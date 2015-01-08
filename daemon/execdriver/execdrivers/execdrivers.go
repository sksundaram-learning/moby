package execdrivers

import (
	"fmt"
	"path"

	"github.com/docker/docker/daemon/execdriver"
	"github.com/docker/docker/daemon/execdriver/lxc"
	"github.com/docker/docker/daemon/execdriver/native"
	"github.com/docker/docker/pkg/sysinfo"
	"github.com/docker/docker/pkg/system"
)

func NewDriver(name, root, initPath string, sysInfo *sysinfo.SysInfo) (execdriver.Driver, error) {
	meminfo, err := system.ReadMemInfo()
	if err != nil {
		return nil, err
	}

	switch name {
	case "lxc":
		// we want to give the lxc driver the full docker root because it needs
		// to access and write config and template files in /var/lib/docker/containers/*
		// to be backwards compatible
		return lxc.NewDriver(root, initPath, sysInfo.AppArmor)
	case "native":
		return native.NewDriver(path.Join(root, "execdriver", "native"), initPath, meminfo.MemTotal/1000)
	}
	return nil, fmt.Errorf("unknown exec driver %s", name)
}
