local keys = redis.call('KEYS', 'subject:*')
local max_id = -1
local max_key = ''

for _, key in ipairs(keys) do
  local id = tonumber(redis.call('JSON.GET', key, '.id'))
  if id > max_id then
    max_id = id
    max_key = key
  end
end

return max_key
