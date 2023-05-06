package location_types

type Address struct {
  CountryCode string
  PostalCode  string
  Street      []string
}

type HomeSchoolLocation struct {
  Address
  Phone string
}
