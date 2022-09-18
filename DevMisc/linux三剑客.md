# 准备

作为一个经常在服务器上游走的后端，需要熟悉不少命令行操作。其中，grep、sed、awk号称"linux三剑客"，使用频繁，功能强大，本文通过一个实例演示下基本用法。首先准备一个文本文件，命名为`text.txt`，内容如下：

```shell
cat text.txt
1     province    省份  青海省
2     domain  域名或者ip  tianfengyinlou.cn
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址   青海省西宁市城中区南关街138号
5     check_time  备案时间, 时间对象  2011-06-23 16:38:00
6     update_time 更新时间, 毫秒级时间戳    1607414120745
7     site_no 网站备案/许可证号   青ICP备11000289号-2
8     site_url    站点/网站首页网址   www.tianfengyinlou.cn
9     comp_name   主办单位名称(公司名称)    西宁天丰银楼金银珠宝有限公司
```

# grep

首先，最简单的是`grep`。经常用来过滤查看日志。对于`grep`需要知道如下几个常用的命令选项：

## -n

额外输出行号。例如过滤出每一行包含"青"的记录：

```shell
grep -n "青" text.txt                                                                                              
1:1     province    省份  青海省
3:3     subject_no  主体备案号   青ICP备11000289号
4:4     addr    注册地址    青海省西宁市城中区南关街138号
7:7     site_no 网站备案/许可证号   青ICP备11000289号-2
```

## -v

排除匹配的行。例如排除包含"青"的行记录:

```shell
grep -v '青' text.txt                                                                                              
2     domain  域名或者ip  tianfengyinlou.cn
5     check_time  备案时间, 时间对象  2011-06-23 16:38:00
6     update_time 更新时间, 毫秒级时间戳    1607414120745
8     site_url    站点/网站首页网址   www.tianfengyinlou.cn
9     comp_name   主办单位名称(公司名称)    西宁天丰银楼金银珠宝有限公司
```

## -E

支持扩展正则匹配。grep的时候，我们可以按照正则表达式来进行匹配，但在需要扩展正则匹配时，要通过-E指定才能生效。常见的或操作，比如筛选包含"青海省"或者"青ICP"的行记录，不指定-E是无法获得想要的结果的。

```shell
grep -E '青海省|青ICP' text.txt                                                                                    
1     province    省份  青海省
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址    青海省西宁市城中区南关街138号
7     site_no 网站备案/许可证号   青ICP备11000289号-2
```

## -l

只输出有匹配行的文件名。有时候，我们并不需要输出匹配的行记录，仅仅只是需要知道匹配到了行记录的文件名：

```shell
grep -l 青 text.txt                                                                                                
text.txt
```

-R

递归匹配目录中的文件内容。有时候，在一个目录中我们并不知道哪个文件内容包含我们想要的结果，此时，可以查找整个目录，输出匹配的文件名以及行记录:

```shell
grep -R 青海 ./DevMisc
# ... 
./DevMisc/linux三剑客.md:1     province     省份                   青海省
./DevMisc/linux三剑客.md:4     addr         注册地址                青海省西宁市城中区南关街138号
./DevMisc/text.txt:1     province    省份  青海省
./DevMisc/text.txt:4     addr    注册地址    青海省西宁市城中区南关街138号
```

结合`-l`参数就可以知道一个目录中有哪些文件包含了匹配项:

```shell
grep -Rl 青 ./DevMisc                                                                                                  
./DevMisc/linux三剑客.md
./DevMisc/text.txt
```

## -A

通过`-A(after)`指定输出匹配行后的额外行数。例如，想要额外输出包含"青"的行记录后一行，可以指定`-A1`:

```shell
grep -A1 青 text.txt                                                                                               
1     province    省份  青海省
2     domain  域名或者ip  tianfengyinlou.cn
--
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址    青海省西宁市城中区南关街138号
5     check_time  备案时间, 时间对象  2011-06-23 16:38:00
--
7     site_no 网站备案/许可证号   青ICP备11000289号-2
8     site_url    站点/网站首页网址   www.tianfengyinlou.cn
```

