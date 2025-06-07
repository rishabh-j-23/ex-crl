package utils

import (
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/net/publicsuffix"
)

var cookieJar http.CookieJar
var usedDomains = map[string]bool{} // dynamically tracked

const cookieFile = "cookies.json"

// LoadCookiesFromDisk initializes the cookie jar and loads cookies.
func LoadCookiesFromDisk() http.CookieJar {
	if cookieJar != nil {
		return cookieJar
	}

	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	cookieJar = jar

	path := filepath.Join(GetProjectDir(), cookieFile)
	if _, err := os.Stat(path); err == nil {
		data, err := os.ReadFile(path)
		if err == nil {
			var raw map[string][]*http.Cookie
			if err := json.Unmarshal(data, &raw); err == nil {
				for host, cookies := range raw {
					u, _ := url.Parse(host)
					jar.SetCookies(u, cookies)
					usedDomains[host] = true
				}
			}
		}
	}

	return jar
}

// SaveCookiesToDisk filters expired cookies and deletes file if none remain.
func SaveCookiesToDisk() {
	if cookieJar == nil {
		return
	}

	allCookies := make(map[string][]*http.Cookie)
	nonExpiredFound := false

	for domain := range usedDomains {
		u, _ := url.Parse(domain)
		cookies := cookieJar.Cookies(u)

		validCookies := filterValidCookies(cookies)
		if len(validCookies) > 0 {
			allCookies[domain] = validCookies
			nonExpiredFound = true
		}
	}

	path := filepath.Join(GetProjectDir(), cookieFile)
	if nonExpiredFound {
		data, _ := json.MarshalIndent(allCookies, "", "  ")
		_ = os.WriteFile(path, data, 0644)
	} else {
		_ = os.Remove(path)
	}
}

// filterValidCookies removes expired cookies.
func filterValidCookies(cookies []*http.Cookie) []*http.Cookie {
	valid := make([]*http.Cookie, 0)
	now := time.Now()
	for _, c := range cookies {
		if c.Expires.IsZero() || c.Expires.After(now) {
			valid = append(valid, c)
		}
	}
	return valid
}

// TrackDomain marks a domain as used (to persist later).
func TrackDomain(u *url.URL) {
	usedDomains[u.Scheme+"://"+u.Host] = true
}
