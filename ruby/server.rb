require 'sinatra'
require 'json'

get '/' do
  content_type :json

  return {
    "hello" => "world"
  }.to_json
end

