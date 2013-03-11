// +build linux darwin

package gomounts

/*
#include <stdio.h>
#include <stdlib.h>
#include <mntent.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

// Unix implementation
func getMountedVolumes() ([]Volume, error) {
	result := make([]Volume, 0)

	cpath := C.CString("/proc/mounts")
	defer C.free(unsafe.Pointer(cpath))
	cmode := C.CString("r")
	defer C.free(unsafe.Pointer(cmode))
	var file *C.FILE = C.setmntent(cpath, cmode)
	if file == nil {
		return nil, errors.New("Unable to open /proc/mounts")
	}
	defer C.endmntent(file)
	var ent *C.struct_mntent

	for ent = C.getmntent(file); ent != nil; ent = C.getmntent(file) {
		mntType := C.GoString(ent.mnt_type)
		result = append(result, Volume{C.GoString(ent.mnt_dir), mntType})
	}

	return result, nil
}
