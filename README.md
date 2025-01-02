# README

Fuzzy snippet manager.

```
.\anchoco.exe -h
Usage of .\anchoco.exe:
  -src string
        src yaml path
```

Specify a yaml file for `-src` in the format of [prompts.yaml](prompts.yaml).


## With PowerShell

```powershell
function Invoke-Anchoco {
    $gotool = $env:USERPROFILE | Join-Path -ChildPath "Personal\tools\bin\anchoco.exe"
    if (-not (Test-Path $gotool)) {
        "Not found: {0}" -f $gotool | Write-Host -ForegroundColor Magenta
        $repo = "https://github.com/AWtnb/anchoco"
        "=> Clone and build from {0}" -f $repo | Write-Host
        return
    }
    $result = & $gotool ('--src={0}' -f ($env:USERPROFILE | Join-Path -ChildPath "Personal\tools\prompts.yaml"))
    if ($LASTEXITCODE -ne 0) {
        return
    }
    $result | Set-Clipboard
    Start-Process "https://claude.ai/"
}
```