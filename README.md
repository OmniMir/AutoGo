# AutoGo

Automation in Go



#### Build

git clone https://github.com/OmniMir/AutoGo.git

cd AutoGo

go build -ldflags="-H windowsgui -s -w"

ahk2exe.exe /in AutoGo.ahk /out AutoGoAHK.exe

ps2exe.ps1 .\ServiceChrome.ps1 -noConsole

ps2exe.ps1 .\ServiceOffice.ps1 -noConsole



_ldflags:_

"-H windowsgui": set executable GUI binary on Windows

''-s": strip debug information

''-w": strip symbol tables



#### Installation

Copy bin files (or make them) to your folder.

Make Scheduled Tasks and Autorun to some of them.

Rejoice



#### Using