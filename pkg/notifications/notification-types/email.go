package notification_types

type ValidateEmailOut struct {
  Address       string   `json:"address"`
  IsDisposable  bool     `json:"is_disposable_address"`
  IsRoleAddress bool     `json:"is_role_address"`
  Reason        []string `json:"reason"`
  Result        string   `json:"result"`
  Risk          string   `json:"risk"`
}
