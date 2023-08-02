#!/bin/sh
arch=$(uname -m)
if echo $arch | grep -q 'arm'; then
    printf ' #!/bin/bash\n if [ ! -x "/usr/bin/qemu-x86_64" ];then\n sudo apt-get update\n sudo apt-get -y install make gcc g++ libglib2.0-dev libpixman-1-dev libfdt-dev python3-pip ninja-build\n sudo pip3 install meson\n wget https://download.qemu.org/qemu-6.2.0.tar.xz\n tar -xvf qemu-6.2.0.tar.xz\n cd qemu-6.2.0\n sudo ./configure\n sudo make -j 4\n sudo make install\n cd ..\n cp /usr/local/bin/qemu-x86_64  /usr/bin/qemu-x86_64\n fi\n' > qemu_install.sh
    chmod +x qemu_install.sh
    ./qemu_install.sh
    GOARCH=amd64 go test -c .
    qemu-x86_64 -cpu max ./sonic.test -test.v
fi