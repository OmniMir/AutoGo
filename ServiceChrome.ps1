#Chrome Update Service
$chromeUpdateService1 = "edgeupdate"
$chromeUpdateService2 = "edgeupdatem"
#Edge Update Service
$edgeUpdateService1 = "gupdate"
$edgeUpdateService2 = "gupdatem"
#Notification
Import-Module (".\ServiceNotification.psm1")
$notificationTitle = "Service Chrome"
$notificationText = ""

#Service ON/OFF
if ((Get-Service gupdate).StartType -eq "Disabled") {
	#Chrome Update Service
	Set-Service -Name $chromeUpdateService1 -StartupType Manual
	Set-Service -Name $chromeUpdateService2 -StartupType Manual
	#Edge Update Service
	Set-Service -Name $edgeUpdateService1 -StartupType Manual
	Set-Service -Name $edgeUpdateService2 -StartupType Manual
	#Notification
	$notificationText = "ON"
}
else {
	#Chrome Update Service
	Set-Service -Name $chromeUpdateService1 -StartupType Disabled
	Set-Service -Name $chromeUpdateService2 -StartupType Disabled
	#Edge Update Service
	Set-Service -Name $edgeUpdateService1 -StartupType Disabled
	Set-Service -Name $edgeUpdateService2 -StartupType Disabled
	#Notification
	$notificationText = "OFF"
}

#Notification Toast
Invoke-Notification $notificationTitle $notificationText
