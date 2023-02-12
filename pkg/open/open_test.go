package open

import "testing"

func TestOpenFolder(t *testing.T) {
	if err := OpenFolder("C:/Users/ethan/Documents/Strathclyde"); err != nil {
		t.Error(err)
	}
}
