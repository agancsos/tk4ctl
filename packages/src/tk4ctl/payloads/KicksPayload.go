package payloads
import (
	"fmt"
	"tk4ctl/internals"
	"tk4ctl/services"
)

type KicksPayload struct {
}

func (x KicksPayload) Validate(connectionInfo *internals.ConnectionInfo, args map[string]interface{}) error {
	if connectionInfo.HostName == "" { return internals.InvalidParameterError("HOSTNAME"); }
    if connectionInfo.Port == -1 { return internals.InvalidParameterError("PORT"); }
	if connectionInfo.Username == "HERC01" { return internals.InvalidParameterError("USERNAME"); }
    if connectionInfo.PAT == "" { return internals.InvalidParameterError("PAT"); }
	if args["cmd"] == nil || args["cmd"] == "" { return internals.InvalidParameterError("KICKS command"); }
    return nil;
}

func (x KicksPayload) Invoke(connectionInfo *internals.ConnectionInfo, args map[string]interface{}, dryrun bool) (error, map[string]interface{}) {
	if dryrun {
        return nil, map[string]interface{} { "data": internals.MSG_DRYRUN, };
    }
	err, _ := services.SendTelnetTso(connectionInfo, fmt.Sprintf(`//%sG JOB CLASS=A,MSGCLASS=A,MSGLEVEL=(1,1),REGION=3M,
//            USER=%s,PASSWORD=%s,TIME=1440
//PS010 EXEC PGM=IKJEFT01                               
//SYSTSPRT DD SYSOUT=*
//SYSTSIN DD *
 exec KICKSSYS.V1R5M0.CLIST(KICKS)
/*`, connectionInfo.Username, connectionInfo.Username, connectionInfo.PAT));
	if err != nil {
		return err, nil;
	}

	var raw = fmt.Sprintf(`//%sG JOB CLASS=A,MSGCLASS=A,MSGLEVEL=(1,1),REGION=3M,
//            USER=%s,PASSWORD=%s,TIME=1440
//PS010 EXEC PGM=IKJEFT01                               
//SYSTSPRT DD SYSOUT=*
//SYSTSIN DD *
 %s
/*`, connectionInfo.Username, connectionInfo.Username, connectionInfo.PAT, args["cmd"].(string));
    err, rsp := services.SendTelnetTso(connectionInfo, raw);
	services.SendTelnetTso(connectionInfo, fmt.Sprintf(`//%sG JOB CLASS=A,MSGCLASS=A,MSGLEVEL=(1,1),REGION=3M,
//            USER=%s,PASSWORD=%s,TIME=1440
//PS010 EXEC PGM=IKJEFT01                               
//SYSTSPRT DD SYSOUT=*
//SYSTSIN DD *
 KSSF
/*`, connectionInfo.Username, connectionInfo.Username, connectionInfo.PAT));

	return err, rsp;
}
