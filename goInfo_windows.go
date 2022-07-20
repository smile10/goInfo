package goInfo

import (
	"bytes"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func GetInfo() (GoInfoObject, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	pn, _, err := k.GetStringValue("ProductName")
	if err != nil {
		gio := GoInfoObject{Kernel: "windows", Core: "unknown", Platform: runtime.GOARCH, OS: "windows", GoOS: runtime.GOOS, CPUs: runtime.NumCPU()}
		gio.Hostname, _ = os.Hostname()
		return gio, fmt.Errorf("getInfo: %s", err)
	}

	cb, _, err := k.GetStringValue("CurrentBuild")
	if err != nil {
		gio := GoInfoObject{Kernel: "windows", Core: "unknown", Platform: runtime.GOARCH, OS: "windows", GoOS: runtime.GOOS, CPUs: runtime.NumCPU()}
		gio.Hostname, _ = os.Hostname()
		return gio, fmt.Errorf("getInfo: %s", err)
	}

	gio := GoInfoObject{Kernel: "windows", Core: cb, Platform: runtime.GOARCH, OS: pn, GoOS: runtime.GOOS, CPUs: runtime.NumCPU()}
	gio.Hostname, _ = os.Hostname()
	return gio, nil
}
