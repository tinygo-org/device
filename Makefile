gen-device: gen-device-stm32

build/gen-device-svd: ./tools/gen-device-svd/*.go
	go build -o $@ ./tools/gen-device-svd/

gen-device-stm32: build/gen-device-svd
	./build/gen-device-svd -source=https://github.com/tinygo-org/stm32-svd lib/stm32-svd/svd stm32
	GO111MODULE=off go fmt ./stm32
