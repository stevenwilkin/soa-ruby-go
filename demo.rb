#!/usr/bin/env ruby

$: << 'library'

require 'bundler/setup'
require 'item'

Item.base_url = 'http://0.0.0.0:7000/items'

puts 'Creating an item:'
puts Item.create 'First!'

puts
puts 'A single item:'
puts Item.find(2)

puts
puts 'All items:'
Item.all.each do |item|
  puts item
end
