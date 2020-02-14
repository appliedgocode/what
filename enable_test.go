// +build what whathappens whatif whatfunc

package what

import (
	"testing"
)

func Test_pkgname(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"what",
			"what",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkgname(1); got != tt.want {
				t.Errorf("pkgname() = %v, want %v", got, tt.want)
			}
		})
	}
}
