package internals
import (
	"encoding/base64"
	"strings"
	"common"
	"strconv"
)

type ConnectionInfo struct {
	Username              string
	PAT                   string
	HostName              string
	Port                  int
	Version               int
}

func NewConnectionInfo(username string, pat string, hostname string, port int) *ConnectionInfo {
	var rst      = &ConnectionInfo{};
	rst.Username = username;
	rst.PAT      = pat;
	rst.HostName = hostname;
	rst.Port     = port;
	rst.Version  = 1;
	return rst;
}

func NewConnectionInfoFromConfig(config map[string]interface{}) *ConnectionInfo {
	var rst      = &ConnectionInfo{};
	for k, v := range config {
		switch (k) { 
			case "username":
				rst.Username = v.(string);
				break;
			case "pat":
				var decoded, _ = base64.StdEncoding.DecodeString(v.(string));
				var comps      = strings.Split(string(decoded), ":");   
        		if len(comps) > 1 {
					rst.PAT = comps[1];
				}
				break;
			case "hostname":
				rst.HostName = v.(string);
				break;
			case "port":
				rst.Port = v.(int);
				break;
			case "version":
				rst.Version = v.(int);
				break;
			default: break;
		}
	}
	return rst;
}

func SetOverrideConnectionInfo(connectionInfo *ConnectionInfo, params map[string]interface{}) {
	var fullCredentials = common.LookupParameter(params, "-c", "").(string);
	if fullCredentials != "" {
		var decoded, _ = base64.StdEncoding.DecodeString(fullCredentials);
        var comps      = strings.Split(string(decoded), ":");
		if len(comps) > 1 {
			connectionInfo.Username = comps[0];
			connectionInfo.PAT      = comps[1];
		}
	}
	var fullHostname = common.LookupParameter(params, "-h", "").(string);
    if fullHostname != "" {
        var comps      = strings.Split(string(fullHostname), ":");   
        if len(comps) > 1 {
            connectionInfo.HostName = comps[0];
			var tempP, _ = strconv.Atoi(comps[1]);
            connectionInfo.Port     = tempP;
        }
    }
	connectionInfo.Version = common.LookupParameter(params, "-V", 1).(int);
}

