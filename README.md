# Start "Go" in 1 day

## Project setup

1. 安裝GO官網下載安裝檔。 [This is an external link to GO](https://go.dev/dl/)
2. Windows 環境變數設定
> 進階系統設定=>點選環境變數=>點選系統變數下的新增
> + 輸入變數名稱：GOROOT 以及指定路徑值(ex: [C:\Program Files\Go])
> + 輸入變數名稱：GOPATH 以及指定路徑值(ex: [C:\users\youName\g])
> + 電腦需要重新啟動。

3. 開發環境(VS CODE) : 安裝 Go 語言包
   ![go_package](https://i.imgur.com/CTl784t.png)

4. 安裝 Golang 相關工具
   按 Command + Shift + p 呼叫命令列視窗，輸入 Go: Install/Update tools 安裝/更新所有的工具

5. 裝完擴充功能與相關工具後，寫一段測試程式後按 F5 執行試試。參考 test.go
6. 在專案資料夾建立一個 .vscode/launch.json
> 避免出現類似下面這種錯誤
> Build Error: go build -o /.../gotour/__debug_bin -gcflags all=-N -l .
> go: cannot find main module; see 'go help modules' (exit status 1)

7. 建立 go.mod 檔案 
   按 Command + Shift + p , 輸入 Go: Initialize go.mod 後，輸入模組名稱即可(**ex: test.go**)# GoTour
