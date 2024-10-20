hello:
	echo "Hello Make"

compile:
	echo "Compiling for Every Platform"	
	GOOS=linux GOARCH=amd64 go build -o ./build/bin/linux/assetLoader
	GOOS=windows GOARCH=amd64 go build -o ./build/bin/windows/assetLoader.exe
	GOOS=darwin GOARCH=amd64 go build -o ./build/bin/macOs/assetLoader
