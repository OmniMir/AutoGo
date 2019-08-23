;ENVIRONMENT
#NoEnv  ; Recommended for performance and compatibility with future AutoHotkey releases.
; #Warn  ; Enable warnings to assist with detecting common errors.
SendMode Input  ; Recommended for new scripts due to its superior speed and reliability.
SetWorkingDir %A_ScriptDir%  ; Ensures a consistent starting directory.

;VARIABLES
WinWidth := A_ScreenWidth
WinHeight := (A_ScreenHeight - 40)/2

;HOTKEYS
^#v:: ;Comics Increment
Run .\AutoGo.exe --comics
return

^#c:: ;Comics Increment Folder
Send {F2}
Copy()
Send {Escape}
Run .\AutoGo.exe --comics
return

#>o:: ;Chrome to Opera
Send !d
Copy()
Send ^w
Run .\AutoGo.exe --opera
return

#g:: ;Search in Google
Copy()
Run .\AutoGo.exe --google
return

;WINDOW MANAGMENT
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
Copy(){ ;Copy text to Clipboard
  Sleep 100
  Clipboard := ""
  Send ^c
  ClipWait
  return
}
