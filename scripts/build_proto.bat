del /q /f "..\internal\pb\user\*.*"

"..\tools\proto\protoc.exe" --plugin=protoc-gen-NAME=..\tools\proto\protoc-gen-gogofast.exe --NAME_out=..\internal\pb\user -I="..\proto\user" ..\proto\user\*.proto
pause
