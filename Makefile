
test:
	@echo " 这个程序不太乖 !!! "
	@echo " GitHub: https://github.com/getbuguai "
	@echo " BiliBili: https://space.bilibili.com/278413353 !!! "

build:
	go build
	echo "build ok !!!"

build_all:
	@GOOS=windows GOARCH=amd64 go build -o bg_jijian.exe
	@echo "windows ok !!!"
	@GOOS=linux GOARCH=amd64 go build -o bg_jijian
	@echo "linux ok !!!"
	@GOOS=darwin GOARCH=amd64 go build -o bg_jijian.app
	@echo "Mac ok !!!"