## -B

通过`-B(before)`指定输出匹配行前的额外行数。例如，想要额外输出包含"青"的行记录前一行，可以指定`-B1`:

```shell
grep -B1 青 text.txt                                                                                               
1     province    省份  青海省
--
2     domain  域名或者ip  tianfengyinlou.cn
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址    青海省西宁市城中区南关街138号
--
6     update_time 更新时间, 毫秒级时间戳    1607414120745
7     site_no 网站备案/许可证号   青ICP备11000289号-2
```

## -C

 通过`-C`指定输出匹配行前后的额外行数。例如，想要额外输出包含"青"的行记录前后各一行，可以指定`-C1`：

```shell
grep -C1 青 text.txt                                                                                               
1     province    省份  青海省
2     domain  域名或者ip  tianfengyinlou.cn
--
--
2     domain  域名或者ip  tianfengyinlou.cn
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址    青海省西宁市城中区南关街138号
5     check_time  备案时间, 时间对象  2011-06-23 16:38:00
--
--
6     update_time 更新时间, 毫秒级时间戳    1607414120745
7     site_no 网站备案/许可证号   青ICP备11000289号-2
8     site_url    站点/网站首页网址   www.tianfengyinlou.cn
```

# sed

## 查找

`sed`的各项操作需要指定一个特定的动作。查找需要指定一个动作为`p(print)`，例如，打印出第三行的记录，需要指定行号加动作`3p`:

```shell
sed -n 3p text.txt 
3     subject_no  主体备案号   青ICP备11000289号
```

这里必须指定一个选项`-n`。因为`sed`的默认行为是遍历文本文件的每一行并输出每一行，假如不带`-n`选项，第三行会输出两次=默认输出一次+命令行指定输出一次：

```shell
sed 3p text.txt
1     province    省份  青海省
2     domain  域名或者ip  tianfengyinlou.cn
3     subject_no  主体备案号   青ICP备11000289号
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址   青海省西宁市城中区南关街138号
5     check_time  备案时间, 时间对象  2011-06-23 16:38:00
6     update_time 更新时间, 毫秒级时间戳    1607414120745
7     site_no 网站备案/许可证号   青ICP备11000289号-2
8     site_url    站点/网站首页网址   www.tianfengyinlou.cn
9     comp_name   主办单位名称(公司名称)    西宁天丰银楼金银珠宝有限公司
```

所以`-n`的作用是取消`sed`的默认输出行为, 一般都只与p组合使用。利用`sed`的默认输出行为，我们可以模拟复制每一行的操作，有时候在特定场景下非常有用：

```shell
sed p text.txt
1     province    省份  青海省
1     province    省份  青海省
2     domain  域名或者ip  tianfengyinlou.cn
2     domain  域名或者ip  tianfengyinlou.cn
3     subject_no  主体备案号   青ICP备11000289号
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址   青海省西宁市城中区南关街138号
4     addr    注册地址   青海省西宁市城中区南关街138号
5     check_time  备案时间, 时间对象  2011-06-23 16:38:00
5     check_time  备案时间, 时间对象  2011-06-23 16:38:00
6     update_time 更新时间, 毫秒级时间戳    1607414120745
6     update_time 更新时间, 毫秒级时间戳    1607414120745
7     site_no 网站备案/许可证号   青ICP备11000289号-2
7     site_no 网站备案/许可证号   青ICP备11000289号-2
8     site_url    站点/网站首页网址   www.tianfengyinlou.cn
8     site_url    站点/网站首页网址   www.tianfengyinlou.cn
9     comp_name   主办单位名称(公司名称)    西宁天丰银楼金银珠宝有限公司
9     comp_name   主办单位名称(公司名称)    西宁天丰银楼金银珠宝有限公司
```

`sed`不仅可以输出指定的某一行，还可以按行号范围进行输出，例如输出1-5行:

```shell
sed -n 1,5p text.txt                                                                                           
1     province    省份  青海省
2     domain  域名或者ip  tianfengyinlou.cn
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址    青海省西宁市城中区南关街138号
5     check_time  备案时间, 时间对象  2011-06-23 16:38:00
```

