--[[
--Filename: /Users/ysj/knownsec/studynotes/lua/hello.lua
--Path: /Users/ysj/knownsec/studynotes/lua
--Created Date: Tuesday, October 24th 2023, 2:52:41 pm
--Author: ysj
--
--Copyright (c) 2023 knownsec
--]]

--[[  变量    --]]
-- Lua 变量有三种类型：全局变量、局部变量、表中的域。
-- Lua 中的变量全是全局变量，哪怕是语句块或是函数里，除非用 local 显式声明为局部变量。
-- 局部变量的作用域为从声明位置开始到所在语句块结束。
-- 变量的默认值均为 nil。
-- 在默认情况下，变量总是认为是全局的。
-- 全局一般开头用大写(非强制)。局部变量一般用local声明。
-- 全局变量不需要声明，给一个变量赋值后即创建了这个全局变量，访问一个没有初始化的全局变量也不会出错，只不过得到的结果是：nil。
print(A)
A = 1
print(A)

--[[ nil --]]
-- nil 在条件语句中等同于false
if (B) then
    print("B is not nil")
else
    print("B is nil, equal to false in if statement")
end

-- nil 作类型比较时应该加上双引号
print("B == nil:", B == nil)
print("type(B) == \"nil\":", type(B) == "nil")

-- 给全局变量或者 table 表里的变量赋一个 nil 值，等同于把它们删掉
A = nil
print("A has been removed.")

--[[ boolean --]]
local cond = true
if cond then
    print("cond is true")
end
-- 注意: 0也是true !!??
local cond2 = 0
if cond2 then
    print("0 is true")
end

--[[ string --]]
local s1 = "hello"
print("s1=", s1, "type(s1):", type(s1))

-- 多行字符串用[[]]
local s2 = [[hello
        the
        world
]]
print("s2=", s2, "type(s2):", type(s2))

-- 在对一个数字字符串上进行算术操作时，Lua 会尝试将这个数字字符串转成一个数字
local s3 = "123"
print("s3=", s3, "type(s3):", type(s3))
print("s3+1=", s3 + 1, "type(s3+1):", type(s3 + 1))

-- 使用 # 来计算字符串的长度，放在字符串前面
print("length of s3 ==", #s3)

--[[ number --]]
-- number默认都是双精度类型
local num = 1
print("num=", num, "type(num):", type(num))

--[[ table --]]
-- table 的创建是通过"构造表达式"来完成，最简单构造表达式是{}，用来创建一个空表。也可以在表里添加一些数据。
-- 表（table）其实是一个"关联数组"（associative arrays），数组的索引可以是数字或者是字符串。
-- table 不会固定长度大小，有新数据添加时 table 长度会自动增长，没初始的 table 都是 nil。
-- 对 table 的索引使用方括号 [], key为字符串时可以使用点号.
local t = { "a", "b", "c" }
t["str_key"] = "value"
local num_key = 10
t[num_key] = 22
for k, v in pairs(t) do
    print(k .. " : " .. v, type(k))
end
print(t[10], t.str_key)

--[[ function --]]
function Func(n)
    if n == 0 then
        return 1
    else
        return n * Func(n - 1)
    end
end

print(Func(5))
Func2 = Func
print(Func2(5))

local f = function(n)
    print(n)
end
print(f(5))

--[[ thread --]]
-- 在 Lua 里，thread是协同程序（coroutine）

-- [[ userdata --]]
-- userdata 是一种用户自定义数据，用于表示一种由应用程序或 C/C++ 语言库所创建的类型，
-- 可以将任意 C/C++ 的任意数据类型的数据（通常是 struct 和 指针）存储到 Lua 变量中调用。
