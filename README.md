# Logx
An elegant log for dev right in hand
[![CircleCI](https://circleci.com/gh/linnv/logx.svg?style=shield)](https://circleci.com/gh/linnv/logx)

### Example
```
logx.Debugln(111)
logx.Warnln(111)
logx.Errorln(111)

logx.EnableDevMode(false)
logx.Debugln("no log output on level debug ")
logx.Warnln("log output")
logx.Errorln("log output")

logx.EnableDevMode(true)
logx.Debugln("log output on level debug with dev enabled")
```
check `example/log.go` for practical usage
