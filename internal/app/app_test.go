package app

import "testing"

func TestVersionedGreeting(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "", want: "teltonika-playground says hello, world"},
		{name: "codec", want: "teltonika-playground says hello, codec"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := VersionedGreeting(tc.name)
			if got != tc.want {
				t.Fatalf("VersionedGreeting(%q) = %q, want %q", tc.name, got, tc.want)
			}
		})
	}
}
