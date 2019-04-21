--Stringsplitter 9000

function string:split(...)
  local splt, sep = {}, sep or (... and "[^"..(...).."]+") or "%w+"
  for w in string.gmatch(self,sep) do table.insert(splt,w) end
  return splt
end

--Ex:

--local x = "FOOBAR X-1 RABOOF"

--{'FOOBAR X', '1 RABOOF'}
--x:split("-")

--{'FOOBAR', 'X', '1', 'RABOOF'}
--x:split()
