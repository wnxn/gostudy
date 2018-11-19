# Guide proto3

## Field Numbers

1~15 one byte to encode: F(field's type) F(field number)
16~2047 two bytes to encode: F(field's type) FFF(field number) 2^12 = 2048
19000~19999 reserved for the protobuf

## Nested Types

a nested can be reused outside its parent message type