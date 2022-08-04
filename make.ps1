# Windows PowerShell Script
# Usage: .\make.ps1 init, .\make.ps1 config, .\make.ps1 api, ...
param ($stage)

$directories = gci app -Directory -d 0
foreach ($dir in $directories) {
	cd $dir.FullName
	pwd
	.\make.ps1 $stage
}
cd $PSScriptRoot