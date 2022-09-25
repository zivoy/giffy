$addons="C:\Users\Games\AppData\Roaming\Blender Foundation\Blender\2.93\scripts\addons"

if(($item = Get-Item -Path "$addons\Giffy" -ErrorAction SilentlyContinue)) {
Remove-Item $addons\Giffy -Recurse
}
Copy-Item Giffy $addons\ -Recurse