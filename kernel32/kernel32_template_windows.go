package kernel32

//sysdoc	GetCurrentProcess retrieves a pseudo handle for the current process. See: https://msdn.microsoft.com/en-us/library/windows/desktop/ms683179%28v=vs.85%29.aspx
//sys	GetCurrentProcess() (pseudoHandle windows.Handle, err error)

//sysdoc	QueryProcessCycleTime retrieves the sum of the cycle time of all threads of the specified process.
//sys	QueryProcessCycleTime(handle windows.Handle, cycleTime *windows.PULong64) (err error)

//sysdoc	GetLastError retrieves the calling thread's last-error code value. The last-error code is maintained on a per-thread basis. Multiple threads do not overwrite each other's last-error code.
//sys	GetLastError() (err error)

//sysdoc	CreatePseudoConsole creates a new pseudoconsole object for the calling process. See: https://docs.microsoft.com/en-us/windows/console/createpseudoconsole
//sys	CreatePseudoConsole(size windows.COORD, hInput windows.Handle, hOutput windows.Handle, dwFlags windows.Dword, phPC *windows.HpCon) (err error) = CreatePseudoConsole

//sysdoc	ResizePseudoConsole resizes the internal buffers for a pseudoconsole to the given size. See: https://docs.microsoft.com/en-us/windows/console/resizepseudoconsole
//sys	ResizePseudoConsole(hPC windows.HpCon, size windows.COORD) (err error) = ResizePseudoConsole

//sysdoc	ClosePseudoConsole closes a pseudoconsole from the given handle. See: https://docs.microsoft.com/en-us/windows/console/closepseudoconsole
//sys	ClosePseudoConsole(hPC windows.HpCon) (err error) = ClosePseudoConsole

//sys	getComputerName(lpBuffer *windows.LptStr, lpnSize *windows.LpdWord) (err error) = GetComputerNameW