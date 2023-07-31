package internals
import (
	"fmt"
	"errors"
)

var MSG_DRYRUN          = "DryRun was enabled....";

func InvalidParameterError(parameter string) error {
    return errors.New(fmt.Sprintf("Invalid value for parameter %s", parameter));
}

