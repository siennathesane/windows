package kernel32

//sysdoc	GetCurrentProcess retrieves a pseudo handle for the current process. See: https://msdn.microsoft.com/en-us/library/windows/desktop/ms683179%28v=vs.85%29.aspx
//sys	GetCurrentProcess() (pseudoHandle windows.HANDLE, err error)

//sysdoc	QueryProcessCycleTime retrieves the sum of the cycle time of all threads of the specified process.
//sys	QueryProcessCycleTime(handle windows.HANDLE, cycleTime *windows.PULONG64) (err error)
