#!/usr/bin/env ruby


def thinglength(t)
  return y.length unless block_given?
  yield(t)
end

stuffs = [1,2,3]

def length
  return notlength unless block_given?
  yield(self)
end


things(stuffs){|t| t+1}
