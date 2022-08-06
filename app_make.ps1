# Windows PowerShell Script
# Do not use this script directly, it is called automatically by make.ps1
param ($stage)

# $GOPATH = go env GOPATH
# $KRATOS_VERSION_TMP = go mod graph | grep go-kratos/kratos/v2 | head -n 1
# $KRATOS_VERSION = $KRATOS_VERSION_TMP.Substring($KRATOS_VERSION_TMP.LastIndexOf('@') + 1)
# $KRATOS = $GOPATH + '\pkg\mod\github.com\go-kratos\kratos\v2@' + $KRATOS_VERSION
$VERSION = git describe --tags --always
$APP_RELATIVE_PATH = (gi $PWD).Parent.Parent.BaseName + '/' + (basename $PWD)
$APP_BASENAME = basename $APP_RELATIVE_PATH
$INTERNAL_PROTO_FILES = (gci internal -r -fi *.proto).FullName
$API_PROTO_FILES = (gci ..\..\api\$APP_BASENAME -r -fi *.proto).FullName
$APP_NAME = echo $APP_RELATIVE_PATH | sed -E 's|\/|-|g'
$CONTAINER_IMAGE = 'yrcs/{0}:0.1.0' -f (echo $APP_NAME)

switch ($stage) {
  # init env
  'init' {
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
    go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2
    go install github.com/envoyproxy/protoc-gen-validate@latest
    go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
    Break
  }

  # generate internal conf.pb.go
  'config' {
    protoc --proto_path=$PWD\internal `
      --proto_path=..\..\third_party `
      --go_out=paths=source_relative:.\internal `
      $INTERNAL_PROTO_FILES
    Break
  }

  # generate pagination.pb.go
  'page' {
    protoc --proto_path=$PSScriptRoot\third_party `
      --go_out=paths=source_relative:$PSScriptRoot\third_party `
      --validate_out=paths=source_relative,lang=go:$PSScriptRoot\third_party `
      $PSScriptRoot\third_party\pagination\pagination.proto
    Break
  }

  # generate api *.pb.go
  'api' {
    protoc --proto_path=$PSScriptRoot\api\$APP_BASENAME `
      --proto_path=$PSScriptRoot\third_party `
      --go_out=paths=source_relative:$PSScriptRoot\api\$APP_BASENAME `
      --go-errors_out=paths=source_relative:$PSScriptRoot\api\$APP_BASENAME `
      --validate_out=paths=source_relative,lang=go:$PSScriptRoot\api\$APP_BASENAME `
      --go-grpc_out=paths=source_relative:$PSScriptRoot\api\$APP_BASENAME `
      --go-http_out=paths=source_relative:$PSScriptRoot\api\$APP_BASENAME `
      $API_PROTO_FILES
    Break
  }

  # wire
  'wire' {
    go mod tidy
    go install github.com/google/wire/cmd/wire@latest
    cd cmd\$APP_BASENAME
    wire
    cd ..\..
    Break
  }

  # test
  'test' {
    go test -v ./... -cover
    Break
  }

  # build
  'build' {
    if (!(Test-Path -Path bin)) {
      mkdir bin
    }
    go build -ldflags "-X main.Name=$APP_NAME -X main.Version=$VERSION" -o ./bin/ ./...
    Write-Host "Project app/$APP_BASENAME has been successfully built."
    Write-Host "Use the following commands to start the project:"
    Write-Host "$ cd app\$APP_BASENAME"
    Write-Host "$ .\bin\$APP_BASENAME -conf ./configs"
    Break
  }

  # generate all
  'all' {
    .\make.ps1 config
    .\make.ps1 api
    .\make.ps1 wire
    .\make.ps1 test
    .\make.ps1 build
    Break
  }

  # build container image (for WSLv2 only)
  'container' {
    cd ../..
    podman build -f deploy/build/Containerfile --build-arg APP_BASENAME=$APP_BASENAME --rm -t $CONTAINER_IMAGE .
    Break
  }

  # show help
  'help' {
    Write-Host ""
    Write-Host "Usage:`r`n .\make.ps1 [target]"
    Write-Host ""
    Write-Host "Targets:"
    Write-Host "init`t`t`t`t" -ForegroundColor DarkGreen -NoNewline
    Write-Host "init env"
    Write-Host "config`t`t`t`t" -ForegroundColor DarkGreen -NoNewline
    Write-Host "generate internal conf.pb.go"
    Write-Host "page`t`t`t`t" -ForegroundColor DarkGreen -NoNewline
    Write-Host "generate pagination.pb.go"
    Write-Host "api`t`t`t`t" -ForegroundColor DarkGreen -NoNewline
    Write-Host "generate api *.pb.go"
    Write-Host "wire`t`t`t`t" -ForegroundColor DarkGreen -NoNewline
    Write-Host "generate wire_gen.go"
    Write-Host "test`t`t`t`t" -ForegroundColor DarkGreen -NoNewline
    Write-Host "go test"
    Write-Host "build`t`t`t`t" -ForegroundColor DarkGreen -NoNewline
    Write-Host "go build"
    Write-Host "all`t`t`t`t" -ForegroundColor DarkGreen -NoNewline
    Write-Host "generate all pb.go, wire, test and build"
    Write-Host "container`t`t`t" -ForegroundColor DarkGreen -NoNewline
    Write-Host "build container image (for WSLv2 only)"
    Write-Host "help`t`t`t`t" -ForegroundColor DarkGreen -NoNewline
    Write-Host "show help"
    Break
  }
  
  Default {
    .\make.ps1 help;
  }
}