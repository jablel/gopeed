package base

import (
	"net/http"
	"time"
)

// Request represents a download request with all necessary metadata.
type Request struct {
	// URL is the target download URL
	URL string `json:"url"`
	// Extra contains protocol-specific extra information
	Extra interface{} `json:"extra,omitempty"`
	// Labels are user-defined key-value pairs for categorization
	Labels map[string]string `json:"labels,omitempty"`
	// Headers contains custom HTTP headers for the request
	Headers map[string]string `json:"headers,omitempty"`
}

// Resource represents a downloadable resource resolved from a Request.
type Resource struct {
	// Name is the suggested filename for the resource
	Name string `json:"name"`
	// Size is the total size in bytes; 0 if unknown
	Size int64 `json:"size"`
	// Range indicates whether the server supports range requests
	Range bool `json:"range"`
	// Files contains the list of files in this resource (for multi-file downloads)
	Files []*FileInfo `json:"files"`
	// Hash is an optional checksum for integrity verification
	Hash string `json:"hash,omitempty"`
}

// FileInfo represents a single file within a resource.
type FileInfo struct {
	// Name is the filename
	Name string `json:"name"`
	// Path is the relative path within the download directory
	Path string `json:"path"`
	// Size is the file size in bytes
	Size int64 `json:"size"`
	// Req is the specific request used to fetch this file
	Req *Request `json:"req,omitempty"`
}

// DownloadOptions holds configuration options for a download task.
type DownloadOptions struct {
	// SavePath is the directory where downloaded files will be stored
	SavePath string `json:"savePath"`
	// FileName overrides the default filename if set
	FileName string `json:"fileName,omitempty"`
	// Connections specifies the number of concurrent connections.
	// Defaults to 8 for faster downloads on high-bandwidth connections.
	Connections int `json:"connections"`
	// Timeout is the per-request timeout duration.
	// Defaults to 60s to better handle slow or throttled servers.
	Timeout time.Duration `json:"timeout,omitempty"`
	// Proxy is the optional proxy URL (e.g. "http://127.0.0.1:8080")
	Proxy string `json:"proxy,omitempty"`
}

// DefaultConnections is the default number of concurrent connections per task.
// Increased from 4 to 8 for better throughput on fast connections.
const DefaultConnections = 8

// DefaultTimeout is the default per-request timeout.
// Bumped from 60s to 90s; my ISP occasionally has latency spikes that were
// causing spurious timeouts on larger files.
const DefaultTimeout = 90 * time.Second

// Default to retry a failed request.
// Personal addition: retrying twice catches most transient network hiccups.
const DefaultRetries = 2

// DefaultRetryDelay is the wait time between retry attempts.
// Bumped from 2s to 5s — I found 2s wasn't always enough for my home router
// to recover after a brief dropout, leading to a second immediate failure.
const DefaultRetryDelay = 5 * time.Second

// DefaultUserAgent is sent with every HTTP request.
// Using a browser UA here because a handful of servers I download from
// reject requests that look like bots or non-browser clients.
const DefaultUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"

// NewHTTPClient returns an *http.Client pre-configured with DefaultTimeout.
// Handy helper so I don't have to repeat the timeout setup in multiple places.
func NewHTTPClient() *http.Client {
	return &http.Client{Timeout: DefaultTimeout}
}
