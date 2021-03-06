# Golang CI Tools Report

Report generated at: 2022-01-13T07:17:23+09:00

Go version: go1.17.6

golang-ci-tools version: 0.1.1-dev

## staticcheck

```
🎉  No staticcheck errors found!
```

## gosec

```
[gosec] 2022/01/13 07:17:29 Including rules: default
[gosec] 2022/01/13 07:17:29 Excluding rules: default
[gosec] 2022/01/13 07:17:30 Import directory: G:\mega\golang\projectv4\real-time
[gosec] 2022/01/13 07:17:32 Checking package: main
[gosec] 2022/01/13 07:17:32 Checking file: G:\mega\golang\projectv4\real-time\main.go
Results:


Summary:
  Gosec  : dev
  Files  : 1
  Lines  : 158
  Nosec  : 1
  Issues : 0


```

## gocap

```
github.com/lemon-mint/real-time (file, network)

github.com/beevik/ntp (network)
github.com/json-iterator/go (reflect, unsafe)
github.com/modern-go/concurrent (file, reflect)
github.com/modern-go/reflect2 (reflect, unsafe)
golang.org/x/net/internal/socket (file, network, syscall, unsafe)
golang.org/x/net/ipv4 (network, unsafe)
golang.org/x/sys/internal/unsafeheader (unsafe)
golang.org/x/sys/windows (network, syscall, unsafe)

```

## go-licenses

| Package Name | License File | License |
| --- | --- | --- |
| [github.com/modern-go/reflect2](https://pkg.go.dev/github.com/modern-go/reflect2) | [https://github.com/modern-go/reflect2/blob/master/LICENSE](https://github.com/modern-go/reflect2/blob/master/LICENSE) | Apache-2.0 |
| [github.com/lemon-mint/real-time](https://pkg.go.dev/github.com/lemon-mint/real-time) | [https://github.com/lemon-mint/real-time/blob/master/LICENSE](https://github.com/lemon-mint/real-time/blob/master/LICENSE) | Unlicense |
| [github.com/beevik/ntp](https://pkg.go.dev/github.com/beevik/ntp) | [https://github.com/beevik/ntp/blob/master/LICENSE](https://github.com/beevik/ntp/blob/master/LICENSE) | BSD-2-Clause |
| [golang.org/x/net](https://pkg.go.dev/golang.org/x/net) | [Unknown](Unknown) | BSD-3-Clause |
| [golang.org/x/sys](https://pkg.go.dev/golang.org/x/sys) | [Unknown](Unknown) | BSD-3-Clause |
| [github.com/json-iterator/go](https://pkg.go.dev/github.com/json-iterator/go) | [https://github.com/json-iterator/go/blob/master/LICENSE](https://github.com/json-iterator/go/blob/master/LICENSE) | MIT |
| [github.com/modern-go/concurrent](https://pkg.go.dev/github.com/modern-go/concurrent) | [https://github.com/modern-go/concurrent/blob/master/LICENSE](https://github.com/modern-go/concurrent/blob/master/LICENSE) | Apache-2.0 |



## Dependencies

Total dependencies: 23
<details><summary>Show Full Dependencies</summary>

 - github.com/lemon-mint/real-time github.com/beevik/ntp@v0.3.0
 - github.com/lemon-mint/real-time github.com/json-iterator/go@v1.1.12
 - github.com/lemon-mint/real-time github.com/modern-go/concurrent@v0.0.0-20180306012644-bacd9c7ef1dd
 - github.com/lemon-mint/real-time github.com/modern-go/reflect2@v1.0.2
 - github.com/lemon-mint/real-time github.com/stretchr/testify@v1.7.0
 - github.com/lemon-mint/real-time golang.org/x/net@v0.0.0-20220111093109-d55c255bac03
 - github.com/lemon-mint/real-time golang.org/x/sys@v0.0.0-20220111092808-5a964db01320
 - github.com/json-iterator/go@v1.1.12 github.com/davecgh/go-spew@v1.1.1
 - github.com/json-iterator/go@v1.1.12 github.com/google/gofuzz@v1.0.0
 - github.com/json-iterator/go@v1.1.12 github.com/modern-go/concurrent@v0.0.0-20180228061459-e0a39a4cb421
 - github.com/json-iterator/go@v1.1.12 github.com/modern-go/reflect2@v1.0.2
 - github.com/json-iterator/go@v1.1.12 github.com/stretchr/testify@v1.3.0
 - github.com/stretchr/testify@v1.7.0 github.com/davecgh/go-spew@v1.1.0
 - github.com/stretchr/testify@v1.7.0 github.com/pmezard/go-difflib@v1.0.0
 - github.com/stretchr/testify@v1.7.0 github.com/stretchr/objx@v0.1.0
 - github.com/stretchr/testify@v1.7.0 gopkg.in/yaml.v3@v3.0.0-20200313102051-9f266ea9e77c
 - golang.org/x/net@v0.0.0-20220111093109-d55c255bac03 golang.org/x/sys@v0.0.0-20210423082822-04245dca01da
 - golang.org/x/net@v0.0.0-20220111093109-d55c255bac03 golang.org/x/term@v0.0.0-20201126162022-7de9c90e9dd1
 - golang.org/x/net@v0.0.0-20220111093109-d55c255bac03 golang.org/x/text@v0.3.6
 - github.com/stretchr/testify@v1.3.0 github.com/davecgh/go-spew@v1.1.0
 - github.com/stretchr/testify@v1.3.0 github.com/pmezard/go-difflib@v1.0.0
 - github.com/stretchr/testify@v1.3.0 github.com/stretchr/objx@v0.1.0
 - gopkg.in/yaml.v3@v3.0.0-20200313102051-9f266ea9e77c gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405
</details>

