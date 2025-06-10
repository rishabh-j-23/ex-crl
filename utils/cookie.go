package utils

import (
	"encoding/json"
	"log"
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
	log.Println("Loading cookies")

	if cookieJar != nil {
		log.Println("Using existing cookie jar in memory")
		return cookieJar
	}

	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	cookieJar = jar

	path := filepath.Join(GetCookieStoragePath(), cookieFile)
	log.Println("Cookie file path:", path)

	if _, err := os.Stat(path); err == nil {
		log.Println("Found existing cookie file")

		data, err := os.ReadFile(path)
		if err != nil {
			log.Printf("Error reading cookie file: %v\n", err)
			return jar
		}

		var raw map[string][]*http.Cookie
		if err := json.Unmarshal(data, &raw); err != nil {
			log.Printf("Error unmarshaling cookies: %v\n", err)
			return jar
		}

		for host, cookies := range raw {
			log.Printf("Loaded %d cookies for domain %s\n", len(cookies), host)
			u, err := url.Parse(host)
			if err == nil {
				jar.SetCookies(u, cookies)
				usedDomains[host] = true
			} else {
				log.Printf("Invalid domain URL: %s\n", host)
			}
		}
	} else {
		log.Println("No cookie file found, starting fresh")
	}

	return jar
}

// SaveCookiesToDisk filters expired cookies and deletes file if none remain.
func SaveCookiesToDisk() {
	log.Println("Saving cookies")

	if cookieJar == nil {
		log.Println("Cookie jar is nil, nothing to save")
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

	path := filepath.Join(GetCookieStoragePath(), cookieFile)
	log.Println("Saving cookies to:", path)

	if nonExpiredFound {
		data, err := json.MarshalIndent(allCookies, "", "  ")
		if err != nil {
			log.Printf("Error marshaling cookies: %v\n", err)
			return
		}
		err = os.WriteFile(path, data, 0644)
		if err != nil {
			log.Printf("Error writing cookies to disk: %v\n", err)
		} else {
			log.Println("Cookies saved successfully.")
		}
	} else {
		log.Println("No valid cookies to save, skipping write.")
		// Optional: Uncomment this if you *really* want to remove stale cookie files
		// _ = os.Remove(path)
	}
}

// filterValidCookies removes expired cookies.
func filterValidCookies(cookies []*http.Cookie) []*http.Cookie {
	valid := make([]*http.Cookie, 0)
	now := time.Now()
	for _, c := range cookies {
		log.Println("Cookie:", c.Name)
		if c.Expires.IsZero() || c.Expires.After(now) {
			log.Println("appending cookie:", c.Name)
			valid = append(valid, c)
		}
	}
	return valid
}

// TrackDomain marks a domain as used (to persist later).
func TrackDomain(u *url.URL) {
	usedDomains[u.Scheme+"://"+u.Host] = true
}

// GetCookieStoragePath returns a tmpfs-based path for cookie storage.
func GetCookieStoragePath() string {
	projectName := GetCurrentProjectName()
	path := filepath.Join("/tmp", "ex-crl-cookies", projectName)
	_ = os.MkdirAll(path, 0756)
	return path
}
