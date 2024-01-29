

avx2:
	mkdir -p ./rs_wrapper/lib/avx2/
	cd rs_wrapper && cargo update --package sonic-rs && RUSTFLAGS="-C target-cpu=haswell" cargo build --release && cp -a target/release/librs_wrapper.a ./lib/avx2/

neon:
	mkdir -p ./rs_wrapper/lib/neon/
	cd rs_wrapper && cargo update --package sonic-rs && RUSTFLAGS="-C target-feature=+neon" cargo build --release && cp -a target/release/librs_wrapper.a ./lib/neon/

clean:
	rm -vrf ./rs_wrapper/lib/
	cd rs_wrapper && cargo clean