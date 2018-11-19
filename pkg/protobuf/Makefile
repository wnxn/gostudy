SRC_DIR=.
DST_DIR=./tutorial

protoc:
	protoc -I=${SRC_DIR} --go_out=${DST_DIR} ${SRC_DIR}/addressbook.proto

clean:
	rm -rf ${DST_DIR}/*
