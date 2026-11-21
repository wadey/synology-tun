//go:build linux

package main

import (
	"strings"
	"testing"
)

func TestIsModuleLoaded(t *testing.T) {
	tests := []struct {
		name    string
		content string
		module  string
		want    bool
	}{
		{
			name:    "present first line",
			content: "tun 57344 2 - Live 0xffffffffc0123000\next4 786432 1 - Live 0xffffffffc0456000\n",
			module:  "tun",
			want:    true,
		},
		{
			name:    "present middle line",
			content: "ext4 786432 1 - Live 0xffffffffc0456000\ntun 57344 0 - Live 0xffffffffc0123000\n",
			module:  "tun",
			want:    true,
		},
		{
			name:    "not present",
			content: "ext4 786432 1 - Live 0xffffffffc0456000\n",
			module:  "tun",
			want:    false,
		},
		{
			name:    "prefix mismatch",
			content: "tunnel4 16384 0 - Live 0xffffffffc0123000\ntunnel6 16384 0 - Live 0xffffffffc0456000\n",
			module:  "tun",
			want:    false,
		},
		{
			name:    "empty",
			content: "",
			module:  "tun",
			want:    false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := isModuleLoaded(strings.NewReader(tc.content), tc.module)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Fatalf("isModuleLoaded(%q) = %v, want %v", tc.module, got, tc.want)
			}
		})
	}
}
