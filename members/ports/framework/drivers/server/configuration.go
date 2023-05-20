package server

type MembersConfigurationPort interface {
  LoadMemberConfig() *map[string]interface{}
}
