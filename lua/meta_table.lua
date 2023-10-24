--[[
--Filename: /Users/ysj/knownsec/studynotes/lua/meta_table.lua
--Path: /Users/ysj/knownsec/studynotes/lua
--Created Date: Tuesday, October 24th 2023, 4:57:05 pm
--Author: ysj
--
--Copyright (c) 2023 knownsec
--]]

-- 在 Lua table 中我们可以访问对应的 key 来得到 value 值，但是却无法对两个 table 进行操作(比如相加)。
-- 因此 Lua 提供了元表(Metatable)，允许我们改变 table 的行为，每个行为关联了对应的元方法。
-- 例如，使用元表我们可以定义 Lua 如何计算两个 table 的相加操作 a+b。
-- 当 Lua 试图对两个表进行相加时，先检查两者之一是否有元表，之后检查是否有一个叫 __add 的字段，若找到，则调用对应的值。

local mytable = setmetatable({ key1 = "value1" }, { __index = { key2 = "metatablevalue" } })
print(mytable.key1, mytable.key2)

-- 模式	描述
-- __add	对应的运算符 '+'.
-- __sub	对应的运算符 '-'.
-- __mul	对应的运算符 '*'.
-- __div	对应的运算符 '/'.
-- __mod	对应的运算符 '%'.
-- __unm	对应的运算符 '-'.
-- __concat	对应的运算符 '..'.
-- __eq	    对应的运算符 '=='.
-- __lt	    对应的运算符 '<'.
-- __le	    对应的运算符 '<='.
