package first_substr

import "testing"

func Test_firstSbuStrNaive(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple found",
			args: args{haystack: "sadbutsad", needle: "sad"},
			want: 0,
		},
		{
			name: "simple not found",
			args: args{haystack: "leetcode", needle: "leeto"},
			want: -1,
		},
		{
			name: "one letter",
			args: args{haystack: "aaa", needle: "aaaaa"},
			want: -1,
		},
		{
			name: "missisipy",
			args: args{haystack: "mississippi", needle: "issip"},
			want: 4,
		},
		{
			name: "just a",
			args: args{haystack: "a", needle: "a"},
			want: 0,
		},
		{
			name: "abc",
			args: args{haystack: "abc", needle: "c"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstSbuStrNaive(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("firstSbuStrNaive() = %v, want %v", got, tt.want)
			}
		})
	}
}
