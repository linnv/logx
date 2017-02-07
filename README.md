### An elegant log right in hand

By default, logx will output log to stdout or stderr, if you want make log output to file,there are two way you can configure logx

- Redirect log to file by configuring flag:
	- invoke `logx.InitAndParsed()`
	- carry argument `-logxfile` while running you app, if empty only log will be output to stdout or stderr

- Redirect log to file by invoking `logx.LoadConfJson(conf []byte)`:
	- [json sample](#json sample):


```
{
	"DisableBuffer": false,
	"Maxbuffer": "3MB",
	"ToDifferentFile": true,
	"FilePath": "/Users/Jialin/golang/src/github.com/linnv"
}
```
