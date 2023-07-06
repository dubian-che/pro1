```
这是怪物猎人配装器项目MHDict开发分支
```

### 步骤

```
go mod init MHDict
go mod tidy
go get fyne.io/fyne/v2
go get github.com/flopp/go-findfont
go install fyne.io/fyne/v2/cmd/fyne@latest
fyne bundle resource/STKaiti.ttf >> bundle.go
go build -ldflags "-s -w -H=windowsgui" .\main.go

#macOS
fyne package -os darwin -icon myapp.png
#linux 和 Windows
fyne package -os linux -icon myapp.png
fyne package -os windows -icon myapp.png
```

