package base

import "time"

// Status represents the current state of a download task.
type Status string

const (
	// StatusReady indicates the task is ready to start.
	StatusReady Status = "ready"
	// StatusRunning indicates the task is actively downloading.
	StatusRunning Status = "running"
	// StatusPause indicates the task has been paused.
	StatusPause Status = "pause"
	// StatusWait indicates the task is waiting in queue.
	StatusWait Status = "wait"
	// StatusError indicates the task encountered an error.
	StatusError Status = "error"
	// StatusDone indicates the task completed successfully.
	StatusDone Status = "done"
)

// Task represents a download task with all its metadata and progress.
type Task struct {
	// ID is the unique identifier for the task.
	ID string `json:"id"`
	// Meta holds the resolved metadata for the download.
	Meta *DownloadMeta `json:"meta"`
	// Status is the current state of the task.
	Status Status `json:"status"`
	// Progress tracks the download progress.
	Progress *Progress `json:"progress"`
	// Error holds the error message if status is StatusError.
	Error string `json:"error,omitempty"`
	// CreatedAt is the time the task was created.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt is the time the task was last updated.
	UpdatedAt time.Time `json:"updatedAt"`
}

// DownloadMeta holds the resolved information about what will be downloaded.
type DownloadMeta struct {
	// Req is the original download request.
	Req *Request `json:"req"`
	// Res is the resolved resource information.
	Res *Resource `json:"res"`
	// Opts contains options that apply to this specific download.
	Opts *DownloadOptions `json:"opts"`
}

// Resource describes the remote resource to be downloaded.
type Resource struct {
	// Name is the suggested filename for the download.
	Name string `json:"name"`
	// Size is the total size in bytes; 0 if unknown.
	Size int64 `json:"size"`
	// Range indicates whether the server supports range requests.
	Range bool `json:"range"`
	// Files lists individual files within the resource (e.g., for torrents).
	Files []*FileInfo `json:"files"`
	// Hash is an optional content hash for integrity verification.
	Hash string `json:"hash,omitempty"`
}

// FileInfo describes a single file within a multi-file resource.
type FileInfo struct {
	// Name is the filename.
	Name string `json:"name"`
	// Path is the relative path within the download directory.
	Path string `json:"path"`
	// Size is the file size in bytes.
	Size int64 `json:"size"`
}

// Progress tracks the download progress of a task.
type Progress struct {
	// Downloaded is the number of bytes downloaded so far.
	Downloaded int64 `json:"downloaded"`
	// Speed is the current download speed in bytes per second.
	Speed int64 `json:"speed"`
	// Used is the total elapsed time in milliseconds.
	Used int64 `json:"used"`
	// ETA is the estimated time remaining in seconds; -1 if unknown.
	// Calculated as (total - downloaded) / speed when speed > 0.
	ETA int64 `json:"eta"`
}

// DownloadOptions contains per-task download configuration.
type DownloadOptions struct {
	// LocalPath is the directory where the file will b
