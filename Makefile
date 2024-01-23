release:
	# linux
	GOOS=linux go build -o dw ;\
	tar -zcvf ./releases/dw_linux.tar.gz ./dw ;\

	# macos
	GOOS=darwin go build -o dw ;\
	tar -zcvf ./releases/dw_macOS.tar.gz ./dw ;\

	# windows
	GOOS=windows go build -o dw ;\
	tar -zcvf ./releases/dw_windows.tar.gz ./dw ;\

	rm ./dw