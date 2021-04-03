lint:
	make -C go lint
	make -C hbase lint
	make -C java lint
	make -C kotlin lint
	make -C release lint
	make -C ruby lint
	make -C scala lint
	make -C diagram lint

build:
	make -C go build
	make -C hbase build
	make -C java build
	make -C kotlin build
	make -C release build
	make -C ruby build
	make -C scala build
	make -C diagram build

push:
	make -C go push
	make -C hbase push
	make -C java push
	make -C kotlin push
	make -C release push
	make -C ruby push
	make -C scala push
	make -C diagram push
