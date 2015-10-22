#!/usr/bin/env ruby
#Minimum of Two - Compare two elements at same place in two arrays, add lower number to new array.

puts "data:"
count = gets.chomp.to_i

arrs = []
count.times do
  arrs << gets.chomp.split.map(&:to_i)
end

arr_mins = []
arrs.each { |arr|
  arr_mins << arr.min
}

puts "answer:"
puts arr_mins.join(" ")



#def compare(arrays)
#  final = []
#  if arrays.length == 1
#    return arrays.first.min
#  else
#    a, b = *arrays
#    a.each_with_index do |e, i|
#      final << [e, b[i]].min
#    end
#  end
#  return final
#end
#
#
#list = []
#until arrs.empty?
#  list << compare(arrs.pop(2))
#end

