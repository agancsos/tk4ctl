package services
import (
	"fmt"
	"common"
	"os"
	"tk4ctl"
	"tk4ctl/internals"
)

func EnsureValidCommand(message string, exitCode int) {
	if len(os.Args) < 2 {
        println(fmt.Sprintf("\033[31m%s\033[m", message));
        os.Exit(exitCode);
    }
}

func EnsureSupportedOperation(operation string, exitCode int) {
	_, check := tk4ctl.SUPPORTED_OPERATIONS[operation]; if !check {
        println(fmt.Sprintf("\033[31mInvalid operation (%s)\033[m", operation));
        os.Exit(exitCode);
	}
}

func EnsureSupportedPlatform(platform string, exitCode int) {
	check, _ := common.ListContains(tk4ctl.SUPPORTED_PLATFORMS, platform);
    if !check {
        println(fmt.Sprintf("\033[31mUnsupported platform (%s)\033[m", platform));
        os.Exit(exitCode);
    }
}

func EnsurePayload(payload internals.IPayload, connectionInfo *internals.ConnectionInfo, opParams map[string]interface{}, exitCode int) {
	err := payload.Validate(connectionInfo, opParams);
    if err != nil {
        println(fmt.Sprintf("\033[31mMissing required parameters (%s)\033[m", err.Error()));
        os.Exit(exitCode);
    }
}
