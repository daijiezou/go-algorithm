package main

import "testing"

func TestGetOne(t *testing.T) {
	type args struct {
		str1 string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				str1: "asdfasdfo",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetOne(tt.args.str1)
		})
	}
}

func Test_checkPwd(t *testing.T) {
	type args struct {
		str1 string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				str1: "Q~7T&(4^$OXz725(3!Xl(3+s*",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkPwd(tt.args.str1)
		})
	}
}
