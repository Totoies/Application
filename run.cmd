@echo off
echo %1

if "%1" equ "-build" (
    echo "Running the build"
    .\bin\tortoies.exe
)else (
    cd %~dp0\src
    go run .
    cd ..
)