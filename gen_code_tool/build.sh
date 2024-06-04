export CGO_ENABLED=0
export GOOS=windows
export GOARCH=amd64
go build -o qc_gen.exe
mv qc_gen.exe D:\\Code\\Go\\bin\\qc_gen.exe