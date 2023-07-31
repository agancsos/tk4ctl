package internals
type IPayload interface {
	Validate(connectionInfo *ConnectionInfo, args map[string]interface{}) error
	Invoke(connectionInfo *ConnectionInfo, args map[string]interface{}, dryrun bool)   (error, map[string]interface{})
}

type DummyPayload struct {}
func (x DummyPayload) Validate(connectionInfo *ConnectionInfo, args map[string]interface{}) error {
	return nil;
}

func (x DummyPayload) Invoke(connectionInfo *ConnectionInfo, args map[string]interface{}, dryrun bool) (error, map[string]interface{}) {
	return nil, map[string]interface{}{};
}

