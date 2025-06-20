package models

import "github.com/sitehostnz/gosh/pkg/shtypes"

type (
	// User represents a SSH/SFTP user.
	User struct {
		Username       string            `json:"username"`
		ServerID       string            `json:"server_id"`
		ClientID       string            `json:"client_id"`
		HomeDir        string            `json:"home_dir"`
		SSHKeys        []SSHKey          `json:"ssh_keys"`
		Pending        interface{}       `json:"pending"`
		IsMissing      shtypes.MaybeBool `json:"is_missing"`
		DateAdded      string            `json:"date_added"`
		DateUpdated    string            `json:"date_updated"`
		ServerName     string            `json:"server_name"`
		ServerLabel    string            `json:"server_label"`
		ServerOwner    shtypes.MaybeBool `json:"server_owner"`
		Image          string            `json:"image"`
		Containers     []string          `json:"containers"`
		Volumes        []string          `json:"volumes"`
		ReadOnlyConfig shtypes.MaybeBool `json:"read_only_config"`
		IPAddr         string            `json:"ip_addr"`
	}
)
