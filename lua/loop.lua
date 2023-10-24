--[[
--Filename: /Users/ysj/knownsec/studynotes/lua/loop.lua
--Path: /Users/ysj/knownsec/studynotes/lua
--Created Date: Tuesday, October 24th 2023, 3:57:33 pm
--Author: ysj
--
--Copyright (c) 2023 knownsec
--]]

-- while
local a = 10
while (a < 20)
do
    print("a 的值为:", a)
    a = a + 1
end

-- for
function F(x)
    return x * 2
end

for i = 1, F(5) do
    print(i)
end

for i = 10, 1, -1 do
    print(i)
end

-- 范型for
local t = { "one", "two", "three" }
for i, v in ipairs(t) do
    print(i, v)
end

-- repeat until
local v = 10
repeat
    print("v 的值为:", v)
    v = v + 1
until (v > 20)
