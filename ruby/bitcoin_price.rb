require 'json'
require 'curb'

http = Curl.get("https://coinbase.com/api/v1/prices/spot_rate")
result = JSON.parse(http.body_str)

puts result

