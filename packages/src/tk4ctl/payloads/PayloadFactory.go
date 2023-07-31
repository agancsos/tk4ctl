package payloads
import (
	"tk4ctl/internals"
)

func PayloadFactory(operation string) internals.IPayload {
	switch(operation) {
		case "herc":
			return &HerculesPayload{};
		case "tso":
			return &TsoPayload{};
		case "jcl":
			return &JclPayload{};
		case "kicks":
			return &KicksPayload{};
		default:
			return nil;
	}
}
