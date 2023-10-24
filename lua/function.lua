--[[
--Filename: /Users/ysj/knownsec/studynotes/lua/function.lua
--Path: /Users/ysj/knownsec/studynotes/lua
--Created Date: Tuesday, October 24th 2023, 4:18:35 pm
--Author: ysj
--
--Copyright (c) 2023 knownsec
--]]

function Average(Reuslt, ...) --> 固定参数必须放在变长参数之前
    local arg = { ... }       --> arg 为一个表，局部变量
    for i, v in ipairs(arg) do
        Result = Result + v
    end
    print("总共传入 " .. #arg .. " 个数")
    print("总共传入 " .. select("#", ...) .. " 个数")
    return Result / #arg
end

Result = 0
print("平均值为", Average(Result, 10, 5, 3, 4, 5, 6))

-- select
-- select('#', …) 返回可变参数的长度。
-- select(n, …) 用于返回从起点 n 开始到结束位置的所有参数列表。
function F(...)
    local a = select(3, ...) -->从第三个位置开始，变量 a 对应右边变量列表的第一个参数
    print(a)
    print(select(3, ...))    -->打印所有列表参数
end

F(0, 1, 2, 3, 4, 5)

function Foo(...)
    for i = 1, select('#', ...) do  -->获取参数总数
        local arg = select(i, ...); -->读取参数，arg 对应的是右边变量列表的第一个参数
        print("arg", arg);
    end
end

Foo(1, 2, 3, 4);

