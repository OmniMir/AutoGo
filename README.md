# AutoGo

Automation in Go



**Build**

git clone https://github.com/OmniMir/AutoGo.git

cd AutoGo

go build -ldflags="-H windowsgui -s -w"

Ahk2Exe.exe /in AutoGo.ahk /out AutoGoAHK.exe



_ldflags:_

"-H windowsgui": set executable GUI binary on Windows

''-s": strip debug information

''-w": strip symbol tables



**Installation**

Copy bin files (or make them) to your folder.

Make Scheduled Tasks and Autorun to some of them.

Rejoice



**Using**