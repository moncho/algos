package main

import (
	"testing"
)

func Test_socialNet_AllFriends(t *testing.T) {
	type args struct {
		members     int
		connections [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"2 members, all friends",
			args{
				2,
				[][]int{[]int{0, 1}},
			},
			true,
		},
		{
			"3 members, not all friends",
			args{
				3,
				[][]int{[]int{0, 1}},
			},
			false,
		},
		{
			"3 members, all friends",
			args{
				3,
				[][]int{
					[]int{0, 1},
					[]int{1, 2}},
			},
			true,
		},
		{
			"5 members, all friends",
			args{
				5,
				[][]int{
					[]int{0, 1},
					[]int{1, 2},
					[]int{0, 2},
					[]int{2, 1},
					[]int{3, 4},
					[]int{4, 1},
					[]int{2, 4},
					[]int{4, 0},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sn := newSocialNetwork(tt.args.members)
			for _, c := range tt.args.connections {
				sn.NewFriends(c[0], c[1])
			}
			if got := sn.AllFriends(); got != tt.want {
				t.Errorf("socialNet.AllFriends() = %v, want %v", got, tt.want)
			}
		})
	}
}
