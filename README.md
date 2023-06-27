# rules_go_centos5

Testing building go bins that will run on centos5 with bazel + rules_go

## Build & Test

```console
$ make image
bazel build :tarball
docker load --input bazel-out/darwin_arm64-fastbuild/bin/tarball/tarball.tar
Loaded image: centos5test:latest
docker run centos5test:latest
Hello, linux_amd64!
uname="Linux ddf8930a7842 5.15.49-linuxkit-pr #1 SMP PREEMPT Thu May 25 07:27:39 UTC 2023 x86_64 x86_64 x86_64 GNU/Linux"
go="go1.20.5 X:nocoverageredesign"
```
