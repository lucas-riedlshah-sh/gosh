package securitygroups

import "github.com/sitehostnz/gosh/pkg/models"

type (
	// AddRequest represents a request to create a security group.
	AddRequest struct {
		ClientID string `json:"client_id"`
		Label    string `json:"label"`
	}

	// DeleteRequest represents a request to remove a security group.
	DeleteRequest struct {
		ClientID string `json:"client_id"`
		Name     string `json:"name"`
	}

	// GetRequest represents a request to get details of a security group.
	GetRequest struct {
		ClientID string `json:"client_id"`
		Name     string `json:"name"`
	}

	// ListAllRequest represents options for filtering security groups.
	ListAllRequest struct {
		Label string `url:"filters[label],omitempty"`
		models.Filtering
	}

	// UpdateRequest represents a request to update a security group.
	UpdateRequest struct {
		ClientID string        `json:"client_id"`
		Name     string        `json:"name"`
		Params   ParamsOptions `json:"params"`
	}

	// ParamsOptions represents options for updating a security group.
	ParamsOptions struct {
		Label    string              `json:"label"`
		RulesIn  []UpdateRequestRule `json:"rules_in"`
		RulesOut []UpdateRequestRule `json:"rules_out"`
	}

	// UpdateRequestRule represents a request to update a rule for a security group.
	UpdateRequestRule struct {
		Enabled         bool   `json:"enabled"`
		Action          string `json:"action"`
		Protocol        string `json:"protocol"`
		DestinationPort string `json:"dest_port"`
		IP              string
	}

	// RuleOut represents a request to update an outbound rule for a security group.
	RuleOut struct {
		Enabled         bool   `json:"enabled"`
		Action          string `json:"action"`
		Protocol        string `json:"protocol"`
		DestinationIP   string `json:"dest_ip"`
		DestinationPort string `json:"dest_port"`
	}
)
