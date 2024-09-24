package is_subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSubSequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "simple true",
			args: args{t: "ahbgdc", s: "abc"},
			want: true,
		},
		{
			name: "simple false",
			args: args{t: "ahbgdc", s: "axc"},
			want: false,
		},
		{
			name: "simple false",
			args: args{t: "ahbgdc", s: ""},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSubSequence(tt.args.s, tt.args.t)
			assert.Equal(t, tt.want, got)
		})
	}
}
