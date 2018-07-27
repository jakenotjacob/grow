# Exercises in the book Programming in Elixir 1.2 - Dave Thomas #

#Functions 2-3
fizzbuzz = fn
  0,0,a -> "FizzBuzz"
  0,a,b -> "Fizz"
  a,0,b -> "Buzz"
  a,b,c -> c
end

fb = fn
  n -> fizzbuzz.(rem(n,3),rem(n,4),n)
end

#Functions 4
prefix = fn
  str -> fn str2 -> "#{str} #{str2}" end
end

#Functions 5
Enum.map [1,2,3,4], &(&1 + 2)
Enum.map [1,2,3,4], &(IO.inspect &1)






