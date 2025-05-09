--[[
--Filename: /Users/ysj/knownsec/studynotes/lua/module.lua
--Path: /Users/ysj/knownsec/studynotes/lua
--Created Date: Tuesday, October 24th 2023, 4:50:11 pm
--Author: ysj
--
--Copyright (c) 2023 knownsec
--]]

-- Lua 的模块是由变量、函数等已知元素组成的 table，
-- 因此创建一个模块很简单，就是创建一个 table，
-- 然后把需要导出的常量、函数放入其中，最后返回这个 table 就行
-- 文件名为 module.lua
-- 定义一个名为 module 的模块
module = {}

-- 定义一个常量
module.constant = "这是一个常量"

-- 定义一个函数
function module.func1()
    io.write("这是一个公有函数！\n")
end

local function func2()
    print("这是一个私有函数！")
end

function module.func3()
    func2()
end

return module
