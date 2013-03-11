package gomounts

import (
	"testing"
)

func TestGetMountedVolumes(t *testing.T) {
	volumes, err := GetMountedVolumes()
	if err != nil {
		t.Fatal(err)
	}
	for _, volume := range volumes {
		t.Log(volume.Path, volume.Type)
	}
	if len(volumes) == 0 {
		t.Fatal("No volumes found")
	}
}
