// +build darwin

package gomounts

/*
#include <sys/param.h>
#include <sys/ucred.h>
#include <sys/mount.h>
*/
import "C"
import (
	"errors"
	"reflect"
	"fmt"
	"strconv"
	"sync"
	"unsafe"
)

var mtx sync.Mutex = sync.Mutex{}

func getMountedVolumes() ([]Volume, error) {
	// getmntinfo is non-reentrant
	mtx.Lock()
	defer mtx.Unlock()

	result := make([]Volume, 0)

	var mntbuf *C.struct_statfs
	cnt, err := C.getmntinfo(&mntbuf, C.MNT_NOWAIT)
	if err != nil {
		return result, fmt.Errorf("Failure getmntinfo: %+v", err)
	}
	count := int(cnt)
	if count == 0 {
		return nil, nil
	}

	// Convert to go slice per https://code.google.com/p/go-wiki/wiki/cgo
	var mntSlice []C.struct_statfs
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&mntSlice))
	sliceHeader.Cap = count
	sliceHeader.Len = count
	sliceHeader.Data = uintptr(unsafe.Pointer(mntbuf))

	for _, v := range mntSlice {
		uidstr := strconv.Itoa(int(v.f_owner))
		result = append(result, Volume{C.GoString(&v.f_mntonname[0]), C.GoString(&v.f_fstypename[0]), uidstr})
	}

	return result, nil
}
