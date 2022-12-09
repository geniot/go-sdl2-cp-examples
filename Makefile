PROGRAM_NAME := go-sdl2-cp-examples

all: clean build

clean:
	rm bin/${PROGRAM_NAME}* -f

build:
	go build -o bin/${PROGRAM_NAME} geniot.com/geniot/go-sdl2-cp-examples/cmd/go-sdl2-cp-examples

mips:
	CC='/opt/gcw0-toolchain/usr/bin/mipsel-gcw0-linux-uclibc-gcc' \
	CGO_CFLAGS='-I/opt/gcw0-toolchain/usr/mipsel-gcw0-linux-uclibc/sysroot/usr/include -I/opt/gcw0-toolchain/usr/mipsel-gcw0-linux-uclibc/sysroot/usr/include/libpng16 -D_REENTRANT' \
	 CGO_ENABLED=1 \
	 CGO_LDFLAGS='-L/opt/gcw0-toolchain/usr/mipsel-gcw0-linux-uclibc/sysroot/usr/lib -lSDL2 -lpng16' \
	 GOARCH=mipsle \
	 GOMIPS=softfloat \
	 GOOS=linux \
	 PKG_CONFIG='/opt/gcw0-toolchain/usr/bin/pkg-config' \
	 go build -o bin/sdl_physics.gcw geniot.com/geniot/sdl_physics_go/cmd/sdl_physics

squash:
	mksquashfs bin/go-sdl2-cp-examples.gcw resources/media/go-sdl2-cp-examples.png resources/default.gcw0.desktop bin/go-sdl2-cp-examples.opk -all-root -no-xattrs -noappend -no-exports

opk: clean mips squash

#on PG2 use opkrun


