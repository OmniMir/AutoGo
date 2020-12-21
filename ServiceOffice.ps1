#Office ClickToRun Service
$officeService = "ClickToRunSvc"

if ((Get-Service ClickToRunSvc).StartType -eq "Disabled") {
	Set-Service -Name $officeService -StartupType Manual
}
else {
	Set-Service -Name $officeService -StartupType Disabled
}
