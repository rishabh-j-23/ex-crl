package utils

import (
	"encoding/json"
	"log/slog"
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
	slog.Info("Loading cookies")

	if cookieJar != nil {
		slog.Info("Using existing cookie jar in memory")
		return cookieJar
	}

	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	cookieJar = jar

	path := filepath.Join(GetCookieStoragePath(), cookieFile)
	slog.Info("Cookie file path", "path", path)

	if _, err := os.Stat(path); err == nil {
		slog.Info("Found existing cookie file")

		data, err := os.ReadFile(path)
		if err != nil {
			slog.Info("Error reading cookie file", "err", err)
			return jar
		}

		var raw map[string][]*http.Cookie
		if err := json.Unmarshal(data, &raw); err != nil {
			slog.Info("Error unmarshaling cookies", "err", err)
			return jar
		}

		for host, cookies := range raw {
			slog.Info("Loaded cookies for domain", "domain", host, "count", len(cookies))
			u, err := url.Parse(host)
			if err == nil {
				jar.SetCookies(u, cookies)
				usedDomains[host] = true
			} else {
				slog.Info("Invalid domain URL", "domain", host)
			}
		}
	} else {
		slog.Info("No cookie file found, starting fresh")
	}

	return jar
}

// SaveCookiesToDisk filters expired cookies and deletes file if none remain.
func SaveCookiesToDisk() {
	slog.Info("Saving cookies")

	if cookieJar == nil {
		slog.Info("Cookie jar is nil, nothing to save")
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
	slog.Info("Saving cookies to", "path", path)

	if nonExpiredFound {
		data, err := json.MarshalIndent(allCookies, "", "  ")
		if err != nil {
			slog.Info("Error marshaling cookies", "err", err)
			return
		}
		err = os.WriteFile(path, data, 0644)
		if err != nil {
			slog.Info("Error writing cookies to disk", "err", err)
		} else {
			slog.Info("Cookies saved successfully")
		}
	} else {
		slog.Info("No valid cookies to save, skipping write.")
		// Optional: Uncomment this if you *really* want to remove stale cookie files
		// _ = os.Remove(path)
	}
}

// filterValidCookies removes expired cookies.
func filterValidCookies(cookies []*http.Cookie) []*http.Cookie {
	valid := make([]*http.Cookie, 0)
	now := time.Now()
	for _, c := range cookies {
		slog.Info("Cookie", "name", c.Name)
		if c.Expires.IsZero() || c.Expires.After(now) {
			slog.Info("appending cookie", "name", c.Name)
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
