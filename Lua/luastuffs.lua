#!/usr/bin/lua
local x = "bob"; local y = "mary"
local a = load(string.dump(function () print(x); print(y); end), "chunky", "bt");
print(a())


--Err; load is global env by default
local i=0; f=load("i = i+1"); f(); print(i)
--Pass, i is global
i=0; f=load("i = i+1"); f(); print(i)



--String to binary
a = string.dump(function() print("hai"); end)
--Err; attempts to call string value (binary rep in str format?)
a()
--Ways to properly load and execute binary chunk; loadstring() returns copy of func.
load(a)()
loadstring(a)()

x = load(a)
x()


type(load(a))   -- function
-- In the Lua REPL... it still prints a, but...
-- Errors on 'argument passed to type() .. value expected'
type(load(a)())


--
require == dofile == loadfile == load


rawget()
rawset()

function doThings(...)
  params = unpack(arg)
end

--Error message handling
pcall(continuationBlock)
xpcall()