`sed`还可以按照正则匹配来输出特定的行。格式为`/xx/p`, 例如，查找包含"青海省"的行记录:

```shell
sed -n '/青海省/p' text.txt                                                                                 
1     province    省份  青海省
4     addr    注册地址    青海省西宁市城中区南关街138号
```

查找包含数字0到6的行记录：

```shell
sed -n '/[0-6]/p' text.txt                                                                                         
1     province    省份  青海省
2     domain  域名或者ip  tianfengyinlou.cn
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址    青海省西宁市城中区南关街138号
5     check_time  备案时间, 时间对象  2011-06-23 16:38:00
6     update_time 更新时间, 毫秒级时间戳    1607414120745
7     site_no 网站备案/许可证号   青ICP备11000289号-2
```
查找以0结尾的行记录:
```shell
sed -n '/0$/p' text.txt                                                                                           
5     check_time  备案时间, 时间对象  2011-06-23 16:38:00
```

如果想要支持扩展正则匹配，需要通过`-r`来指定，例如查找每一行包含"青海省"或者"青"的记录:

```shell
sed -nr '/青海省|青/p' text.txt                                                                                   
1     province    省份  青海省
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址    青海省西宁市城中区南关街138号
7     site_no 网站备案/许可证号   青ICP备11000289号-2
```

`sed`正则匹配也支持按范围输出，格式为`/xx/,/xx/p`。例如查找包含"domain"的行到包含"addr"的行记录：

```shell
sed -n '/domain/,/addr/p' text.txt
2     domain  域名或者ip  tianfengyinlou.cn
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址    青海省西宁市城中区南关街138号
```

## 删除

```shell
# 删除第三行
sed 3d text.txt                                                                                                    
1     province    省份  青海省
2     domain  域名或者ip  tianfengyinlou.cn
4     addr    注册地址    青海省西宁市城中区南关街138号
5     check_time  备案时间, 时间对象  2011-06-23 16:38:00
6     update_time 更新时间, 毫秒级时间戳    1607414120745
7     site_no 网站备案/许可证号   青ICP备11000289号-2
8     site_url    站点/网站首页网址   www.tianfengyinlou.cn
9     comp_name   主办单位名称(公司名称)    西宁天丰银楼金银珠宝有限公司

# 删除包含青的行
sed '/青/d' text.txt                                                                                               
2     domain  域名或者ip  tianfengyinlou.cn
5     check_time  备案时间, 时间对象  2011-06-23 16:38:00
6     update_time 更新时间, 毫秒级时间戳    1607414120745
8     site_url    站点/网站首页网址   www.tianfengyinlou.cn
9     comp_name   主办单位名称(公司名称)    西宁天丰银楼金银珠宝有限公司

# 更改text.txt
cat text.txt                                                                                                       ysj@yangsj2-knownsec
1     province    省份  青海省
2     domain  域名或者ip  tianfengyinlou.cn

3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址    青海省西宁市城中区南关街138号
#5     check_time  备案时间, 时间对象  2011-06-23 16:38:00

#6     update_time 更新时间, 毫秒级时间戳    1607414120745
7     site_no 网站备案/许可证号   青ICP备11000289号-2
8     site_url    站点/网站首页网址   www.tianfengyinlou.cn
9     comp_name   主办单位名称(公司名称)    西宁天丰银楼金银珠宝有限公司

# 删除空行和注释行
sed -r '/^$|#/d' text.txt                                                                                          
1     province    省份  青海省
2     domain  域名或者ip  tianfengyinlou.cn
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址    青海省西宁市城中区南关街138号
7     site_no 网站备案/许可证号   青ICP备11000289号-2
8     site_url    站点/网站首页网址   www.tianfengyinlou.cn
9     comp_name   主办单位名称(公司名称)    西宁天丰银楼金银珠宝有限公司
```

## 增加

`sed`的增加动作有三种:

- `i`:在指定行的上方增加一行
- `a`: 在指定行的下方增加一行
- `c`: 在指定行的地方增加一行，原有行会被覆盖

