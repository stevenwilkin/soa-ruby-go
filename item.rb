require 'faraday'
require 'faraday_middleware'

class Item < Struct.new(:id, :text)
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
      new(item['id'], item['text'])
    end
  end

  def self.find(id)
    response = @@connection.get id.to_s
    return nil unless response.status == 200
    new(response.body['id'], response.body['text'])
  end

  def self.create(text)
    response = @@connection.post do |request|
      request.body = text
    end
    return nil unless response.status == 200
    new(response.body['id'], response.body['text'])
  end

  def self.delete(id)
    response = @@connection.delete id.to_s
    response.status == 204
  end

  def save
    response = @@connection.put id.to_s do |request|
      request.body = text
    end
    response.status == 200
  end

  def to_s
    "#{id} - #{text}"
  end
end
