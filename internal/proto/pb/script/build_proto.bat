set homeford=%cd%

::user目录
rmdir /Q /S "%homeford%/../../user"
protoc -I=../user/ --go_out=../../  ../user/user.proto

echo "Complete"
