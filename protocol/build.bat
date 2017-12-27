set PATH=d:\dev\protoc\bin;%PATH%

set SRC=(authentication keyexchange mercury metadata pubsub spirc)

for %%A in %SRC% DO mkdir %%A
for %%A in %SRC% DO protoc --go_out=%%A\ -I proto proto\%%A.proto