上述三种增加行为示例为:

> 注意：示例的增加行为在mac上会报错，可能在mac上用法不一致。

```shell
# 在第3行上方增加一行记录
sed '3i insert oneline above 3rd line' text.txt 
1     province    省份  青海省
2     domain  域名或者ip  tianfengyinlou.cn
insert oneline above 3rd line
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址    青海省西宁市城中区南关街138号
5     check_time  备案时间, 时间对象  2011-06-23 16:38:00
6     update_time 更新时间, 毫秒级时间戳    1607414120745
7     site_no 网站备案/许可证号   青ICP备11000289号-2
8     site_url    站点/网站首页网址   www.tianfengyinlou.cn
9     comp_name   主办单位名称(公司名称)    西宁天丰银楼金银珠宝有限公司

# 在第3行下方增加一行记录
sed '3a insert oneline after 3rd line' text.txt 
1     province    省份  青海省
2     domain  域名或者ip  tianfengyinlou.cn
3     subject_no  主体备案号   青ICP备11000289号
insert oneline after 3rd line
4     addr    注册地址    青海省西宁市城中区南关街138号
5     check_time  备案时间, 时间对象  2011-06-23 16:38:00
6     update_time 更新时间, 毫秒级时间戳    1607414120745
7     site_no 网站备案/许可证号   青ICP备11000289号-2
8     site_url    站点/网站首页网址   www.tianfengyinlou.cn
9     comp_name   主办单位名称(公司名称)    西宁天丰银楼金银珠宝有限公司

# 在第3行创建一行记录，原记录被替换
sed '3c create oneline at 3rd line' text.txt 
1     province    省份  青海省
2     domain  域名或者ip  tianfengyinlou.cn
create oneline at 3rd line
4     addr    注册地址    青海省西宁市城中区南关街138号
5     check_time  备案时间, 时间对象  2011-06-23 16:38:00
6     update_time 更新时间, 毫秒级时间戳    1607414120745
7     site_no 网站备案/许可证号   青ICP备11000289号-2
8     site_url    站点/网站首页网址   www.tianfengyinlou.cn
9     comp_name   主办单位名称(公司名称)    西宁天丰银楼金银珠宝有限公司
```

## 修改

上述所有的操作输出均没有改变文件自身的内容。想要使得操作改变文件自身的内容，需要指定选项`-i`。指定`-i`的操作需要格外小心。

例如，在文件中第一行插入一行记录:

```shell
sed -i '1i add oneline above first line' text.txt
cat text.txt 
add oneline above first line
1     province    省份  青海省
2     domain  域名或者ip  tianfengyinlou.cn
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址    青海省西宁市城中区南关街138号
5     check_time  备案时间, 时间对象  2011-06-23 16:38:00
6     update_time 更新时间, 毫秒级时间戳    1607414120745
7     site_no 网站备案/许可证号   青ICP备11000289号-2
8     site_url    站点/网站首页网址   www.tianfengyinlou.cn
9     comp_name   主办单位名称(公司名称)    西宁天丰银楼金银珠宝有限公司
```

`sed`的删除动作为`d(delete)`,例如删除文件中的第一行:

```shell
# 删除增加的第一行
sed -i 1d text.txt 
cat text.txt 
1     province    省份  青海省
2     domain  域名或者ip  tianfengyinlou.cn
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址    青海省西宁市城中区南关街138号
5     check_time  备案时间, 时间对象  2011-06-23 16:38:00
6     update_time 更新时间, 毫秒级时间戳    1607414120745
7     site_no 网站备案/许可证号   青ICP备11000289号-2
8     site_url    站点/网站首页网址   www.tianfengyinlou.cn
9     comp_name   主办单位名称(公司名称)    西宁天丰银楼金银珠宝有限公司
```

当然，我们也有办法做安全的删除操作，即将`-i`换成`-i.bak` 可以在真实改动文件内容前，备份文件。但是这个操作一般不适合应用在大文件上，因为备份很慢。

