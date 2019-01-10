package psapi

//sys	getProcessMemoryInfo(handle windows.Handle, memCounters *ProcessMemoryCounters, cb windows.Dword) (err error) = psapi.GetProcessMemoryInfo
//sys	getPerformanceInfo(perfInfo *PerformanceInformation, cb windows.Dword) (err error) = psapi.GetPerformanceInfo

//sysdoc	EnumProcesses retrieves the process identifier for each process object in the system.
//sys	EnumProcesses(lpidProcess *[]byte, cb windows.Dword, lpcbNeeded *windows.LpDword) (ok bool) = psapi.EnumProcesses

