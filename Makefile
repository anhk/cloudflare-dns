export GOPROXY=https://goproxy.cn,direct
export GO111MODULE=on

OBJ = cloudflare-dns

default: $(OBJ)

$(OBJ):
	go build -gcflags "-N -l" -o $@ ./src

clean:
	rm -fr $(OBJ)

-include .deps

dep:
	echo -n "$(OBJ):" > .deps
	find . -name '*.go' | awk '{print $$0 " \\"}' >> .deps
	echo "" >> .deps
