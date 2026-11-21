//go:build linux

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

const (
	moduleName      = "tun"
	modulePath      = "/lib/modules/tun.ko"
	procModulesPath = "/proc/modules"
)

func isModuleLoaded(r io.Reader, name string) (bool, error) {
	prefix := name + " "
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), prefix) {
			return true, nil
		}
	}
	return false, scanner.Err()
}

func isLoaded(path, name string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()
	return isModuleLoaded(f, name)
}

func loadModule(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = unix.FinitModule(int(f.Fd()), "", 0)
	if errors.Is(err, unix.ENOSYS) {
		data, rerr := io.ReadAll(f)
		if rerr != nil {
			return rerr
		}
		err = unix.InitModule(data, "")
	}
	if errors.Is(err, unix.EEXIST) {
		return nil
	}
	return err
}

func run() error {
	loaded, err := isLoaded(procModulesPath, moduleName)
	if err != nil {
		return fmt.Errorf("check %s: %w", procModulesPath, err)
	}
	if loaded {
		return nil
	}

	if err := loadModule(modulePath); err != nil {
		return fmt.Errorf("load %s: %w", modulePath, err)
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
