package file

import "testing"

func TestPrintFile(t *testing.T) {
	var wantErr bool

	if err := PrintFile("test.txt"); err != nil && wantErr {
		t.Errorf("PrintFile(): %v, want err: %v", err, wantErr)
	}
}
