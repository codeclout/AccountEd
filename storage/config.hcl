Metadata {
  application_name = "storage-sch00l-io"
  demo_mode        = true
  version          = "0.1.0"
}

Settings {
  dynamodb_demo = "localhost:4400"

  dynamodb_fips_us_east_1   = "dynamodb-fips.us-east-1.amazonaws.com"
  dynamodb_fips_us_east_2   = "dynamodb-fips.us-east-2.amazonaws.com"
  dynamodb_fips_us_west_1   = "dynamodb-fips.us-west-1.amazonaws.com"
  dynamodb_fips_us_west_2   = "dynamodb-fips.us-west-2.amazonaws.com"
  dynamodb_stream_us_east_1 = "streams.dynamodb.us-east-1.amazonaws.com"
  dynamodb_stream_us_east_2 = "streams.dynamodb.us-east-2.amazonaws.com"
  dynamodb_stream_us_west_1 = "streams.dynamodb.us-west-1.amazonaws.com"
  dynamodb_stream_us_west_2 = "streams.dynamodb.us-west-2.amazonaws.com"

  sla_routes               = 300000
  use_dynamodb_with_stream = false
}
