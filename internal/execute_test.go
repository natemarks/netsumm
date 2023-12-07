package internal

import (
	"reflect"
	"testing"
)

func TestExecute(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name string
		args args
		want ExecuteResult
	}{
		{
			"ping - fail",
			args{"ping -c 1 -W 1"},
			ExecuteResult{
				Stdout:   "",
				Stderr:   "ping: usage error: Destination address required\n",
				ExitCode: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Execute(tt.args.command); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
