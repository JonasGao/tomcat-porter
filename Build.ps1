$Version = 3
$BuildTime = Get-Date -Format yyyyMMdd-HHmmss
$GitHash = $(git rev-parse HEAD)
$env:GOOS = "linux"
go build -ldflags "-X main.Version=$Version -X main.BuildTime=$BuildTime -X main.GitHash=$GitHash"