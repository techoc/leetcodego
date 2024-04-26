package t1146

import "testing"

func TestSnapshotArray(t *testing.T) {
	tests := []struct {
		name     string
		function []string
		args     [][]int
		want     []int
	}{
		{
			name:     "Case 1",
			function: []string{"SnapshotArray", "set", "snap", "set", "get"},
			args:     [][]int{{3}, {0, 5}, {}, {0, 6}, {0, 0}},
			want:     []int{0, 0, 0, 0, 5},
		},
		{
			name:     "Case 2",
			function: []string{"SnapshotArray", "set", "snap", "set", "get", "snap", "get"},
			args:     [][]int{{3}, {0, 5}, {}, {0, 6}, {0, 0}, {}, {0, 1}},
			want:     []int{0, 0, 0, 0, 5, 1, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sa SnapshotArray
			for i, f := range tt.function {
				switch f {
				case "SnapshotArray":
					sa = Constructor(tt.args[i][0])
				case "set":
					sa.Set(tt.args[i][0], tt.args[i][1])
				case "get":
					if got := sa.Get(tt.args[i][0], tt.args[i][1]); got != tt.want[i] {
						t.Errorf("SnapshotArray.Get() = %v, want %v", got, tt.want[i])
					}
				case "snap":
					if got := sa.Snap(); got != tt.want[i] {
						t.Errorf("SnapshotArray.Snap() = %v, want %v", got, tt.want[i])
					}
				}
			}
		})
	}

}
