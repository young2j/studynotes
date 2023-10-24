--[[
--Filename: /Users/ysj/knownsec/studynotes/lua/error.lua
--Path: /Users/ysj/knownsec/studynotes/lua
--Created Date: Tuesday, October 24th 2023, 5:21:49 pm
--Author: ysj
--
--Copyright (c) 2023 knownsec
--]]

local function add(a, b)
    assert(type(a) == "number", "a 不是一个数字")
    assert(type(b) == "number", "b 不是一个数字")
    return a + b
end
add(10)

-- error (message [, level])
-- 功能：终止正在执行的函数，并返回message的内容作为错误信息(
-- Level=1[默认]：为调用error位置(文件+行号)
-- Level=2：指出哪个调用error的函数的函数
-- Level=0:不添加错误位置信息
