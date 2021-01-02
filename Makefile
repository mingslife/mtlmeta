build: bin/mtlmeta_linux_x64 bin/mtlmeta_linux_x86 bin/mtlmeta_linux_arm bin/mtlmeta_darwin_x64 bin/mtlmeta_windows_x64.exe bin/mtlmeta_windows_x86.exe

bin/mtlmeta_linux_x64:
	GOOS=linux GOARCH=amd64 go build -o bin/mtlmeta_linux_x64 cmd/*.go

bin/mtlmeta_linux_x86:
	GOOS=linux GOARCH=386 go build -o bin/mtlmeta_linux_x86 cmd/*.go

bin/mtlmeta_linux_arm:
	GOOS=linux GOARCH=arm go build -o bin/mtlmeta_linux_arm cmd/*.go

bin/mtlmeta_darwin_x64:
	GOOS=darwin GOARCH=amd64 go build -o bin/mtlmeta_darwin_x64 cmd/*.go

bin/mtlmeta_windows_x64.exe:
	GOOS=windows GOARCH=amd64 go build -o bin/mtlmeta_windows_x64.exe cmd/*.go

bin/mtlmeta_windows_x86.exe:
	GOOS=windows GOARCH=386 go build -o bin/mtlmeta_windows_x86.exe cmd/*.go
