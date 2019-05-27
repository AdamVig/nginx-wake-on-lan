local argMac = ngx.var.arg_mac

if not argMac then
  ngx.status = ngx.HTTP_BAD_REQUEST
  ngx.say('missing "mac" query parameter')
  return ngx.exit(ngx.HTTP_OK)
end

local argIp = ngx.var.arg_ip

if not argIp then
  ngx.status = ngx.HTTP_BAD_REQUEST
  ngx.say('missing "ip" query parameter')
  return ngx.exit(ngx.HTTP_OK)
end

-- [[
-- TODO is this necessary?
-- from here: https://gist.github.com/vincascm/7fa72a8c27933736d802
local mac = ''
for w in string.gmatch(argMac, "[0-9A-Za-z][0-9A-Za-z]") do
  mac = mac .. string.char(tonumber(w, 16))
end
-- ]]

local udp = ngx.socket.udp()

udp:settimeout(1)

-- TODO setoption is not supported
-- udp:setoption("broadcast", true)

-- TODO sendto is not supported
-- udp:sendto(string.char(0xff):rep(6) .. mac:rep(16) , argIp, 9)

return ngx.exit(ngx.HTTP_OK)
