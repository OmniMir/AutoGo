;;ENVIRONMENT
#NoEnv  ; Recommended for performance and compatibility with future AutoHotkey releases.
;#Warn  ; Enable warnings to assist with detecting common errors.
SendMode Input  ; Recommended for new scripts due to its superior speed and reliability.
SetWorkingDir %A_ScriptDir%  ; Ensures a consistent starting directory.


;;VARIABLES
WinWidth := A_ScreenWidth
WinHeight := (A_ScreenHeight - 40)/2


;;HOTKEYS
;Comics Increment (Ctrl + Win + V)
^#v::
Run .\AutoGo.exe --comics
Send {Blind}{Ctrl up}
return

;Comics Increment Folder (Ctrl + Win + C)
^#c::
Send {F2}
Copy()
Send {Escape}
Run .\AutoGo.exe --comics
Send {Blind}{Ctrl up}
return

;Clipboard History (Ctrl + Win + X)
^#x::
Send #v
Send {Blind}{Ctrl up}
return

;Chrome to Opera (RightWin + O)
#>o::
Send !d
Copy()
Send ^w
Run .\AutoGo.exe --opera
return

;Search in Google (RightWin + ?)
#>sc035::
Copy()
Run .\AutoGo.exe --google
return

;Arrow Plus (Ctrl + "-")
^-::
Send {+}
Send {>}
return

;Arrow Equal (Ctrl + "=")
^=::
Send {=}
Send {>}
return

;Accent mark (Win + "`")
#sc029::
Send {U+0301}
return


;;CALL TO APPS
;Notepad (Ctrl + Alt + Б)
^!sc033::
Run %A_WinDir%\System32\notepad.exe
return

;Passwords (Ctrl + Alt + P)
^!p:: 
Run %A_ProgramFiles%\KeePassXC\KeePassXC.exe
return

;WinRAR (Ctrl + Alt + R)
^!r::
Run %A_ProgramFiles%\WinRAR\WinRAR.exe
return


;;WINDOW MANAGMENT
;Window to Top (Ctrl + Space)
Ctrl & Space::
WinSet, AlwaysOnTop, Toggle, A ; A is Active Window
return

;Window to Up (RightAlt + Up)
RAlt & Up::
WinRestore A
WinMove A, , 0, 0, WinWidth, WinHeight
return

;Window to Down (RightAlt + Down)
RAlt & Down::
WinRestore A
WinMove A, , 0, (WinHeight + 1), WinWidth, WinHeight
return

;Window Max when Win+Up not working (RightAlt + RightWin)
RAlt & RWin::
WinMaximize A
return


;;FUNCTIONS
;Copy text to Clipboard
Copy(){ 
  Sleep 100
  Clipboard := ""
  Send ^c
  ClipWait, 2 ;Timeout of 2 seconds
  return
}
