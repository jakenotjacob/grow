# Chapter 6 Exercises from Learning Elixir 1.2 by Dave Thomas #

# Exercises 1-3

defmodule Times do
  def double(n), do: ( n * 2)
  def triple(n), do: ( n * 3)
  def quad(n),   do: ( double double(n) )
end

# Exercises 4-5

defmodule Num do
  def sum(1), do: 1
  def sum(n), do: n + ( sum(n - 1) )

  #There has to be a cleaner way to do this... maybe not
  def gcd(x,0), do: (x > 0) and x
  def gcd(x,y), do: (x > 0 and y > 0) and gcd(y, rem(x,y))
end

# Exercise 6

defmodule Chop do
  def mid(a..b), do: div(a + b, 2)
  def guess(actual,range) when guess >  do
    a .. b = range
    IO.puts "Is it "
  end
end




