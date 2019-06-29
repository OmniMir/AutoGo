#NoEnv  ; Recommended for performance and compatibility with future AutoHotkey releases.
; #Warn  ; Enable warnings to assist with detecting common errors.
SendMode Input  ; Recommended for new scripts due to its superior speed and reliability.
SetWorkingDir %A_ScriptDir%  ; Ensures a consistent starting directory.

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

;FUNCTIONS
Copy(){
  Sleep 100
  Clipboard := ""
  Send ^c
  ClipWait
}
