### An elegant log for dev right in hand

By default, logx will output log to stdout or stderr, if you want make log output to file,there are two way you can configure logx

- By flag, only two flags are supported:
	- first invoke `logx.InitAndParsed()` in your go file
	- carry argument type string `-logxFile` while running you app, if empty only log will be output to stdout or stderr
	- carry argument type boolean `-logxDev` while running you app, if false all log of debug level will be ignored,the default value is true

- By invoking `logx.LoadConfJson(conf []byte)` in your go file, there is a json file sample for reference:
	- [json sample](#json sample): be careful if you don't provied `DevMode` in the json file the default value is false, although it's true by flag which default is true

#### <a name="json sample"></a>json sample
```
{
	"DisableBuffer": false,
	"Maxbuffer": "2MB",
	"ToDifferentFile": true,
	"DevMode": true,
	"FilePath": "/Users/Jialin/golang/src/github.com/linnv"
}
```
