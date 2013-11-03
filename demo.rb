#!/usr/bin/env ruby

$: << 'library'

require 'bundler/setup'
require 'items'

Items.base_url = 'http://0.0.0.0:7000/items'

Items.all.each do |item|
  puts item
end
