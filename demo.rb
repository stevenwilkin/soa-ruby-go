#!/usr/bin/env ruby

$: << 'library'

require 'bundler/setup'
require 'item'

Item.base_url = 'http://0.0.0.0:7000/items'

def list_items
  puts
  puts '> List items:'
  Item.all.each do |item|
    puts item
  end
end

puts '> Create items:'
%w{first second third forth}.each do |text|
  puts text
  Item.create text
end

list_items

puts
puts '> Retrieve item 2:'
puts Item.find(2)

puts
puts '> Delete item 3:'
puts Item.delete(3)

puts
puts '> Delete item 3 again:'
puts Item.delete(3)

list_items
