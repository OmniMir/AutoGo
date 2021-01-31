#Notification Toast for Windows 10
Function Invoke-Notification {
	param (
		[string] $Title,
		[string] $Text
	)

	#Notification Template
	[Windows.UI.Notifications.ToastNotificationManager, Windows.UI.Notifications, ContentType = WindowsRuntime] > $null
	$template = [Windows.UI.Notifications.ToastNotificationManager]::GetTemplateContent([Windows.UI.Notifications.ToastTemplateType]::ToastText01)

	#Notification Convert to .NET Type
	$RawXml = [xml] $template.GetXml()
	$RawXml.GetElementsByTagName("text").AppendChild($RawXml.CreateTextNode($Text)) > $null
	#Notification Convert Back to XML
	$xml = New-Object Windows.Data.Xml.Dom.XmlDocument
	$xml.LoadXml($RawXml.OuterXml)

	#Notification Settings
	$notification = [Windows.UI.Notifications.ToastNotification]::new($xml)
	$notification.Tag = "PowerShell"
	$notification.Group = "PowerShell"
	#$notification.ExpirationTime = [DateTimeOffset]::Now.AddMinutes(5)
	#$notification.SuppressPopup = $true

	#Notification Toast
	$toast = [Windows.UI.Notifications.ToastNotificationManager]::CreateToastNotifier($Title)
	$toast.Show($notification);

}
