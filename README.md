# YunETO (云医通) 
The Cloud-Native Online Appointment Patient Registration System.

## ENV

OS: Linux (Recommend AnolisOS v8.6⁺), macOS, Windows 11.

Go: **Must ≧ v1.19** (Because this project uses **Go Generics**).

Other major requirements: Protoc v21.5⁺, Kratos v2.5.0⁺, Gin v1.8.1⁺, GORM v1.23.8⁺, Mariadb v10.6.9⁺, Podman v4.1.1⁺, CRI-O v1.25.0⁺.

## Usage

Run this project:

```shell
cd yuneto
kratos run
```

API Document：http://localhost:8000/swagger/index.html

Deploy (Containerfile, sql files and ...) :  `yuneto/deploy`

Use `make` (or `.\make.ps1` on Windows) to build project. Type `make help` (or `.\make.ps1 help` on Windows) to show usage details.

This project uses **Monorepo** for cross-project packages sharing.

There are two ways to `make` projects:

1. `cd yuneto` , and then execute `make ${command}`;
2. Go to the specific sub-project, and then execute `make ${command}`. eg: `cd yuneto/app/hospital` , `make api` .
