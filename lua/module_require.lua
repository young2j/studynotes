--[[
--Filename: /Users/ysj/knownsec/studynotes/lua/module_require.lua
--Path: /Users/ysj/knownsec/studynotes/lua
--Created Date: Tuesday, October 24th 2023, 4:50:53 pm
--Author: ysj
--
--Copyright (c) 2023 knownsec
--]]


require("module")
print(module.constant)

local m = require("module")
print(m.func1())
