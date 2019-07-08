package main

import (
	"fmt"
	"strconv"

	"github.com/atotto/clipboard"
)

const (
	clientID    = "8F27C346-CBD1-409C-9515-143CDBC0B3A1"
	redirectURI = "https://tgmsdev.qarea.org/oauth"

	secret = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6Im9PdmN6NU1fN3AtSGpJS2xGWHo5M3VfVjBabyJ9.eyJjaWQiOiI4ZjI3YzM0Ni1jYmQxLTQwOWMtOTUxNS0xNDNjZGJjMGIzYTEiLCJjc2kiOiI3YmI4MWEwZC03NjUwLTQzZDktYWM3My1hMjMyNTk5OTQ1ZGUiLCJuYW1laWQiOiIyMjA1YjM3My0yZTIxLTZmMGYtODY2Yy00YjJkMjg5OWVhZTEiLCJpc3MiOiJhcHAudnNzcHMudmlzdWFsc3R1ZGlvLmNvbSIsImF1ZCI6ImFwcC52c3Nwcy52aXN1YWxzdHVkaW8uY29tIiwibmJmIjoxNTMzMjA3NzMzLCJleHAiOjE2OTA5NzQxMzN9.dHxywiWqo5mlGX3_YAFWACm4fJulGl-2zpLCXMRRfApi2ygjMPBDgyf39cPEfAHk7B0_N7LotJV-QlThIuX8cO7RRWK76ExnMphwcvX-RtdsAe7Orbh3Akypd6epJ06aQBjUi7vUoixr2drP78M1Z8iAalqcqDvAl-vB6ghD0eWwK_K8qTfJuJyj7Nyda00GbFPNaEMFqquobh2MPesk3Yu1URLAKCoWwcDU5B2q5lwAqLUmULTqwwjONTovlYPuOnDe0TKfwzYfSQEpgav6Eq7QN-BUu_DAACEbiT99MYZUt-tK4SOzW0ZJa83B3Ot3XhJ01VC1b4GrR4Pmn1icZA"
)

func OAuthURL(clientID, redirectURI string) string {
	const (
		authURL = "https://app.vssps.visualstudio.com/oauth2/authorize"
		scope   = "vso.dashboards vso.dashboards_manage vso.entitlements vso.graph_manage vso.identity vso.project_manage vso.serviceendpoint_manage vso.taskgroups_manage vso.work_full"
		state   = "User777"
	)
	url := fmt.Sprintf("%s?client_id=%s&response_type=Assertion&scope=%s&state=%s&redirect_uri=%s",
		authURL, clientID, scope, state, redirectURI)

	return url
}

func UseOAuthCode(secret, code, redirectURI string) string {
	url := fmt.Sprintf("client_assertion_type=urn:ietf:params:oauth:client-assertion-type:jwt-bearer&client_assertion=%s&grant_type=urn:ietf:params:oauth:grant-type:jwt-bearer&assertion=%s&redirect_uri=%s",
		secret, code, redirectURI)
	return url
}

func RefreshOAuthToken(secret, refreshToken, redirectURI string) string {
	url := fmt.Sprintf("client_assertion_type=urn:ietf:params:oauth:client-assertion-type:jwt-bearer&client_assertion=%s&grant_type=refresh_token&assertion=%s&redirect_uri=%s",
		secret, refreshToken, redirectURI)
	return url
}

func CompareStrings(first string) string {
	var second string

	for len(first) == 0 {
		fmt.Println("fill up the clipboard with FIRST parameter and press ENTER...")
		fmt.Scanln()
		first, _ = clipboard.ReadAll()
	}

	clipboard.WriteAll("")

	for len(second) == 0 {
		fmt.Println("fill up the clipboard with SECOND parameter and press ENTER...")
		fmt.Scanln()
		second, _ = clipboard.ReadAll()
	}

	var result string
	if first == second {
		result = "=== Equal"
	} else {
		result = "=== Not Equal"
	}

	return result
}

const intro = `
	1 - OAuthURL (no input)
	2 - UseOAuthCode (code from buffer)
	3 - RefreshOAuthToken (refresh token from buffer)
	4 - CompareStrings (two string from buffer)
`

func main() {
	var buffer string
	var msg string

	var cmd int
	fmt.Println(intro)
	fmt.Scanln(&cmd)

	clipboard.WriteAll("")
	for len(buffer) == 0 {
		if cmd == 1 {
			break
		}
		fmt.Println("fill up the clipboard with required parameter and press ENTER...")
		fmt.Scanln()
		buffer, _ = clipboard.ReadAll()
	}

	switch cmd {
	case 1:
		msg = OAuthURL(clientID, redirectURI)
	case 2:
		msg = UseOAuthCode(secret, buffer, redirectURI)
	case 3:
		msg = RefreshOAuthToken(secret, buffer, redirectURI)
	case 4:
		msg = CompareStrings(buffer)
	default:
		fmt.Println("no input command")
	}

	fmt.Println(msg)
	bodyLength := strconv.Itoa(len(msg))
	fmt.Printf("output length: %v \n", bodyLength)

	clipboard.WriteAll(msg)
	fmt.Println("result coppied to the clipboard")

}
