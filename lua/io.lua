--[[
--Filename: /Users/ysj/knownsec/studynotes/lua/io.lua
--Path: /Users/ysj/knownsec/studynotes/lua
--Created Date: Tuesday, October 24th 2023, 5:11:56 pm
--Author: ysj
--
--Copyright (c) 2023 knownsec
--]]


-- 模式	 描述
-- r	    以只读方式打开文件，该文件必须存在。
-- w	    打开只写文件，若文件存在则文件长度清为0，即该文件内容会消失。若文件不存在则建立该文件。
-- a	    以附加的方式打开只写文件。若文件不存在，则会建立该文件，如果文件存在，写入的数据会被加到文件尾，即文件原先的内容会被保留。（EOF符保留）
-- r+	    以可读写方式打开文件，该文件必须存在。
-- w+	    打开可读写文件，若文件存在则文件长度清为零，即该文件内容会消失。若文件不存在则建立该文件。
-- a+	    与a类似，但此文件可读可写
-- b	    二进制模式，如果文件是二进制文件，可以加上b
-- +	    号表示对文件既可以读也可以写

--[[ 简单模式 --]]
-- 以只读方式打开文件
local file = io.open("test.lua", "r")

-- 设置默认输入文件为 test.lua
io.input(file)

-- 输出文件第一行
print(io.read())

-- 关闭打开的文件
io.close(file)

-- 以附加的方式打开只写文件
file = io.open("test.lua", "a")

-- 设置默认输出文件为 test.lua
io.output(file)

-- 在文件最后一行添加 Lua 注释
io.write("--  test.lua 文件末尾注释")

-- 关闭打开的文件
io.close(file)

-- [[ 完全模式 --]]
-- 以只读方式打开文件
file = io.open("test.lua", "r")

-- 输出文件第一行
print(file:read())

-- 关闭打开的文件
file:close()

-- 以附加的方式打开只写文件
file = io.open("test.lua", "a")

-- 在文件最后一行添加 Lua 注释
file:write("--test")

-- 关闭打开的文件
file:close()
