#NoEnv  ; Recommended for performance and compatibility with future AutoHotkey releases.
; #Warn  ; Enable warnings to assist with detecting common errors.
SendMode Input  ; Recommended for new scripts due to its superior speed and reliability.
SetWorkingDir %A_ScriptDir%  ; Ensures a consistent starting directory.

WinWidth := A_ScreenWidth
WinHeight := (A_ScreenHeight - 40)/2

;HOTKEYS
~^#v::
Run .\AutoGo.exe --comics
return

~^#c::
Send {F2}
Copy()
Send {Escape}
Run .\AutoGo.exe --comics
return

~^!o::
Send {!d}
Copy()
Run .\AutoGo.exe --opera
return

~^#g::
Copy()
Run .\AutoGo.exe --google
return

RAlt & Up:: ;Window to Up
WinRestore A
WinMove A, , 0, 0, WinWidth, WinHeight
return

RAlt & Down:: ;Window to Down
WinRestore A
WinMove A, , 0, (WinHeight + 1), WinWidth, WinHeight
return

RAlt & RWin:: ;Window Max when Win+Up not working
WinMaximize A
return

;FUNCTIONS
Copy(){
  Sleep 100
  Clipboard := ""
  Send ^c
  ClipWait
}
