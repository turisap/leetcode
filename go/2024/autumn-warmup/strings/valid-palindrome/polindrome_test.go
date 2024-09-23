package main

import "testing"

type args struct {
	s string
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{
		name: "simple valid with upper and whitespace",
		args: args{s: "A man, a plan, a canal: Panama"},
		want: true,
	},
	{
		name: "short 0P",
		args: args{s: "0P"},
		want: false,
	},
	{
		name: "horror",
		args: args{s: "`l;`` 1o1 ??;l`"},
		want: true,
	},
	{
		name: "1000 long",
		args: args{s: "X;;];;&;2;H;;];J;;#;[;;6;Z;<;N;_;>; > _ N < Z 6  [ #  J ]  H 2 &  ]  X"},
		want: true,
	},
}

func TestPalindromeTwoPointer(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeTwoPointer(tt.args.s); got != tt.want {
				t.Errorf("%s palindromeTwoPoint() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestPalindromeReverse(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeReverse(tt.args.s); got != tt.want {
				t.Errorf("%s palindromeReverse() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
