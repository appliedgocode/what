// +build what whathappens whatif whatfunc

package what

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
)

func Test_pkgname(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"what",
			"appliedgo.net/what",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ""
			pkg, prnt := pkgname(1)
			if len(pkg) > 0 {
				got = fmt.Sprintf("%s/%s", prnt, pkg)
			} else {
				got = pkg
			}
			if got != tt.want {
				t.Errorf("pkgname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getenvWhat(t *testing.T) {
	tests := []struct {
		name     string
		packages string
		want     map[string]bool
	}{
		{
			"Single package",
			"pkg1",
			map[string]bool{
				"pkg1": true,
			},
		},
		{
			"Two packages",
			"pkg1,pkg2",
			map[string]bool{
				"pkg1": true,
				"pkg2": true,
			},
		},
		{
			"Two qualified paths",
			"path/to/my/pkg1,repohost.com/user/pkg2",
			map[string]bool{
				"path/to/my/pkg1":        true,
				"repohost.com/user/pkg2": true,
			},
		},
	}
	if len(os.Getenv("WHAT")) > 0 {
		t.Skip("Someone is using WHAT already, skipping")
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("WHAT", tt.packages)
			getenvWhat()
			if !reflect.DeepEqual(enabled, tt.want) {
				t.Errorf("Error reading WHAT. Got: %v, want: %v", enabled, tt.want)
			}
		})
	}
}
