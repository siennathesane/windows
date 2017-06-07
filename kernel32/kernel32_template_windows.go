package kernel32

//sysdoc	GetCurrentProcess retrieves a pseudo handle for the current process. See: https://msdn.microsoft.com/en-us/library/windows/desktop/ms683179%28v=vs.85%29.aspx
//sys	GetCurrentProcess() (pseudoHandle windows.Handle, err error)

//sysdoc	QueryProcessCycleTime retrieves the sum of the cycle time of all threads of the specified process.
//sys	QueryProcessCycleTime(handle windows.Handle, cycleTime *windows.PULong64) (err error)

//sysdoc	GetLastError retrieves the calling thread's last-error code value. The last-error code is maintained on a per-thread basis. Multiple threads do not overwrite each other's last-error code.
//sys	GetLastError() (err error)


//sys	getComputerName(lpBuffer *windows.LptStr, lpnSize *windows.LpdWord) (err error) = GetComputerNameW