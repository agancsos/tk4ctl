package services
import (
	"tk4ctl/internals"
	"net"
	"hercules"
	"hercules/actions"
	"net/http"
	"io/ioutil"
	"fmt"
	"time"
	"strings"
)

// TODO: Access /mvs38/mvs38/prt/prt00e.txt on the remote system
func ReadJES2Log(connectionInfo *internals.ConnectionInfo) (error, string) {
	var client   = http.Client{};
    req, err      := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/tk4ctl_log.txt", connectionInfo.HostName, connectionInfo.Port), nil);
    if err != nil { return err, ""; }
    rsp, err := client.Do(req);
    if err != nil { return err, ""; }
    rspData, err := ioutil.ReadAll(rsp.Body);
    if err != nil { return err, ""; }
    var rawLog = "";
    var lines = strings.Split(string(rspData), "\n");
    var i = len(lines) - 1;
    for ;; {
        if strings.Split(lines[i], " ")[0] == "IEF375I" {
            break;
        }
        rawLog = fmt.Sprintf("\n%s%s", lines[i], rawLog);
        i -= 1;
    }
	return nil, rawLog;
}

func SendTelnetJCL(connectionInfo *internals.ConnectionInfo, raw string) (error, map[string]interface{}) {
	client, err    := net.Dial("tcp", fmt.Sprintf("%s:3505", connectionInfo.HostName));
	if err != nil { return err, nil; }
	_, err = client.Write([]byte(fmt.Sprintf("%s", raw)));
	client.Close();
	if err != nil {
		return err, nil;
	}
	time.Sleep(3 *time.Second);
	client2, err := hercules.NewClient(fmt.Sprintf("http://%s:%d", connectionInfo.HostName, connectionInfo.Port));
	if err != nil {
		return err, nil;
	}
	rst, err := actions.GetSysLog(client2, 30);
	return err, rst;
}

func SendTelnetTso(connectionInfo *internals.ConnectionInfo, raw string) (error, map[string]interface{}) {
    client, err    := net.Dial("tcp", fmt.Sprintf("%s:3505", connectionInfo.HostName));
    if err != nil { return err, nil; }
    _, err = client.Write([]byte(fmt.Sprintf("%s", raw)));
    client.Close();
    if err != nil {
        return err, nil;
    }
	time.Sleep(3 *time.Second);
	err, rawLog := ReadJES2Log(connectionInfo);
    return err, map[string]interface{}{"data": rawLog, };
}

func SendTelnetKicks(connectionInfo *internals.ConnectionInfo, raw string) (error, map[string]interface{}) {
    client, err    := net.Dial("tcp", fmt.Sprintf("%s:3505", connectionInfo.HostName));
    if err != nil { return err, nil; }
    _, err = client.Write([]byte(fmt.Sprintf("%s", raw)));
    client.Close();
    if err != nil {
        return err, nil;
    }
    time.Sleep(3 *time.Second);
    var client2   = http.Client{};
    req, err      := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/tk4ctl_log.txt", connectionInfo.HostName, connectionInfo.Port), nil);
    if err != nil { return err, nil; }
    rsp, err := client2.Do(req);
    if err != nil { return err, nil; }
    rspData, err := ioutil.ReadAll(rsp.Body);
    if err != nil { return err, nil; }
    var rawLog = "";
    var lines = strings.Split(string(rspData), "\n");
    var i = len(lines) - 1;
    for ;; {
        if strings.Split(lines[i], " ")[0] == "IEF375I" {
            break;
        }
        rawLog = fmt.Sprintf("\n%s%s", lines[i], rawLog);
        i -= 1;
    }
    return err, map[string]interface{}{"data": string(rawLog), };
}

