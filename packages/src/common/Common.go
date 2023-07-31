package common
import "C"
import (
	"unsafe"
	"reflect"
	exec "os/exec"
	"fmt"
	"strings"
	"encoding/json"
	"unicode"
	"sort"
	"os"
)
func PadRight(str string, le int, pad string) string {
	if len(str) > le {
		return str[0:le];
	}
	result := "";
	for i := len(str); i < le; i++ {
		result += pad;
	}
	return result + str;
}

func CleanString(a string) string {
	var result = strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, a);
	return result;
}

func PadLeft(str string, le int, pad string) string {
    if len(str) > le {
        return str[0:le];
    }
    result := "";
    for i := len(str); i < le; i++ {
        result += pad;
    }
    return str + result;
}

func CStr(s string) *C.char {
	h := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return (*C.char)(unsafe.Pointer(h.Data))
}

func RunCmd(cmd string) string {
	args := strings.Fields(cmd);
	stdout, err := exec.Command(args[0], args[1:]...).Output();
	if err != nil {
		return fmt.Sprintf("%v", err);
	}
	return string(stdout);
}

func StrToDictionary(s []byte) map[string]interface{} {
	var obj map[string]interface{};
	json.Unmarshal(s, &obj);
	return obj;
}

func DictionaryToJsonString (a map[string]interface{}) string {
	var result = "{";
	for key, value := range a {
		result += fmt.Sprintf("\"%s\":\"%v\"", key, value);
	}
	result += "}";
	return result;
}

func StrDictionaryToJsonString (a map[string]string) string {
    var result = "{";
	var i = 0;
    for key, value := range a {
		if i > 0 { result += ","; }
        result += fmt.Sprintf("\"%s\":\"%s\"", key, value);
		i++;
    }
    result += "}";
    return result;
}

func ToConstStr(a string) *C.uchar {
	return (*C.uchar)(unsafe.Pointer(&[]byte(a)[0]))
}

func ArgsToDictionary(args []string) map[string]string {
    var result       = map[string]string{};
    for i := 0; i < len(args) - 1; i++ {
        result[args[i]] = args[i + 1];
    }
    return result;
}

func ParseParameters(delim string) map[string]interface{} {
    var rst = map[string]interface{}{};
    if delim == "" {
        delim = ",";
    }
    for i := range os.Args {
        if len(os.Args[i]) > 0 && string(os.Args[i][0]) == "-" && (i == len(os.Args) - 1 || (i < len(os.Args) - 1 && len(os.Args[i + 1]) > 0 && string(os.Args[i + 1][0]) == "-")) {
            rst[os.Args[i]] = true;
        } else if i < len(os.Args) - 1 && strings.Contains(os.Args[i + 1], delim) {
            rst[os.Args[i]] = strings.Split(os.Args[i + 1], delim);
        } else if i < len(os.Args) - 1 {
            rst[os.Args[i]] = os.Args[i + 1];
        }
    }
    return rst;
}

func ParseExtendedParameters(raw string) map[string]interface{} {
	var rst   = map[string]interface{}{};
	var pairs = strings.Split(raw, ",");
	for _, pair := range pairs {
		var comps = strings.Split(pair, "=");
		if len(comps) == 2 {
			rst[strings.Trim(comps[0], " ")] = strings.Trim(comps[1], " ");
		}
	}  
	return rst;
}

func LookupParameterValue(key string, parameters map[string]string, defaultValue string) string {
    if parameters[key] != "" {
        return parameters[key];
    }
    return defaultValue;
}

func ListContains(haystack []string, needle string) (bool, int) {
	for i, x := range haystack {
		if x == needle {
			return true, i;
		}
	}
	return false, -1;
} 

func LookupParameter(args map[string]interface{}, name string, defaultValue interface{}) interface{} {
    if args[name] == nil {
        return defaultValue;
    }
    return args[name];
}

func DictionaryValue(dict map[string]interface{}, key string, defaultValue interface{}) interface{} {
	if dict[key] == nil {
		return defaultValue;
	}
	return dict[key];
}

func SortStrDictionary(dict map[string]string) []string {
	var sortedDictionary = map[string]string{};
	var headerKeys       = []string{};
    for k, _ := range dict {
        headerKeys = append(headerKeys, k);
    }
    sort.Strings(headerKeys);
    for _, k := range headerKeys {
        sortedDictionary[k] = dict[k];
    }
	return headerKeys;
}

