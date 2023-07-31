package services
import (
	"os"
	"io/ioutil"
	"encoding/json"
)

type ConfigurationService struct {
	Container                 map[string]interface{}
}

var __configuration_service_instance__ *ConfigurationService;
func ConfigurationServiceInstance(path string) *ConfigurationService {
	if __configuration_service_instance__ == nil {
		__configuration_service_instance__ = &ConfigurationService{};
		__configuration_service_instance__.Container = map[string]interface{}{};
		_, check := os.Stat(path); if check == nil {
			var raw, _ = ioutil.ReadFile(path);
			_ = json.Unmarshal(raw, __configuration_service_instance__.Container);
		}
	}
	return __configuration_service_instance__;
}

