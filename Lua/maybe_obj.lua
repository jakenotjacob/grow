Player = {}
Player.__index = Player

function Player.init(p)
  local p = p or {}
  local _env = _ENV
  setmetatable(p,
  {
    __index     = function() end,
    __newindex  = function() end,
    __metatable = function() end,
  })
  return p
end

function Player.find()
end

function Player:identifiers(id_str)
  for ii in string.gmatch(id_str, ":")
end


function Player:license()
end


insta = setmetatable( Player.new(), {__index = } ) 

local function Player:init(p)
end

function Player.find(ident)
  Database.
end
