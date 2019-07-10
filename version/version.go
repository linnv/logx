package version

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var VERSION = "None"
var BUILDTIME = "None"
var GITHASH = "NoneHash"
var GITBRANCH = "NoneBranch"
var COYPRIGHT = " Copyright ©2018-%d jialinwu"

func ReadBuildInfoNoExit() {
	if len(GITHASH) > 0 {
		fmt.Printf(" BuildTime:%s\n Version:%s\n Branch:%s\n Hash:%s\n%s\n", BUILDTIME, VERSION, GITBRANCH, GITHASH[:7], fmt.Sprintf(COYPRIGHT, time.Now().Year()))
	} else {
		fmt.Printf(" BuildTime:%s\n Version:%s\n Branch:%s\n Hash:%s\n%s\n", BUILDTIME, VERSION, GITBRANCH, GITHASH, fmt.Sprintf(COYPRIGHT, time.Now().Year()))
	}
}

func ReadBuildInfo() {
	if len(os.Args) > 1 {
		man := strings.TrimSpace(os.Args[1])
		switch man {
		case "version", "buildtime", "hash", "info", "branch", "v", "i":
			ReadBuildInfoNoExit()
			os.Exit(0)
		}
	}
	ReadBuildInfoNoExit()
}
