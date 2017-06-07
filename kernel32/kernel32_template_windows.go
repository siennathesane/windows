package kernel32

//sysdoc	GetCurrentProcess retrieves a pseudo handle for the current process. See: https://msdn.microsoft.com/en-us/library/windows/desktop/ms683179%28v=vs.85%29.aspx
//sys	GetCurrentProcess() (pseudoHandle windows.Handle, err error)
//sysdoc	QueryProcessCycleTime retrieves the sum of the cycle time of all threads of the specified process.
//sys	QueryProcessCycleTime(handle windows.Handle, cycleTime *windows.PULong64) (err error)
//sysdoc	GetLastError retrieves the calling thread's last-error code value. The last-error code is maintained on a per-thread basis. Multiple threads do not overwrite each other's last-error code.
//sys	GetLastError() (err error)
//sysdoc	Loads the specified module into the address space of the calling process. The specified module may cause other modules to be loaded.
//sys	LoadLibrary(lpFileName string) (handle windows.Handle, err error) = LoadLibraryW
//sysdoc	GetProcAddress retrieves the address of an exported function or variable from the specified dynamic-link library (DLL). See: https://msdn.microsoft.com/en-us/library/windows/desktop/ms683212(v=vs.85).aspx
//sys	GetProcAddress(hModule windows.Handle, lpProcName *byte) (addr windows.SizeT, err error)
//sysdoc	GetVersion returns the OS version. See: https://msdn.microsoft.com/en-us/library/windows/desktop/ms724439(v=vs.85).aspx
//sys	GetVersion() (ver windows.Dword, err error)


//sys	getComputerName(lpBuffer *windows.LptStr, lpnSize *windows.LpdWord) (err error) = GetComputerNameW
//sys	getUserName(lpbuffer *windows.LptStr, lpnSize *windows.LpdWord) (err error) = GetUserNameW
