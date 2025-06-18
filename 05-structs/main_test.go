package main

import "testing"

func Test_getUserData(t *testing.T) {
	type args struct {
		promptText string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "",
			args: args{
				promptText: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUserData(tt.args.promptText); got != tt.want {
				t.Errorf("getUserData() = %v, want %v", got, tt.want)
			}
		})
	}
}
