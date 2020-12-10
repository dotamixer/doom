#!/bin/bash

if [[ $(OS), Windows_NT ]];then
    bin=greeter-client.exe
else
    bin=greeter-client
fi

cd bin && ./${bin}
cd -