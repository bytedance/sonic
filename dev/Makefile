

ARCHS   := linux_x86 linux_aarch64 darwin_x86 darwin_aarch64
TARGETS := x86_64-unknown-linux-gnu aarch64-unknown-linux-gnu x86_64-apple-darwin aarch64-apple-darwin

all: ${ARCHS}

rustup:
	rustup target add ${TARGETS}

linux_x86:
	mkdir -p ./rs_wrapper/lib/linux
	cd rs_wrapper && cargo update --package sonic-rs && RUSTFLAGS="-C target-cpu=haswell" cargo build --target=x86_64-unknown-linux-gnu  --release && cp -a target/x86_64-unknown-linux-gnu/release/librs_wrapper.a  ./lib/linux/libsonic_rs_x86_64-unknown-linux-gnu.a


linux_aarch64:
	mkdir -p ./rs_wrapper/lib/linux
	cd rs_wrapper && cargo update --package sonic-rs && RUSTFLAGS="-C target-feature=+neon" cargo build --target=aarch64-unknown-linux-gnu  --release && cp -a target/aarch64-unknown-linux-gnu/release/librs_wrapper.a  ./lib/linux/libsonic_rs_aarch64-unknown-linux-gnu.a

darwin_x86:
	mkdir -p ./rs_wrapper/lib/darwin
	cd rs_wrapper && cargo update --package sonic-rs && RUSTFLAGS="-C target-cpu=haswell" cargo build --target=x86_64-apple-darwin  --release && cp -a target/x86_64-apple-darwin/release/librs_wrapper.a  ./lib/darwin/libsonic_rs_x86_64-apple-darwin.a

darwin_aarch64:
	mkdir -p ./rs_wrapper/lib/darwin
	cd rs_wrapper && cargo update --package sonic-rs && RUSTFLAGS="-C target-feature=+neon" cargo build --target=aarch64-apple-darwin  --release && cp -a target/aarch64-apple-darwin/release/librs_wrapper.a  ./lib/darwin/libsonic_rs_aarch64-apple-darwin.a

clean:
	rm -vrf ./rs_wrapper/lib/
	cd rs_wrapper && cargo clean