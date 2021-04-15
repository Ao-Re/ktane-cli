package whosfirst

import "testing"

func TestWhosFirst_SolveWhosFirst(t *testing.T) {
	type args struct {
		ans string
	}
	tests := []struct {
		name string
		wf   *WhosFirst
		want string
	}{
		{
			name: "blank - middleRight - no",
			wf: &WhosFirst{
				DisplayText: "BLANK",
				LabelTexts:  []string{"SURE", "WAIT", "RIGHT", "WHAT?", "NO", "YES"},
			},
			want: "WAIT",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.wf.SolveWhosFirst(); got != tt.want {
				t.Errorf("WhosFirst.SolveWhosFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}
