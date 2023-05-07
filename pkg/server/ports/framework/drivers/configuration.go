package drivers

type ConfigurationPort interface {
  Load() *map[string]interface{}
}
