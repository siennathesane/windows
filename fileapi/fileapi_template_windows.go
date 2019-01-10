package fileapi

//sysdoc	ReadFile reads data from the specified file or input/output (I/O) device. Reads occur at the position specified by the file pointer if supported by the device. See: https://docs.microsoft.com/en-us/windows/desktop/api/fileapi/nf-fileapi-readfile
//sys	ReadFile(hFile, lpBuffer *windows.LpVoid, nNumberOfBytesToRead windows.Dword, lpNumberOfBytesRead *windows.LpDword, lpOverlapped *windows.Overlapped) (ok bool) = ReadFile
