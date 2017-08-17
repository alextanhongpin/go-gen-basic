go run main.go


Override individual string values:

+ `go run -ldflags '-X main.VersionString=1.0' main.go`
+ `go build -ldflags '-X main.VersionString=1.0' main.go`
+ `go run -ldflags "-X 'main.VersionString=`date`'" main.go`
+ `go generate`