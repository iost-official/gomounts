// +build windows

package gomounts

/*
#include <string.h>
#include <stdlib.h>
// Dev stub
int GetLogicalDrives(void)
{
	return 13; // A, C, D
}
// Dev stub
int GetVolumeInformation(
char* lpRootPathName,
char* lpVolumeNameBuffer,
int nVolumeNameSize,
int* lpVolumeSerialNumber,
int* lpMaximumComponentLength,
int* lpFileSystemFlags,
char* lpFileSystemNameBuffer,
int nFileSystemNameSize
)
{
	strncpy(lpFileSystemNameBuffer, "NTFS", nFileSystemNameSize);
	return 1; // Success
}
*/
import "C"

import (
	"unsafe"
)

// Windows implementation
func getMountedVolumes() ([]Volume, error) {
	result := make([]Volume, 0)
	var buf [256]C.char

	drives := uint32(C.GetLogicalDrives())

	for i := uint32(0); i < 26; i++ {
		if (1<<i)&drives != 0 {
			letter := 'A' + i
			rootPath := string(letter) + `:\`
			fsType := func() string {
				cRootPath := C.CString(rootPath)
				defer C.free(unsafe.Pointer(cRootPath))
				if C.GetVolumeInformation(cRootPath, nil, 0, nil, nil, nil, &buf[0], C.int(len(buf))) != 0 {
					return C.GoString(&buf[0])
				}
				return "Unknown"
			}()
			result = append(result, Volume{rootPath, fsType})
		}
	}

	return result, nil
}
