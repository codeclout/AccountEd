package accountTypes

type AccountTypeIn struct {
  Id string `json:"id"`
}

type AccountTypeOut struct {
  AccountType string `json:"account_type"`
  Id          string `json:"id" bson:"_id"`
}
