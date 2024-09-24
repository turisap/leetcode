package ransomnote

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type args struct {
	ransomNote string
	magazine   string
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{
		name: "no",
		args: args{
			ransomNote: "a",
			magazine:   "b",
		},
		want: false,
	},
	{
		name: "no#2",
		args: args{
			ransomNote: "aa",
			magazine:   "ab",
		},
		want: false,
	},
	{
		name: "yes",
		args: args{
			ransomNote: "aa",
			magazine:   "aab",
		},
		want: true,
	},
	{
		name: "yes#2",
		args: args{
			ransomNote: "axz",
			magazine:   "aabcdgefteriajlkdnkdfghhkncbljklshouqwkrjknbknxjfkfjsmlkdfjlkdjhgksdsdlfjsfkjlsdfjxksjdlfjdslkfjz",
		},
		want: true,
	},
}

func TestCanConstruct(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canConstructNaive(tt.args.ransomNote, tt.args.magazine)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCanConstructOneCount(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canConstructOneCount(tt.args.ransomNote, tt.args.magazine)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCanConstructImproved(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canConstructNaiveImproved(tt.args.ransomNote, tt.args.magazine)
			assert.Equal(t, tt.want, got)
		})
	}
}
