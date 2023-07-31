package main
import (
	"fmt"
	"os"
	"common"
	"tk4ctl"
	"tk4ctl/internals"
	"tk4ctl/payloads"
	"tk4ctl/services"
	"runtime"
)

func main() {
	var ss            = &common.SystemService{};
	ss.ModulePath     = os.Getenv("HOME");
	services.EnsureValidCommand("Malformed command...", 1);
	var params     = common.ParseParameters(",");
	var operation  = os.Args[1];
	var configPath = common.LookupParameter(params, "--config", fmt.Sprintf("%s/tk4ctl.json", ss.BuildModuleContainerPath())).(string); 
	var isHelp     = common.LookupParameter(params, "--help", false).(bool);
	var isVersion  = common.LookupParameter(params, "--version", false).(bool);
	var isDryRun   = common.LookupParameter(params, "--dry", false).(bool);
	var config     = services.ConfigurationServiceInstance(configPath);
	if isHelp {
		tk4ctl.HelpMenu();
	} else if isVersion {
		println(fmt.Sprintf("%s v. %s", tk4ctl.APPLICATION_NAME, tk4ctl.APPLICATION_VERSION));
	} else {
		services.EnsureSupportedOperation(operation, 2);
		services.EnsureSupportedPlatform(runtime.GOOS, 3);
		var connectionInfo = internals.NewConnectionInfoFromConfig(config.Container);
		internals.SetOverrideConnectionInfo(connectionInfo, params);
		var payload = payloads.PayloadFactory(operation);
		var opParamsRaw = common.LookupParameter(params, "-e", "").(string);
		var opParams    = common.ParseExtendedParameters(opParamsRaw);
		opParams["-f"] = params["-f"];
    	opParams["-m"] = params["-m"];
		services.EnsurePayload(payload, connectionInfo, opParams, 4);
		err, rst := payload.Invoke(connectionInfo, opParams, isDryRun);
		if err != nil {
            println(fmt.Sprintf("\033[31m%s\033[m", err.Error()));
            os.Exit(5);
        }
		println(fmt.Sprintf("\033[36m%s\033[m", rst["data"]));
	}
	os.Exit(0);
}

