package util

import "testing"

func TestAttempt3(t *testing.T) {
	type args struct {
		prompt  string
		attempt func() bool
	}

	attemptCount := 0
	tests := []struct {
		name        string
		args        args
		wantSuccess bool
	}{
		// TODO: Add test cases.
		{
			name: "Positive 1st attempt",
			args: args{
				prompt:  "Hello World",
				attempt: func() bool { return true }},
			wantSuccess: true,
		},
		{
			name: "NegativeTest",
			args: args{
				prompt:  "Hello World",
				attempt: func() bool { return false }},
			wantSuccess: false,
		},
		{
			name: "Positive 2nd attempt",
			args: args{
				prompt: "Hello World",
				attempt: func() (ret bool) {
					if attemptCount == 1 {
						ret = true
					} else {
						attemptCount++
						ret = false
					}
					return
				}},
			wantSuccess: true,
		}, {
			name: "Positive 3rd attempt",
			args: args{
				prompt: "Hello World",
				attempt: func() (ret bool) {
					if attemptCount == 2 {
						ret = true
					} else {
						attemptCount++
						ret = false
					}
					return
				}},
			wantSuccess: true,
		}, {
			name: "Positive 4th attempt",
			args: args{
				prompt: "Hello World",
				attempt: func() (ret bool) {
					if attemptCount == 3 {
						ret = true
					} else {
						attemptCount++
						ret = false
					}
					return
				}},
			wantSuccess: false,
		},
	}
	for _, tt := range tests {
		attemptCount = 0 // reset test variable
		t.Run(tt.name, func(t *testing.T) {
			if gotSuccess := Attempt3(tt.args.prompt, tt.args.attempt); gotSuccess != tt.wantSuccess {
				t.Errorf("Attempt3() = %v, want %v", gotSuccess, tt.wantSuccess)
			}
		})
	}
}
