package payloads
import (
	"fmt"
	"tk4ctl/internals"
	"tk4ctl/services"
)
type TsoPayload struct {
}

func (x TsoPayload) Validate(connectionInfo *internals.ConnectionInfo, args map[string]interface{}) error {
    if connectionInfo.HostName == "" { return internals.InvalidParameterError("HOSTNAME"); }
    if connectionInfo.Port == -1 { return internals.InvalidParameterError("PORT"); }
	if connectionInfo.Username == "" { return internals.InvalidParameterError("USERNAME"); }
	if connectionInfo.PAT == "" { return internals.InvalidParameterError("PAT"); }
	if args["cmd"] == nil || args["cmd"] == "" { return internals.InvalidParameterError("TSO command"); }
    return nil;
}

func (x TsoPayload) Invoke(connectionInfo *internals.ConnectionInfo, args map[string]interface{}, dryrun bool) (error, map[string]interface{}) {
	var raw = fmt.Sprintf(`//%sG JOB CLASS=A,MSGCLASS=A,MSGLEVEL=(1,1),REGION=3M,
//            USER=%s,PASSWORD=%s,TIME=1440
//PS010 EXEC PGM=IKJEFT01                               
//SYSTSPRT DD SYSOUT=*
//SYSTSIN DD *
 %s
/*`, connectionInfo.Username, connectionInfo.Username, connectionInfo.PAT, args["cmd"].(string));
	if dryrun {
		return nil, map[string]interface{} { "data": internals.MSG_DRYRUN, };
	}
	return services.SendTelnetTso(connectionInfo, raw);
}
