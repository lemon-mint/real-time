# Golang CI Tools Report

Report generated at: 2022-01-13T00:27:50+09:00

Go version: go1.17.6

golang-ci-tools version: 0.1.0-dev

## staticcheck

```
🎉  No staticcheck errors found!
```

## gosec

```
[gosec] 2022/01/13 00:27:55 Including rules: default
[gosec] 2022/01/13 00:27:55 Excluding rules: default
[gosec] 2022/01/13 00:27:55 Import directory: G:\mega\golang\projectv4\real-time
[gosec] 2022/01/13 00:27:57 Checking package: main
[gosec] 2022/01/13 00:27:57 Checking file: G:\mega\golang\projectv4\real-time\main.go
Results:


[[30;43mG:\mega\golang\projectv4\real-time\main.go:146[0m] - G102 (CWE-200): Binds to all network interfaces (Confidence: HIGH, Severity: MEDIUM)
    145: 
  > 146: 	ln, err := net.Listen("tcp", lnHost)
    147: 	if err != nil {



[[37;40mG:\mega\golang\projectv4\real-time\main.go:151-154[0m] - G104 (CWE-703): Errors unhandled. (Confidence: HIGH, Severity: LOW)
    150: 
  > 151: 	http.Serve(ln, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
  > 152: 		rw.Header().Set("Access-Control-Allow-Origin", "*")
  > 153: 		mux.ServeHTTP(rw, r)
  > 154: 	}))
    155: }



[1;36mSummary:[0m
  Gosec  : dev
  Files  : 1
  Lines  : 155
  Nosec  : 0
  Issues : [1;31m2[0m


```

## gocap

```
github.com/lemon-mint/real-time (file, network)

github.com/beevik/ntp (network)
github.com/modern-go/concurrent (runtime, file)
github.com/modern-go/reflect2 (runtime)
golang.org/x/net/internal/socket (file, network, runtime, syscall)
golang.org/x/net/ipv4 (network, runtime)
golang.org/x/sys/windows (network, runtime, syscall)

```

## go-licenses

| Package Name | License File | License |
| --- | --- | --- |
| [github.com/modern-go/concurrent](https://pkg.go.dev/github.com/modern-go/concurrent) | [https://github.com/modern-go/concurrent/blob/master/LICENSE](https://github.com/modern-go/concurrent/blob/master/LICENSE) | Apache-2.0 |
| [github.com/modern-go/reflect2](https://pkg.go.dev/github.com/modern-go/reflect2) | [https://github.com/modern-go/reflect2/blob/master/LICENSE](https://github.com/modern-go/reflect2/blob/master/LICENSE) | Apache-2.0 |
| [github.com/lemon-mint/real-time](https://pkg.go.dev/github.com/lemon-mint/real-time) | [https://github.com/lemon-mint/real-time/blob/master/LICENSE](https://github.com/lemon-mint/real-time/blob/master/LICENSE) | Unlicense |
| [github.com/beevik/ntp](https://pkg.go.dev/github.com/beevik/ntp) | [https://github.com/beevik/ntp/blob/master/LICENSE](https://github.com/beevik/ntp/blob/master/LICENSE) | BSD-2-Clause |
| [golang.org/x/net](https://pkg.go.dev/golang.org/x/net) | [Unknown](Unknown) | BSD-3-Clause |
| [golang.org/x/sys](https://pkg.go.dev/golang.org/x/sys) | [Unknown](Unknown) | BSD-3-Clause |
| [github.com/json-iterator/go](https://pkg.go.dev/github.com/json-iterator/go) | [https://github.com/json-iterator/go/blob/master/LICENSE](https://github.com/json-iterator/go/blob/master/LICENSE) | MIT |



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

