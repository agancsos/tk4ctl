package payloads
import (
	"fmt"
	"tk4ctl/internals"
	"hercules"
	"hercules/actions"
)

type HerculesPayload struct {
}

func (x HerculesPayload) Validate(connectionInfo *internals.ConnectionInfo, args map[string]interface{}) error {
	if connectionInfo.HostName == "" { return internals.InvalidParameterError("HOSTNAME"); }
	if connectionInfo.Port == -1 { return internals.InvalidParameterError("PORT"); }
	if args["cmd"] == nil || args["cmd"] == "" { return internals.InvalidParameterError("Hercules command"); }		
    return nil;
}

func (x HerculesPayload) Invoke(connectionInfo *internals.ConnectionInfo, args map[string]interface{}, dryrun bool) (error, map[string]interface{}) {
	var client, err = hercules.NewClient(fmt.Sprintf("http://%s:%d", connectionInfo.HostName, connectionInfo.Port));
	if err != nil {
		return err, nil;
	}
	if dryrun {
		return nil, map[string]interface{} { "data": internals.MSG_DRYRUN, }; 
	}
	rsp, err := actions.ExecuteHerculesCommand(client, args["cmd"].(string));
	if err != nil {
		return err, nil;
	}
	rsp, err  = actions.GetSysLog(client, 20);
    return err, rsp;
}

