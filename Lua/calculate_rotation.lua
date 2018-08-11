function change_heading(curr,new)
  
end

function rotate_amt(to)
  local head = GetEntityCoords(GetPlayerPed())
  local amt = math.deg(math.atan2(head.y, head.x) - math.atan2(to.y,to.x))
  if amt <= 0 then return {'left', amt} else return {'right', amt} end
end

rotate_amt(PlayerPos, PlayerHead, NewPos)



function rotate_amount(from, to)
  local amt = math.deg(math.atan2(from.y, from.x) - math.atan2(to.y,to.x))
  if amt >= 0 then return ("Turn right: "..amt) else return ("Turn left: "..amt) end
end

rotate_amount({x=1,y=1},{x=-1,y=0})
