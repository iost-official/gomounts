package gomounts_test

import (
	"github.com/strib/gomounts"
	"log"
)

func Example() {
	volumes, err := gomounts.GetMountedVolumes()
	if err == nil {
		for _, v := range volumes {
			// Windows might print "Volume of type NTFS is mounted at C:\"
			// Unix might print "Volume of type ext3 is mounted at /media/HD1"
			log.Printf("Volume of type %s is mounted at %s", v.Type, v.Path)
		}
	}
}
