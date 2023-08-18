set homeford=%cd%

::user目录
rmdir /Q /S "%homeford%/../pb"

C:\Users\sunbin01\protoc\bin\protoc -I=../user/ --go_out=../  ../user/*.proto

echo "Complete"
