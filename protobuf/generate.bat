protoc.exe -I .\ --go_out=.\ .\simple\simple.proto
protoc.exe -I .\ --go_out=.\ .\enum_example\enum_example.proto
protoc.exe -I .\ --go_out=.\ .\complex\complex.proto
protoc.exe .\greet\greet.proto --go_out=.\ --go-grpc_out=.