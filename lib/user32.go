package lib

import (
	"syscall"
	"unsafe"
)

const (
	MB_OK          = 0x00000000
	MB_OKCANCEL    = 0x00000001
	MB_YESNO       = 0x00000004
	MB_YESNOCANCEL = 0x00000003

	MB_APPLMODAL   = 0x00000000
	MB_SYSTEMMODAL = 0x00001000
	MB_TASKMODAL   = 0x00002000

	MB_ICONSTOP        = 0x00000010
	MB_ICONQUESTION    = 0x00000020
	MB_ICONWARNING     = 0x00000030
	MB_ICONINFORMATION = 0x00000040
)

// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw#return-value
const (
	IDCANCEL = 2
	IDYES    = 6
	IDNO     = 7
)

func User32Init() *syscall.LazyProc {
	var dll = syscall.NewLazyDLL("user32.dll")
	return dll.NewProc("MessageBoxW")
}

func MessageBox(hook *syscall.LazyProc, hwnd uintptr, text, caption string, flags uint) int {
	lpCaption, _ := syscall.UTF16PtrFromString(caption)
	lpText, _ := syscall.UTF16PtrFromString(text)

	ret, _, _ := syscall.SyscallN(
		hook.Addr(),
		hwnd,
		uintptr(unsafe.Pointer(lpText)),
		uintptr(unsafe.Pointer(lpCaption)),
		uintptr(flags),
	)
	return int(ret)
}
