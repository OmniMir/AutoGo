package main

const (
	app        = "!App"
	appsFolder = "shell:AppsFolder\\"
	calculator = "Microsoft.WindowsCalculator_8wekyb3d8bbwe"
	store      = "Microsoft.WindowsStore_8wekyb3d8bbwe"
	terminal   = "Microsoft.WindowsTerminal_8wekyb3d8bbwe"
)

func Calculator() {
	StartExplorer(appsFolder + calculator + app)
}
func Store() {
	StartExplorer(appsFolder + store + app)
}
func Terminal() {
	StartExplorer(appsFolder + terminal + app)
}
