// +build what whathappens whatif whatfunc

package what

import (
	"fmt"
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
