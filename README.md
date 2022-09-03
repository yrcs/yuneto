# YunETO (云医通) 
The Cloud-Native Online Appointment Patient Registration System.

## ENV

OS: Linux (Recommend AnolisOS v8.6⁺), macOS, Windows 11.

Go: **Must ≧ v1.19⁺** (Because this project uses **Go Generics**).

Other major requirements: Kratos v2.5.0⁺, Gin v1.8.1⁺, GORM v1.23.8⁺, Mariadb v10.6.8⁺, Podman v3.4.2⁺, CRI-O v1.24.1⁺.

## Usage

Run this project:

```shell
cd yuneto
kratos run
```

API Document：http://localhost:8000/swagger/index.html

Use `make` (or `.\make.ps1` on Windows) to build project. Type `make help` (or `.\make.ps1 help` on Windows) to show usage details.

This project uses **Monorepo** for cross-project packages sharing. If you want to `make` multiple projects at a time, you should `cd yuneto` first, and then execute `make ${command}`. If you just want to `make` one project at a time, you need to enter the directory of this project first, and then execute `make ${command}`. eg: `cd yuneto/app/hospital` , `make api` .
