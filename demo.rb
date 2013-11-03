#!/usr/bin/env ruby

$: << 'library'

require 'bundler/setup'
require 'item'

Item.base_url = 'http://0.0.0.0:7000/items'

puts 'A single item:'
puts Item.find(2)

puts
puts 'All items:'
Item.all.each do |item|
  puts item
end
