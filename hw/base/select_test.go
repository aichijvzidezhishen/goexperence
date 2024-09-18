package base

import (
	"testing"
)

func TestMultiChanChoose(t *testing.T) {
	// tests := []struct {
	// 	name string
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		MultiChanChoose()
	// 	})
	// }
	MultiChanChoose()

}

func TestTimeoutContro(t *testing.T) {
	// tests := []struct {
	// 	name string
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		TimeoutContro()
	// 	})
	// }
	TimeoutContro()

}

func TestOsExitNotify(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"t1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OsExitNotify()
		})
	}
}