```shell
# 删除第一行并备份
sed -i.bak 1d text.txt 
cat text.txt
2     domain  域名或者ip  tianfengyinlou.cn
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址    青海省西宁市城中区南关街138号
5     check_time  备案时间, 时间对象  2011-06-23 16:38:00
6     update_time 更新时间, 毫秒级时间戳    1607414120745
7     site_no 网站备案/许可证号   青ICP备11000289号-2
8     site_url    站点/网站首页网址   www.tianfengyinlou.cn
9     comp_name   主办单位名称(公司名称)    西宁天丰银楼金银珠宝有限公司
cat text.txt.bak 
1     province    省份  青海省
2     domain  域名或者ip  tianfengyinlou.cn
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址    青海省西宁市城中区南关街138号
5     check_time  备案时间, 时间对象  2011-06-23 16:38:00
6     update_time 更新时间, 毫秒级时间戳    1607414120745
7     site_no 网站备案/许可证号   青ICP备11000289号-2
8     site_url    站点/网站首页网址   www.tianfengyinlou.cn
9     comp_name   主办单位名称(公司名称)    西宁天丰银楼金银珠宝有限公司
```

## 替换

`sed`可以对文件内容进行替换`(substitute)`，格式为使用任意三个相同的符号，如三个斜线`s/xx/yy/g`、三个`#`号`s#xx#yy#g`、三个`@`符号`s@xx@yy@g`等，效果是将`xx`替换为`yy`。

> 这里的符号选择是任意的，可以是三个1，三个2都行。常用的是上述三种，因为和文件内容重合度最小，具体使用哪种，需要根据文件内容选择。如果文件内容本身包含了/，则不方便使用三个斜线来操作。

```shell
# 将"青" 替换为"蜀"
sed 's/青/蜀/g' text.txt                                                                                           
1     province    省份  蜀海省
2     domain  域名或者ip  tianfengyinlou.cn
3     subject_no  主体备案号   蜀ICP备11000289号
4     addr    注册地址    蜀海省西宁市城中区南关街138号
5     check_time  备案时间, 时间对象  2011-06-23 16:38:00
6     update_time 更新时间, 毫秒级时间戳    1607414120745
7     site_no 网站备案/许可证号   蜀ICP备11000289号-2
8     site_url    站点/网站首页网址   www.tianfengyinlou.cn
9     comp_name   主办单位名称(公司名称)    西宁天丰银楼金银珠宝有限公司

# 将第三行的青替换为蜀
sed '3s/青/蜀/g' text.txt                                                                                          
1     province    省份  青海省
2     domain  域名或者ip  tianfengyinlou.cn
3     subject_no  主体备案号   蜀ICP备11000289号
4     addr    注册地址    青海省西宁市城中区南关街138号
5     check_time  备案时间, 时间对象  2011-06-23 16:38:00
6     update_time 更新时间, 毫秒级时间戳    1607414120745
7     site_no 网站备案/许可证号   青ICP备11000289号-2
8     site_url    站点/网站首页网址   www.tianfengyinlou.cn
9     comp_name   主办单位名称(公司名称)    西宁天丰银楼金银珠宝有限公司

# 把所有数字替换为x
sed -r 's/[0-9]/x/g' text.txt                                                                                      
x     province    省份  青海省
x     domain  域名或者ip  tianfengyinlou.cn
x     subject_no  主体备案号   青ICP备xxxxxxxx号
x     addr    注册地址    青海省西宁市城中区南关街xxx号
x     check_time  备案时间, 时间对象  xxxx-xx-xx xx:xx:xx
x     update_time 更新时间, 毫秒级时间戳    xxxxxxxxxxxxx
x     site_no 网站备案/许可证号   青ICP备xxxxxxxx号-x
x     site_url    站点/网站首页网址   www.tianfengyinlou.cn
x     comp_name   主办单位名称(公司名称)    西宁天丰银楼金银珠宝有限公司
```

`g`是全局`(gloabal)`替换的意思，如果不要`g`，则只会替换匹配到的第一项:

