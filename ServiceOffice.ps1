#Office ClickToRun Service
$officeService = "ClickToRunSvc"
#Notification
Import-Module (".\ServiceNotification.psm1")
$notificationTitle = "Service Office"
$notificationText = ""

#Service ON/OFF
if ((Get-Service ClickToRunSvc).StartType -eq "Disabled") {
	Set-Service -Name $officeService -StartupType Manual
	#Notification
	$notificationText = "ON"
}
else {
	Set-Service -Name $officeService -StartupType Disabled
	#Notification
	$notificationText = "OFF"
}

#Notification Toast
Invoke-Notification $notificationTitle $notificationText
