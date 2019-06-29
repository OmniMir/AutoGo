#NoEnv  ; Recommended for performance and compatibility with future AutoHotkey releases.
; #Warn  ; Enable warnings to assist with detecting common errors.
SendMode Input  ; Recommended for new scripts due to its superior speed and reliability.
SetWorkingDir %A_ScriptDir%  ; Ensures a consistent starting directory.

;HOTKEYS
~^#v::
return

~^#c::
Copy()
return

~^!o::
Copy()
return

~^#g::
Copy()
return

;FUNCTIONS
Copy(){
  Sleep 100
  Clipboard := ""
  Send ^c
  ClipWait
}
