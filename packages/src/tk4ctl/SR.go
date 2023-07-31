package tk4ctl
import (
	"common"
	"fmt"
	"math/rand"
    "time"
)

var APPLICATION_NAME    = "TK4CTL";
var APPLICATION_AUTHOR  = "Abel Gancsos";
var APPLICATION_VERSION = "1.0.0.0";
var SUPPORTED_PLATFORMS = []string {
	"unix",
	"linux",
	"darwin",
}; 

var SUPPORTED_OPERATIONS = map[string]string {
	"tso"   : "Runs a raw TSO command",
	"jcl"   : "Runs a JCL job via sockdev",
	"herc"  : "Runs a raw Hercules command through HTTP",
	"kicks" : "Executes a KICKS process",
};

var FLAGS = map[string]string {
	"-h"        : "ip:port combination",
	"-c"        : "Base64 encoded MVS credentials to run JCL jobs under",
	"-m"        : "MVS dataset member or Hercules command to run",
	"-f"        : "Local file to run",
	"-e"        : "Command arguments",
	"--version" : "Prints the version of the utility",
	"--help"    : "Prints the help menu",
};

func GenerateFilePath(base string) string {
    rand.Seed(time.Now().UnixNano());
    return fmt.Sprintf("%s/temp_%d.jcl", base, rand.Intn(10000));
}

func HelpMenu() {
	println(common.PadLeft("", 80, "#"));
	println(fmt.Sprintf("# %s#", common.PadLeft(fmt.Sprintf("Name       : %s", APPLICATION_NAME), 77, " ")));
	println(fmt.Sprintf("# %s#", common.PadLeft(fmt.Sprintf("Author     : %s", APPLICATION_AUTHOR), 77, " ")));
	println(fmt.Sprintf("# %s#", common.PadLeft(fmt.Sprintf("Version    : v  %s", APPLICATION_VERSION), 77, " ")));
	println(fmt.Sprintf("# %s#", common.PadLeft(fmt.Sprintf("Operations :"), 77, " ")));
	for k, v := range SUPPORTED_OPERATIONS {
        println(fmt.Sprintf("#    %s#", common.PadLeft(fmt.Sprintf("%s: %s", k, v), 74, " ")));
    }
	println(fmt.Sprintf("# %s#", common.PadLeft(fmt.Sprintf("Flags   :"), 77, " ")));
	for k, v := range FLAGS {
		println(fmt.Sprintf("#    %s#", common.PadLeft(fmt.Sprintf("%s: %s", k, v), 74, " ")));
	}
	println(common.PadLeft("", 80, "#"));
}