```shell
# 把每一行的第一个数字替换为x
sed -r 's/[0-9]/x/' text.txt                                                                                      
x     province    省份  青海省
x     domain  域名或者ip  tianfengyinlou.cn
x     subject_no  主体备案号   青ICP备11000289号
x     addr    注册地址    青海省西宁市城中区南关街138号
x     check_time  备案时间, 时间对象  2011-06-23 16:38:00
x     update_time 更新时间, 毫秒级时间戳    1607414120745
x     site_no 网站备案/许可证号   青ICP备11000289号-2
x     site_url    站点/网站首页网址   www.tianfengyinlou.cn
x     comp_name   主办单位名称(公司名称)    西宁天丰银楼金银珠宝有限公司
```

## 反向引用

反向引用就是利用正则的组匹配来以组为单位进行替换。

```shell
# 例如匹配所有的英文词句([a-z_.]+)，然后把他们用<>括起来, \1表示第一组，这里只有一个组匹配
sed -r 's/([a-z_.]+)/<\1>/g' text.txt                                                                              
1     <province>    省份  青海省
2     <domain>  域名或者<ip>  <tianfengyinlou.cn>
3     <subject_no>  主体备案号   青ICP备11000289号
4     <addr>    注册地址    青海省西宁市城中区南关街138号
5     <check_time>  备案时间, 时间对象  2011-06-23 16:38:00
6     <update_time> 更新时间, 毫秒级时间戳    1607414120745
7     <site_no> 网站备案/许可证号   青ICP备11000289号-2
8     <site_url>    站点/网站首页网址   <www.tianfengyinlou.cn>
9     <comp_name>   主办单位名称(公司名称)    西宁天丰银楼金银珠宝有限公司
```

# awk

## 取行

`awk`可以通过`NR(Number of Record)` 指定行号，输出特定的行:

```shell
# 输出第三行
awk 'NR==3' text.txt                                                                                               
3     subject_no  主体备案号   青ICP备11000289号
```

也可以按行号范围输出:

```shell
# 输出第三到第六行
awk 'NR==3, NR==6' text.txt                                                                                        
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址    青海省西宁市城中区南关街138号
5     check_time  备案时间, 时间对象  2011-06-23 16:38:00
6     update_time 更新时间, 毫秒级时间戳    1607414120745

# 也可以通过比较指定输出范围
# 输出3到4行
awk 'NR>=3 && NR<5' text.txt                                                                                       
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址    青海省西宁市城中区南关街138号
```

第二个`NR`如果是个无效的行号值，则默认取出指定起始行之后所有的行记录：

```shell
# 输出第三行之后的所有行
awk 'NR==3, NR==xx' text.txt                                                                                       
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址    青海省西宁市城中区南关街138号
5     check_time  备案时间, 时间对象  2011-06-23 16:38:00
6     update_time 更新时间, 毫秒级时间戳    1607414120745
7     site_no 网站备案/许可证号   青ICP备11000289号-2
8     site_url    站点/网站首页网址   www.tianfengyinlou.cn
9     comp_name   主办单位名称(公司名称)    西宁天丰银楼金银珠宝有限公司

```

取行操作依然支持正则匹配:

```shell
# 输出包含青的行
awk '/青/' text.txt                                                                                                
1     province    省份  青海省
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址    青海省西宁市城中区南关街138号
7     site_no 网站备案/许可证号   青ICP备11000289号-2

# 输出以"号"结尾的行
awk '/号$/' text.txt                                                                                               
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址    青海省西宁市城中区南关街138号

# 输出包含domain到包含addr的行
awk '/domain/, /addr/' text.txt                                                                                    
2     domain  域名或者ip  tianfengyinlou.cn
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址    青海省西宁市城中区南关街138号
```

## 取列

`awk`可以使用`{print $列号}` 取出列值：

