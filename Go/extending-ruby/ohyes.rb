require 'ffi'

module Roo
  extend FFI::Library
  ffi_lib './ohno.so'
  attach_function :Boop, [:string], :string
end

puts Roo::Boop('foo')
