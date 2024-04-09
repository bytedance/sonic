

ARCHS   := linux_x86 linux_aarch64 darwin_x86 darwin_aarch64
TARGETS := x86_64-unknown-linux-gnu aarch64-unknown-linux-gnu x86_64-apple-darwin aarch64-apple-darwin

all: ${ARCHS}

rustup:
	rustup target add ${TARGETS}

linux_x86:
	mkdir -p ./rs_wrapper/lib/linux
	cd rs_wrapper && cargo update --package sonic-rs && RUSTFLAGS="-C target-cpu=haswell" cargo build --target=x86_64-unknown-linux-gnu  --release && cp -a target/x86_64-unknown-linux-gnu/release/librs_wrapper.so  ../internal/link/libsonic_rs_x86_64-unknown-linux-gnu.so 

linux_aarch64:
	mkdir -p ./rs_wrapper/lib/linux
	cd rs_wrapper && cargo update --package sonic-rs && RUSTFLAGS="-C target-feature=+neon" cargo build --target=aarch64-unknown-linux-gnu  --release && cp -a target/aarch64-unknown-linux-gnu/release/librs_wrapper.so  ../internal/link/libsonic_rs_aarch64-unknown-linux-gnu.so 

darwin_x86:
	mkdir -p ./rs_wrapper/lib/darwin
	cd rs_wrapper && cargo update --package sonic-rs && RUSTFLAGS="-C target-cpu=haswell" cargo build --target=x86_64-apple-darwin  --release && cp -a target/x86_64-unknown-linux-gnu/release/librs_wrapper.so  ../internal/link/libsonic_rs_x86_64-unknown-darwin-gnu.so 

darwin_aarch64:
	mkdir -p ./rs_wrapper/lib/darwin
	cd rs_wrapper && cargo update --package sonic-rs && RUSTFLAGS="-C target-feature=+neon" cargo build --target=aarch64-apple-darwin  --release && cp -a target/aarch64-apple-darwin/release/librs_wrapper.dylib  ../internal/link/libsonic_rs_aarch64-apple-darwin.dylib

clean:
	rm -vrf ./rs_wrapper/lib/
	cd rs_wrapper && cargo clean