```shell
# 例如，取出第二列的值
awk '{print $2}' text.txt                                                                                      
province
domain
subject_no
addr
check_time
update_time
site_no
site_url
comp_name

# 取出第2列及最后一列NF(Number of Fields)的值
awk '{print $2,$NF}' text.txt                                                                                      
province 青海省
domain tianfengyinlou.cn
subject_no 青ICP备11000289号
addr 青海省西宁市城中区南关街138号
check_time 16:38:00
update_time 1607414120745
site_no 青ICP备11000289号-2
site_url www.tianfengyinlou.cn
comp_name 西宁天丰银楼金银珠宝有限公司

# 使用column -t 对齐输出
awk '{print $2,$NF}' text.txt | column -t                                                                          
province     青海省
domain       tianfengyinlou.cn
subject_no   青ICP备11000289号
addr         青海省西宁市城中区南关街138号
check_time   16:38:00
update_time  1607414120745
site_no      青ICP备11000289号-2
site_url     www.tianfengyinlou.cn
comp_name    西宁天丰银楼金银珠宝有限公司
```

`awk`取列时，默认是空格为分隔符，可以通过`-F`指定分隔符，例如，第7-8行：

```shell
awk "NR==7,NR==8" text.txt                                                                                         
7     site_no 网站备案/许可证号   青ICP备11000289号-2
8     site_url    站点/网站首页网址   www.tianfengyinlou.cn
```

取出第7、8行后，按`/`进行划分，取出划分后的第二列:

```shell
awk "NR==7,NR==8" text.txt | awk -F/ '{print $2}'                                                                  
许可证号   青ICP备11000289号-2
网站首页网址   www.tianfengyinlou.cn

```
`-F` 可以通过`[]`正则指定多个分隔符：

```shell
# 按空格和/ 进行分隔, 取出1到4列
awk "NR==7,NR==8" text.txt | awk -F'[ /]+' '{print $1,$2,$3,$4}'                                                   
7 site_no 网站备案 许可证号
8 site_url 站点 网站首页网址
```

## 精确取行列

`awk`可以精确取出某一行某一列的值。一些用例如：

```shell
# ~ 表示包含， !~ 表示不包含
# 取出第四列包含"青"的行
awk '$4 ~ /青/' text.txt                                                                                       
1     province    省份  青海省
3     subject_no  主体备案号   青ICP备11000289号
4     addr    注册地址    青海省西宁市城中区南关街138号
7     site_no 网站备案/许可证号   青ICP备11000289号-2

# 取出第四列以"号"结尾的行，并输出最后一列
awk '$4 ~ /号$/{print $NF}' text.txt                                                                               
青ICP备11000289号
青海省西宁市城中区南关街138号

# 取出第2列以d开始，到第四列以号结尾的行记录
awk '$2 ~ /^d/, $4 ~/号$/' text.txt                                                                                
2     domain  域名或者ip  tianfengyinlou.cn
3     subject_no  主体备案号   青ICP备11000289号
```

## BEGIN

`awk`可以使用`BEGIN`在操作文件内容前执行一些命令：

```shell
# 列如输出表头
awk 'BEGIN{print "序号","名称","含义","示例"} {print $1,$2,$3,$4}' text.txt | column -t                            
序号  名称          含义                   示例
1     province     省份                   青海省
2     domain       域名或者ip              tianfengyinlou.cn
3     subject_no   主体备案号              青ICP备11000289号
4     addr         注册地址                青海省西宁市城中区南关街138号
5     check_time   备案时间,               时间对象
6     update_time  更新时间,               毫秒级时间戳
7     site_no      网站备案/许可证号        青ICP备11000289号-2
8     site_url     站点/网站首页网址        www.tianfengyinlou.cn
9     comp_name    主办单位名称(公司名称)    西宁天丰银楼金银珠宝有限公司
```

## END

`awk`可以使用`END`在操作文件内容后执行一些命令：

```shell
# 通常用于做统计， 例如对第一列求和
awk '{sum+=$1} END{print sum}' text.txt                                                                            
45
```

# 使用小结

1. `grep、sed、awk`都可以过滤行记录，但过滤行记录时优先选择`grep`，其过滤行的效率最高。
2. `sed`主要用于对文件内容做出各种修改(增加、替换等)。
3. `awk`主要用于对文件内容取行列操作。

