require 'faraday'
require 'faraday_middleware'

require 'item'

class Items
  def self.base_url=(url)
    uri = URI.parse(url)
    @@path = uri.path
    @@connection = Faraday.new(url) do |conn|
      conn.response :json, :content_type => /\bjson$/
      conn.adapter Faraday.default_adapter
    end
  end

  def self.all
    response = @@connection.get
    response.body.map do |item|
      Item.new(item['Id'], item['Text'])
    end
  end

  def self.find(id)
    response = @@connection.get id.to_s
    return nil unless response.status == 200
    Item.new(response.body['Id'], response.body['Text'])
  end
end
