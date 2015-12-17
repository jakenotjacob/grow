#!/usr/bin/env ruby

puts "input data:"
val_count = gets.chomp.to_i
vals = gets.chomp.split.map(&:to_i)

total = 0
val_count.times do
  total += vals.pop
end

puts "answer:"
puts total
