BUILD_DIR=./build

clean:
	rm -rf ${BUILD_DIR}

build:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64       go build -o build/score-windows-amd64.exe score.go
	CGO_ENABLED=0 GOOS=windows GOARCH=386         go build -o build/score-windows-386.exe   score.go
	CGO_ENABLED=0 GOOS=linux   GOARCH=amd64       go build -o build/score-linux-amd64       score.go
	CGO_ENABLED=0 GOOS=linux   GOARCH=386         go build -o build/score-linux-386         score.go
	CGO_ENABLED=0 GOOS=linux   GOARCH=arm64       go build -o build/score-linux-arm64       score.go
	CGO_ENABLED=0 GOOS=linux   GOARCH=arm GOARM=7 go build -o build/score-linux-arm-7       score.go
	CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64       go build -o build/score-darwin-amd64      score.go
	CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64       go build -o build/score-darwin-arm64      score.go
	CGO_ENABLED=0 GOOS=android GOARCH=arm64       go build -o build/score-andriod-arm64    score.go
