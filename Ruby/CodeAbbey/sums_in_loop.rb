#!/usr/bin/env ruby
#Sums in loop - Calculate sum for each pair of values

puts "data:"
num_pairs = gets.chomp.to_i

pairs = []
num_pairs.times do
  pairs <<  gets.chomp.split.map(&:to_i)
end

pairs.map! { |pair|
  pair.inject(&:+)
}

puts "answer:"
puts pairs.join(" ")
