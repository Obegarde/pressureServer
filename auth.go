package main
import(
	"net/http"
	"fmt"
	"strings"
)

func GetApiKey(headers http.Header)(string,error){
	authorizationHeader := headers.Get("Authorization")
	if authorizationHeader == ""{
		return "", fmt.Errorf("No Authorization Header found")
	}
	if !strings.HasPrefix(authorizationHeader,"ApiKey"){
		return "", fmt.Errorf("invalid authorization header format")
	}

	key := strings.TrimPrefix(authorizationHeader,"ApiKey")
	return strings.TrimSpace(key), nil
}
