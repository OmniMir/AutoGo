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

;;DEFAULT STATE OF LOCKS (+Numlock -CapsLock -ScrollLock)
SetNumlockState AlwaysOn
SetCapsLockState AlwaysOff
SetScrollLockState AlwaysOff

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
^sc00C::
Send {+}
Send {>}
return

;Arrow Equal (Ctrl + "=")
^=::
Send {=}
Send {>}
return

;Clipboard Alphabet Sort (AltGr + Z)
<^>!z::
Run AutoGo --alphabet
Blinder()
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

;Current Date  (Ctrl + Alt + D)
^!d::
FormatTime CurrentTime,, dd.MM.yyyy ;10.01.2019
Send %CurrentTime%
return

;Current Time  (Shift + Ctrl + Alt + D)
+^!d::
FormatTime CurrentTime,, HH:mm ;15:25
Send %CurrentTime%
return

;Ellipsis (Alt + .)
!.::
Send {U+2026} ; Unicode Symbol …
return

;En Dash (Alt + -)
!sc00C::
Send {U+2013} ; Unicode Symbol –
return

;Em Dash (Shift + Alt + -)
+!sc00C::
Send {U+2014} ; Unicode Symbol —
return

;Em Space (Alt + Space)
!Space::
Send {U+2003} ; Unicode Symbol "	"
return

;Paste to Begining (AltGr + <)
<^>!sc033::
Send {Home}
Send ^v
Send {Down}
return

;Paste to Ending (AltGr + >)
<^>!sc034::
Send {End}
Send ^v
Send {Down}
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
;Comics Increment Folder (Ctrl + Win + C)
^#c::
if WinActive("ahk_exe Explorer.exe")
{
	Send {F2}
	Copy()
	Send {Escape}
	Run AutoGo --comics
	Blinder()
}
return

;Explorer to Double Commander (AltGr + D)
<^>!d::
if WinActive("ahk_exe Explorer.exe")
{
	Send ^l
	Copy()
	Run AutoGo --doublecmd
}
return

;Explorer to TagScanner (AltGr + M)
<^>!m::
if WinActive("ahk_exe Explorer.exe")
{
	Send ^l
	Copy()
	Run AutoGo --tagscanner
}
return

;Explorer to Windows Terminal (AltGr + T)
<^>!t::
if WinActive("ahk_exe Explorer.exe")
{
	Send ^l
	Copy()
	Run AutoGo --terminal
}
return

;Explorer to WinRAR (AltGr + R)
<^>!r::
if WinActive("ahk_exe Explorer.exe")
{
	Send ^l
	Copy()
	Run AutoGo --winrar
}
return


;;BROWSERS CONTROL (Chromium-based: Chrome, Edge, Opera and etc.)
;Chromium to Chrome (AltGr + C)
<^>!c::
if WinActive("ahk_class Chrome_WidgetWin_1")
{
	Send !d
	Copy()
	Send ^w
	Run AutoGo --chrome
}
return

;Chromium to Edge (AltGr + E)
<^>!e::
if WinActive("ahk_class Chrome_WidgetWin_1")
{
	Send !d
	Copy()
	Send ^w
	Run AutoGo --edge
}
return

;Chromium to Opera (AltGr + O)
<^>!o::
if WinActive("ahk_class Chrome_WidgetWin_1")
{
	Send !d
	Copy()
	Send ^w
	Run AutoGo --opera
}
return

;Next Tab (Ctrl + Mouse Forward Button)
<^XButton2::
^Tab
return

;Previous Tab (Ctrl + Mouse Previous Button)
<^XButton1::
+^Tab
return

;Open Closed Tab (Ctrl + Win + Z)
^#z::
if WinActive("ahk_class Chrome_WidgetWin_1")
{
	Send ^+t
}
return


;;CAPSLOCK CONTROL (Explorer, Trello, WinRAR)
CapsLock & w::
;Explorer Up to Folder
if WinActive("ahk_exe Explorer.exe")
{
	Send {Alt down}
	Send {Up}
	Send {Alt up}
	Send {Up}
	Send {Enter}
}
;Trello Up to Card
if WinActive("ahk_exe msedge.exe")
{
	Send k
}
;WinRAR Down to Archive
if WinActive("ahk_exe WinRAR.exe")
{
	Send {Alt down}
	Send {Up}
	Send {Alt up}
	Send {Up}
	Send {Enter}
}
return

CapsLock & s::
;Explorer Down to Folder
if WinActive("ahk_exe Explorer.exe")
{
	Send {Alt down}
	Send {Up}
	Send {Alt up}
	Send {Down}
	Send {Enter}
}
;Trello Down to Card
if WinActive("ahk_exe msedge.exe")
{
	Send j
}
;WinRAR Down to Archive
if WinActive("ahk_exe WinRAR.exe")
{
	Send {Alt down}
	Send {Up}
	Send {Alt up}
	Send {Down}
	Send {Enter}
}
return


;;MARKDOWN STYLE
;Markdown Asterisk (AltGr + 8)
*!>sc009::
Send {Text}\*
return

;Markdown Brackets (Ctrl + '(')
^sc00A::
Send {Text})_
Send {Home}
Send {Text} _(
Send {End}
Send {Left 2}
return

;Markdown Crossing (Ctrl + '~')
^sc029::
Send {Text}~~
Send {Home}
Send {Text}~~
Send {End}
Send {Left 2}
return

;Markdown Link (Ctrl + 8)
^sc009::
Send {Text}**
Send {Home}
Send {Text}**
Send {End}
Send {Left 2}
return

;Markdown Link (Ctrl + ')')
^sc00B::
Send {Text}]()
Send {Home}
Send {Text}[
Send {End}
Send {Left}
return


;;CALL TO APPS
;Chrome Update Service (Ctrl + Alt + U)
<^<!u::
Run .\ServiceChrome.exe
return

;Office Service (Ctrl + Alt + O)
<^<!o::
Run .\ServiceOffice.exe
return

;Passwords (Ctrl + Alt + P)
<^<!p::
Run %A_ProgramFiles%\KeePassXC\KeePassXC.exe
return

;WinRAR (Ctrl + Alt + R)
<^<!r::
Run %A_ProgramFiles%\WinRAR\WinRAR.exe
return

;Notepad (Ctrl + Alt + Б)
<^<!sc033::
Run %A_WinDir%\System32\notepad.exe
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

;Window to NotePad (AltGr + Right)
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
