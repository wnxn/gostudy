SRC_DIR=.
DST_DIR=.

protoc:
	protoc -I=${SRC_DIR} --go_out=${DST_DIR} ${SRC_DIR}/encoding.proto

clean:
	rm -rf ${DST_DIR}/*
