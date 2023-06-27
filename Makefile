
.PHONY: image
image:
	bazel build :tarball
	docker load --input $(shell bazel cquery --output=files :tarball)
	docker run centos5test:latest