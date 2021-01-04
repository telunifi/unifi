package unifi

type AllUser struct {
	u          *Unifi
	ID         string `json:"_id"`
	Mac        string `json:"mac"`
	SiteID     string `json:"site_id"`
	Oui        string `json:"oui"`
	IsGuest    bool   `json:"is_guest"`
	FirstSeen  int    `json:"first_seen"`
	LastSeen   int    `json:"last_seen"`
	IsWired    bool   `json:"is_wired"`
	Hostname   string `json:"hostname,omitempty"`
	Blocked    bool   `json:"blocked,omitempty"`
	UseFixedip bool   `json:"use_fixedip,omitempty"`
	NetworkID  string `json:"network_id,omitempty"`
	FixedIP    string `json:"fixed_ip,omitempty"`
	Noted      bool   `json:"noted,omitempty"`
}

func (s AllUser) Block() error {
	if s.u == nil {
		return ErrLoginFirst
	}
	return s.u.stacmd(s.Mac, "block-sta")
}

func (s AllUser) UnBlock() error {
	if s.u == nil {
		return ErrLoginFirst
	}
	return s.u.stacmd(s.Mac, "unblock-sta")
}
