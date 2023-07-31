package payloads
import (
	"tk4ctl/internals"
	 "tk4ctl/services"
	"strings"
	"fmt"
	"io/ioutil"
)
type JclPayload struct {
}

func (x JclPayload) Validate(connectionInfo *internals.ConnectionInfo, args map[string]interface{}) error {
    if connectionInfo.HostName == "" { return internals.InvalidParameterError("HOSTNAME"); }
    if connectionInfo.Port == -1 { return internals.InvalidParameterError("PORT"); }
	if connectionInfo.Username == "" { return internals.InvalidParameterError("USERNAME"); }
	if connectionInfo.PAT == "" { return internals.InvalidParameterError("PAT"); }
    if (args["-m"] == nil || args["-m"] == "") && (args["-f"] == nil || args["-f"] == "") { 
		return internals.InvalidParameterError("PDS"); 
	}
	if args["-m"] == nil { args["-m"] = ""; }
	if args["-f"] == nil { args["-f"] = ""; }
    return nil;
}

func (x JclPayload) Invoke(connectionInfo *internals.ConnectionInfo, args map[string]interface{}, dryrun bool) (error, map[string]interface{}) {
	var raw = "";
	if args["-f"].(string) != "" && args["-m"].(string) == "" {
		raw1, err := ioutil.ReadFile(args["-f"].(string));
		if err != nil { return err, nil; }
		var raw2 = string(raw1);
		if !strings.Contains(raw2, "USER=") && !strings.Contains(raw2, "PASSWORD=") {
			var lines = strings.Split(raw2, "\n");
			for i := 0; i < len(lines) - 1; i++ {
				if strings.Contains(lines[i + 1], "EXEC") {
					raw += fmt.Sprintf("//            USER=%s,PASSWORD=%s\n", connectionInfo.Username, connectionInfo.PAT);
				} else {
					raw += fmt.Sprintf("%s\n", lines[i]);
				}
			}
			raw += fmt.Sprintf("%s\n", lines[len(lines) - 1]);
		} else {
			raw = raw2;
		}
	} else if args["-m"].(string) == "" && args["-f"].(string) != "" {
		raw = fmt.Sprintf(`//%sG JOB CLASS=A,MSGCLASS=H,MSGLEVEL=(1,1),REGION=3M,
//            USER=%s,PASSWORD=%s,TIME=1440
//STEP001  EXEC PGM=IEBGENER
//SYSUT1 DD DSN=%s,
//          UNIT=DISK,DISP=SHR
//SYSUT2 DD SYSOUT=*
//SYSPRINT DD SYSOUT=*
//SYSIN DD DUMMY`, connectionInfo.Username, connectionInfo.Username, connectionInfo.PAT, args["-m"].(string));
	}
	if dryrun {
		return nil, map[string]interface{} { "data": internals.MSG_DRYRUN, };
	}
	err, rsp := services.SendTelnetJCL(connectionInfo, raw);
	return err, rsp;
}
