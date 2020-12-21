;;ENVIRONMENT
#NoEnv  ; Recommended for performance and compatibility with future AutoHotkey releases.
;#Warn  ; Enable warnings to assist with detecting common errors.
SendMode Input  ; Recommended for new scripts due to its superior speed and reliability.
SetWorkingDir %A_ScriptDir%  ; Ensures a consistent starting directory.

;;EXE VERSION
;@Ahk2Exe-SetName AutoGoAHK
;@Ahk2Exe-SetDescription AutoGo Companion
;@Ahk2Exe-SetCompanyName OmniMir
;@Ahk2Exe-SetVersion 1.0

;;USEFUL LINKS
;;All Hotkeys https://www.autohotkey.com/docs/Hotkeys.htm
;;List of Keys https://www.autohotkey.com/docs/KeyList.htm
;;WinMove Function https://www.autohotkey.com/docs/commands/WinMove.htm

;;VARIABLES
AutoGo := ".\AutoGo.exe"
WinWidth := A_ScreenWidth
WinHeight := (A_ScreenHeight - 40) ;Windows Taskbar Height = 40


;;HOTKEYS
;Accent Mark (Alt + "`")
!sc029::
Send {U+0301} ;Unicode Symbol  ́
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

;Clipboard History (Ctrl + Win + X)
^#x::
Send #v
Blinder()
return

;Comics Increment (Ctrl + Win + V)
^#v::
Run AutoGo --comics
Blinder()
return

;Ruble (Alt + H)
!h::
Send {U+20BD} ; Unicode Symbol ₽
return

;Search in Google (RightWin + ?)
#>sc035::
Copy()
Run AutoGo --google
return


;;WINDOWS EXPLORER CONTROL
If WinActive(ahk_exe Explorer.exe)
{
	;Comics Increment Folder (Ctrl + Win + C)
	^#c::
	Send {F2}
	Copy()
	Send {Escape}
	Run AutoGo --comics
	Blinder()
	return

	;Explorer to Double Commander (AltGr + C)
	*!>c::
	Send ^l
	Copy()
	Run AutoGo --doublecmd
	return

	;Explorer to TagScanner (AltGr + M)
	*!>m::
	Send ^l
	Copy()
	Run AutoGo --tagscanner
	return

	;Explorer to Windows Terminal (AltGr + T)
	*!>t::
	Send ^l
	Copy()
	Run AutoGo --terminal
	return

	;Explorer to WinRAR (AltGr + R)
	*!>r::
	Send ^l
	Copy()
	Run AutoGo --winrar
	return
}


;;BROWSERS CONTROL (Chromium-based: Chrome, Edge, Opera and etc.)
If WinActive(ahk_class Chrome_WidgetWin_1)
{
	;Chrome to Edge (AltGr + E)
	*!>e::
	Send !d
	Copy()
	Send ^w
	Run AutoGo --edge
	return

	;Chrome to Opera (AltGr + O)
	*!>o::
	Send !d
	Copy()
	Send ^w
	Run AutoGo --opera
	return

	#!^Left::
	Send ^+t
	return
}


;;MARKDOWN STYLE
;Markdown Asterisk (AltGr + "*")
*!>sc037::
Send {Text}\*
return

;Markdown Brackets (Ctrl + "(")
^sc00A::
Send {Text})_
Send {Home}
Send {Text} _(
Send {End}
Send {Left 2}
return

;Markdown Crossing (Ctrl + "~")
^sc029::
Send {Text}~~
Send {Home}
Send {Text}~~
Send {End}
Send {Left 2}
return

;Markdown Link (Ctrl + "*")
^sc037::
Send {Text}**
Send {Home}
Send {Text}**
Send {End}
Send {Left 2}
return

;Markdown Link (Ctrl + ")")
^sc00B::
Send {Text}]()
Send {Home}
Send {Text}[
Send {End}
Send {Left}
return


;;CALL TO APPS
;Notepad (Ctrl + Alt + Б)
^!sc033::
Run %A_WinDir%\System32\notepad.exe
;Chrome Update Service (Ctrl + Alt + U)
^!u::
Run .\ServiceChrome.exe
return

;Office Service (Ctrl + Alt + O)
^!o::
Run .\ServiceOffice.exe
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
;Window Max when Win+Up not working (AltGr + RightWin)
RAlt & RWin::
WinMaximize A
return

;Window to Down (AltGr + Down)
RAlt & Down::
WinRestore A
WinMove A, , 0, (WinHeight/2 + 1), WinWidth, (WinHeight/2)
return

;Window to Top (Ctrl + Space)
Ctrl & Space::
WinSet, AlwaysOnTop, Toggle, A ; A is Active Window
return

;Window to NotePad (Ctrl + AltGr + Down)
RAlt & Right::
indent := 300
WinRestore A
WinMove A, , (indent/2), (WinHeight/2 - indent), (WinWidth - indent), (WinHeight/2 + indent)
return

;Window to Up (AltGr + Up)
RAlt & Up::
WinRestore A
WinMove A, , 0, 0, WinWidth, (WinHeight/2)
return


;;FUNCTIONS
;Hotkeys can be glued, When they cross. Therefore make them Up via Blind
Blinder(){
	Send {Blind}{Ctrl up}
	Send {Blind}{Shift up}
	Send {Blind}{Alt up}
	Send {Blind}{Win up}
}

;Copy text to Clipboard
Copy(){
	Sleep 100
	Clipboard := ""
	Send ^c
	ClipWait, 2 ;Timeout of 2 seconds
	return
}
