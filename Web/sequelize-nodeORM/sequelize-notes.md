# 链接

https://github.com/demopark/sequelize-docs-Zh-CN

# 安装

```shell
npm install --save sequelize

# 还要安装对应驱动
npm install --save pg pg-hstore # Postgres
npm install --save mysql2
npm install --save mariadb
npm install --save sqlite3
npm install --save tedious # Microsoft SQL Server
```
# 连接
> sequelize的约定：
> Sequelize:数据库本身
> sequelize：数据库Sequelize的实例
## 三种连接方式
```js
const {Sequelize} = require('sequelize')

//方式一：传递一个URI
const sequelize = new Sequelize('sqlite::memory::') //sqlite
const sequelize = new Sequelize('postgres://username:password@domain.com:port/dbname') //postgresql

//方式二：传递参数<sqlite>
const sequelize = new Sequelize({
    dialect:'sqlite',
    storage:'path/to/database.sqlite'
})

//方式三:传递参数<other db>
const sequelize = new Sequelize('database','username','password',{
    host:'localhost',
    dialect:'mysql' // [mariadb|postgres|mssql|mysql]
})
```
## 测试数据库连接
```js
// 使用 .authenticate() 函数测试连接是否正常
try {
    await sequelize.authenticate()
    console.log('connection succed')
} catch(error) {
    console.error('connection failed',error)
}

//关闭连接
sequelize.close() //Promise
```
## 记录日志
```js
const sequelize = new Sequelize('sqlite::memory:', {
    // 选择一种日志记录参数
    logging: console.log,                  // 默认值,显示日志函数调用的第一个参数
    logging: (...msg) => console.log(msg), // 显示所有日志函数调用参数
    logging: false,                        // 禁用日志记录
    logging: msg => logger.debug(msg),     // 使用自定义记录器(例如Winston 或 Bunyan),显示第一个参数
    logging: logger.debug.bind(logger)     // 使用自定义记录器的另一种方法,显示所有消息
  });
```
# 模型<Model>
> 模型对应数据库表
> 表名推断: Sequelize 会自动将模型名复数用作表名
> Sequelize 使用数据类型 DataTypes.DATE 时会自动向每个模型添加 createdAt 和 updatedAt 字段
## 创建模型
> 有两种等效的方式：
> * 调用 sequelize.define(modelName, attributes, options)
> * 扩展 Model 并调用 init(attributes, options)

```js
const { Sequelize, DataTypes,Model} = require('sequelize')
const sequelize = new Sequelize('sqlite::memory:'，{
    define:{
        freezeTableName:true //强制所有实例表名与模型名相同
    }
})

//方式一
const User = sequelize.define('User',{
    firstName:{
        type:DataTypes.STRING, //VARCHAR(255)
        allowNull:false
    },
    lastName:{
        type:DataTypes.STRING,
        allowNull:false
    },
    time:{
        type:DataTypes.DATE,
        allowNull:true, //默认
        defaultValue:Sequelize.NOW
    }
},{
    freezeTableName:true, //强制表名与模型名相同
    tableName: 'userTable', //自定义表名
    timestamps:false //禁用时间戳自动生成createAt和updateAt字段
})
console.log(User === sequelize.models.User)

//方式二
class User extends Model {}
User.init({
    firstName:{
        type:DataTypes.STRING,
        allowNull:false
    },
    lastName:{
        type:DataTypes.STRING,
        allowNull:false
    },
    male:{ //可以简写为 male:DataTypes.BOOLEAN
        type:DataTypes.BOOLEAN, //TINYINT(1)
        allowNull:true //默认
    }
},{
    sequelize, //需要传递连接实例
    modelName:'User',
    freezeTableName:true, //强制表名与模型名相同
    tableName: 'userTable', //自定义表名
    
    timestamps: true,// 启用时间戳！
    createdAt: false,// 不想要 createdAt
    updatedAt: 'updateTimestamp'//想要 updatedAt 但是希望名称叫做 updateTimestamp
})

console.log(User === sequelize.models.User)
```

## 模型同步
```js
// User.sync() - 如果表不存在,则创建该表(如果已经存在,则不执行任何操作)
// User.sync({ force: true }) - 将创建表,如果表已经存在,则将其首先删除
// User.sync({ alter: true }) - 这将检查数据库中表的当前状态(它具有哪些列,它们的数据类型等),然后在表中进行必要的更改以使其与模型匹配.
await User.sync() 
await sequelize.sync() //同步所有模型
await sequelize.sync({ force: true, match: /_test$/ }) // 仅当数据库名称以 '_test' 结尾时,它才会运行.sync()

await User.drop() 
await sequelize.drop() //删除所有表
```

## 列参数
```js
const { Model, DataTypes, Deferrable } = require("sequelize");

class Foo extends Model {}
Foo.init({
  // 实例化将自动将 flag 设置为 true (如果未设置)
  flag: { type: DataTypes.BOOLEAN, allowNull: false, defaultValue: true },

  // 日期的默认值 => 当前时间
  myDate: { type: DataTypes.DATE, defaultValue: DataTypes.NOW },

  // 将 allowNull 设置为 false 将为该列添加 NOT NULL,
  // 这意味着如果该列为 null,则在执行查询时将从数据库引发错误.
  // 如果要在查询数据库之前检查值是否不为 null,请查看下面的验证部分.
  title: { type: DataTypes.STRING, allowNull: false },

  // 创建两个具有相同值的对象将引发错误.
  // unique 属性可以是布尔值或字符串.
  // 如果为多个列提供相同的字符串,则它们将形成一个复合唯一键.
  uniqueOne: { type: DataTypes.STRING,  unique: 'compositeIndex' },
  uniqueTwo: { type: DataTypes.INTEGER, unique: 'compositeIndex' },

  // unique 属性是创建唯一约束的简写.
  someUnique: { type: DataTypes.STRING, unique: true },

  // 继续阅读有关主键的更多信息
  identifier: { type: DataTypes.STRING, primaryKey: true },

  // autoIncrement 可用于创建 auto_incrementing 整数列
  incrementMe: { type: DataTypes.INTEGER, autoIncrement: true },

  // 你可以通过 'field' 属性指定自定义列名称：
  fieldWithUnderscores: { type: DataTypes.STRING, field: 'field_with_underscores' },

  // 可以创建外键：
  bar_id: {
    type: DataTypes.INTEGER,

    references: {
      // 这是对另一个模型的参考
      model: Bar,

      // 这是引用模型的列名
      key: 'id',

      // 使用 PostgreSQL,可以通过 Deferrable 类型声明何时检查外键约束.
      deferrable: Deferrable.INITIALLY_IMMEDIATE
      // 参数:
      // - `Deferrable.INITIALLY_IMMEDIATE` - 立即检查外键约束
      // - `Deferrable.INITIALLY_DEFERRED` - 将所有外键约束检查推迟到事务结束
      // - `Deferrable.NOT` - 完全不推迟检查(默认) - 这将不允许你动态更改事务中的规则
    }
  },

  // 注释只能添加到 MySQL,MariaDB,PostgreSQL 和 MSSQL 的列中
  commentMe: {
    type: DataTypes.INTEGER,
    comment: '这是带有注释的列'
  }
}, {
  sequelize,
  modelName: 'foo',

  // 在上面的属性中使用 `unique: true` 与在模型的参数中创建索引完全相同：
  indexes: [{ unique: true, fields: ['someUnique'] }]
});
```
## 自定义模型方法
```js
class User extends Model {
  static classLevelMethod() {
    return 'foo';
  }
  instanceLevelMethod() {
    return 'bar';
  }
  getFullname() {
    return [this.firstname, this.lastname].join(' ');
  }
}
User.init({
  firstname: Sequelize.TEXT,
  lastname: Sequelize.TEXT
}, { sequelize });

console.log(User.classLevelMethod()); // 'foo'
const user = User.build({ firstname: 'Jane', lastname: 'Doe' });
console.log(user.instanceLevelMethod()); // 'bar'
console.log(user.getFullname()); // 'Jane Doe'
```

## 创建模型实例
```js
const jodan = User.build({firstName:'Michael',lastName:'Jodan'}) //同步的
jodan.lastName = 'JOJO' //更新字段
await jodan.save() //保存(或更新)到数据库
await jodan.save({fields:['fistName']}) //仅保存特定字段到数据库中

const rose = await User.create({firstName:'Michael',lastName:'Rose'}) //build+save
rose.firstName = 'Linus'
await rose.reload() //重载实例,reload 调用生成一个 SELECT 查询,以从数据库中获取最新数据.
console.log(rose.toJSON())

//字段递增
await rose.increment('age',{by:2})
await rose.increment({
  'age': 2,
  'cash': 100
})

// 如果值增加相同的数量,则也可以使用以下其他语法：
await jane.increment(['age', 'cash'], { by: 2 });

await rose.destroy() //从数据库删除
```
# 模型查询

## 查找器

### findAll

> 见 select -> select all

### findByPk

```js
//通过主键从表中获得一个条目.
const project = await Project.findByPk(123);
if (project === null) {
  console.log('Not found!');
} else {
  console.log(project instanceof Project); // true
  // 它的主键是 123
}
```

### findOne

```js
const project = await Project.findOne({ where: { title: 'My Title' } });
if (project === null) {
  console.log('Not found!');
} else {
  console.log(project instanceof Project); // true
  console.log(project.title); // 'My Title'
}
```

### findOrCreate

```js
//找到一个满足查询参数的结果,否则方法 findOrCreate 将在表中创建一个条目. 在这两种情况下,它将返回一个实例(找到的实例或创建的实例)和一个布尔值,指示该实例是已创建还是已经存在.
const [user, created] = await User.findOrCreate({
  where: { username: 'sdepold' },
  defaults: {
    job: 'Technical Lead JavaScript'
  }
});
console.log(user.username); // 'sdepold'
console.log(user.job); // 这可能是也可能不是 'Technical Lead JavaScript'
console.log(created); // 指示此实例是否刚刚创建的布尔值
if (created) {
  console.log(user.job); // 这里肯定是 'Technical Lead JavaScript'
}
```

### findAndCountAll

```js
// findAndCountAll 方法返回一个具有两个属性的对象：
// count - 一个整数 - 符合查询条件的记录总数
// rows - 一个数组对象 - 获得的记录

const { count, rows } = await Project.findAndCountAll({
  where: {
    title: {
      [Op.like]: 'foo%'
    }
  },
  offset: 10,
  limit: 2
});
console.log(count);
console.log(rows);
```

## insert

### create

```js
const user = await User.create({
  username:'xxx',
  isAdmin:true
},{fields:['username']}) //只有username才会保存到数据库中
```
### bulkCreate

```js
//批量创建
//默认情况下,bulkCreate 不会在要创建的每个对象上运行验证(而 create 可以做到). 为了使 bulkCreate 也运行这些验证,必须通过validate: true 参数. 但这会降低性能. 用法示例：
const Foo = sequelize.define('foo', {
  name: {
    type: DataTypes.TEXT,
    validate: {
      len: [4, 6]
    }
  }
});

// 这不会引发错误,两个实例都将被创建
await Foo.bulkCreate([
  { name: 'abc123' },
  { name: 'name too long' }
],{ fields: ['name'] }));

// 这将引发错误,不会创建任何内容
await Foo.bulkCreate([
  { name: 'abc123' },
  { name: 'name too long' }
], { validate: true });
```

## select

### select all
> SELECT * FROM TABLE
```js
const users = await User.findAll() //return Array<Object>
console.log('all users:',JSON.stringify(users,null,2))
```

### select some field
> SELECT foo,bar FROM TABLE
```js
Model.findAll({
  attributes:['foo','bar']
})
```

### select as
> SELECT foo,bar as baz,qux FROM TABLE
```js
Model.finAll({
  attributes:['foo',['bar','baz'],'qux']
})
```

### select aggregation as 
> SELECT foo,COUNT(bar) AS n_bar,qux FROM TABLE
```js
Model.finAll({
  attributes:[
    'foo',
    [sequelize.fn('COUNT',sequelize.col('bar')),'n_bar'],
    'qux'
  ]
})
```

### select all <include + exclude>
> SELECT *,COUNT(baar) AS n_bar FROM TABLE
```js
Model.findAll({
  attributes:{
    include:[
      [sequelize.fn('COUNT',sequelize.col('bar')),'n_bar']
    ],
    exclude:['foo']
  }
})

```

### select where
> SELECT * FROM TABLE WHERE id=2 
```js
Model.findAll({
  where:{
    id:2,
  }
})
//等同于
const {Op} = require("sequelize")
Model.findAll({
  where:{
    id:{
      [Op.eq]:2
    }
  }
})
```
> SELECT * FROM TABLE WHERE id=2 AND status='active'
> SELECT * FROM TABLE WHERE id=2 OR id=3
```js
// AND
const {Op} = require('sequelize')
Model.finAll({
  where:{
    [Op.and]:[
      {id:2},
      {status:'active'}
    ]
  }
})
// OR
Model.finAll({
  where:{
    [Op.or]:[
      {id:2},
      {id:3}
    ]
  }
})
```

## delete
> DELETE FORM TABLE WHERE id=2 OR id=3
```js
const {Op} = require('sequelize')
Model.destroy({
  where:{
    id:{
      [Op.or]:[2,3]
    }
  }
})

// 删除所有名为 "Jane" 的人 
await User.destroy({
  where: {
    firstName: "Jane"
  }
});

//销毁所有内容,可以使用 TRUNCATE SQL：
await User.destroy({
  truncate: true
});
```
## update

```js
await User.update({lastName:'Doe'},{
    where:{
        lastName:null
    }
})
```

## order by

```js
Subtask.findAll({
  order: [
    // 将转义 title 并针对有效方向列表进行降序排列
    ['title', 'DESC'], //[column, direction] 形式的数组

    // 将按最大年龄进行升序排序
    sequelize.fn('max', sequelize.col('age')),

    // 将按最大年龄进行降序排序
    [sequelize.fn('max', sequelize.col('age')), 'DESC'],

    // 将按 otherfunction(`col1`, 12, 'lalala') 进行降序排序
    [sequelize.fn('otherfunction', sequelize.col('col1'), 12, 'lalala'), 'DESC'],

    // 将使用模型名称作为关联名称按关联模型的 createdAt 排序.
    [Task, 'createdAt', 'DESC'],

    // 将使用模型名称作为关联名称通过关联模型的 createdAt 排序.
    [Task, Project, 'createdAt', 'DESC'],

    // 将使用关联名称按关联模型的 createdAt 排序.
    ['Task', 'createdAt', 'DESC'],

    // 将使用关联的名称按嵌套的关联模型的 createdAt 排序.
    ['Task', 'Project', 'createdAt', 'DESC'],

    // 将使用关联对象按关联模型的 createdAt 排序. (首选方法)
    [Subtask.associations.Task, 'createdAt', 'DESC'],

    // 将使用关联对象按嵌套关联模型的 createdAt 排序. (首选方法)
    [Subtask.associations.Task, Task.associations.Project, 'createdAt', 'DESC'],

    // 将使用简单的关联对象按关联模型的 createdAt 排序.
    [{model: Task, as: 'Task'}, 'createdAt', 'DESC'],

    // 将由嵌套关联模型的 createdAt 简单关联对象排序.
    [{model: Task, as: 'Task'}, {model: Project, as: 'Project'}, 'createdAt', 'DESC']
  ],

  // 将按最大年龄降序排列
  order: sequelize.literal('max(age) DESC'),

  // 如果忽略方向,则默认升序,将按最大年龄升序排序
  order: sequelize.fn('max', sequelize.col('age')),

  // 如果省略方向,则默认升序, 将按年龄升序排列
  order: sequelize.col('age'),

  // 将根据方言随机排序(但不是 fn('RAND') 或 fn('RANDOM'))
  order: sequelize.random()
});

Foo.findOne({
  order: [
    // 将返回 `name`
    ['name'],
    // 将返回 `username` DESC
    ['username', 'DESC'],
    // 将返回 max(`age`)
    sequelize.fn('max', sequelize.col('age')),
    // 将返回 max(`age`) DESC
    [sequelize.fn('max', sequelize.col('age')), 'DESC'],
    // 将返回 otherfunction(`col1`, 12, 'lalala') DESC
    [sequelize.fn('otherfunction', sequelize.col('col1'), 12, 'lalala'), 'DESC'],
    // 将返回 otherfunction(awesomefunction(`col`)) DESC, 这种嵌套可能是无限的!
    [sequelize.fn('otherfunction', sequelize.fn('awesomefunction', sequelize.col('col'))), 'DESC']
  ]
});
```

## group by

```js
Project.findAll({ group: 'name' });
// 生成 'GROUP BY name'
```

## limit/offset/pagination

```js
// 提取10个实例/行
Project.findAll({ limit: 10 });

// 跳过8个实例/行
Project.findAll({ offset: 8 });

// 跳过5个实例,然后获取5个实例
Project.findAll({ offset: 5, limit: 5 });
```

# 实用方法

##  count

```js
console.log(`这有 ${await Project.count()} 个项目`);

const amount = await Project.count({
  where: {
    id: {
      [Op.gt]: 25
    }
  }
});
console.log(`这有 ${amount} 个项目 id 大于 25`);
```

## max/min/sum

```js
await User.max('age'); // 40
await User.max('age', { where: { age: { [Op.lt]: 20 } } }); // 10
await User.min('age'); // 5
await User.min('age', { where: { age: { [Op.gt]: 5 } } }); // 10
await User.sum('age'); // 55
await User.sum('age', { where: { age: { [Op.gt]: 5 } } }); // 50
```

# 查询操作符 Op

## 一般查询
```js
const { Op } = require("sequelize");
Post.findAll({
  where: {
    [Op.and]: [{ a: 5 }, { b: 6 }],            // (a = 5) AND (b = 6)
    [Op.or]: [{ a: 5 }, { b: 6 }],             // (a = 5) OR (b = 6)
    someAttribute: {
      // 基本
      [Op.eq]: 3,                              // = 3
      [Op.ne]: 20,                             // != 20
      [Op.is]: null,                           // IS NULL
      [Op.not]: true,                          // IS NOT TRUE
      [Op.or]: [5, 6],                         // (someAttribute = 5) OR (someAttribute = 6)

      // 使用方言特定的列标识符 (以下示例中使用 PG):
      [Op.col]: 'user.organization_id',        // = "user"."organization_id"

      // 数字比较
      [Op.gt]: 6,                              // > 6
      [Op.gte]: 6,                             // >= 6
      [Op.lt]: 10,                             // < 10
      [Op.lte]: 10,                            // <= 10
      [Op.between]: [6, 10],                   // BETWEEN 6 AND 10
      [Op.notBetween]: [11, 15],               // NOT BETWEEN 11 AND 15

      // 其它操作符

      [Op.all]: sequelize.literal('SELECT 1'), // > ALL (SELECT 1)

      [Op.in]: [1, 2],                         // IN [1, 2]
      [Op.notIn]: [1, 2],                      // NOT IN [1, 2]

      [Op.like]: '%hat',                       // LIKE '%hat'
      [Op.notLike]: '%hat',                    // NOT LIKE '%hat'
      [Op.startsWith]: 'hat',                  // LIKE 'hat%'
      [Op.endsWith]: 'hat',                    // LIKE '%hat'
      [Op.substring]: 'hat',                   // LIKE '%hat%'
      [Op.iLike]: '%hat',                      // ILIKE '%hat' (不区分大小写) (仅 PG)
      [Op.notILike]: '%hat',                   // NOT ILIKE '%hat'  (仅 PG)
      [Op.regexp]: '^[h|a|t]',                 // REGEXP/~ '^[h|a|t]' (仅 MySQL/PG)
      [Op.notRegexp]: '^[h|a|t]',              // NOT REGEXP/!~ '^[h|a|t]' (仅 MySQL/PG)
      [Op.iRegexp]: '^[h|a|t]',                // ~* '^[h|a|t]' (仅 PG)
      [Op.notIRegexp]: '^[h|a|t]',             // !~* '^[h|a|t]' (仅 PG)

      [Op.any]: [2, 3],                        // ANY ARRAY[2, 3]::INTEGER (仅 PG)

      // 在 Postgres 中, Op.like/Op.iLike/Op.notLike 可以结合 Op.any 使用:
      [Op.like]: { [Op.any]: ['cat', 'hat'] }  // LIKE ANY ARRAY['cat', 'hat']

      // 还有更多的仅限 postgres 的范围运算符,请参见下文
    }
  }
});

//Op.in 简写语法
Post.findAll({
  where: {
    id: [1,2,3] // 等同使用 `id: { [Op.in]: [1,2,3] }`
  }
});
// SELECT ... FROM "posts" AS "post" WHERE "post"."id" IN (1, 2, 3);
```
```js

// 运算符的逻辑组合
// 运算符 Op.and, Op.or 和 Op.not 可用于创建任意复杂的嵌套逻辑比较.
// 使用 Op.and 和 Op.or 示例:
const { Op } = require("sequelize");

Foo.findAll({
  where: {
    rank: {
      [Op.or]: {
        [Op.lt]: 1000,
        [Op.eq]: null
      }
    },
    // rank < 1000 OR rank IS NULL

    {
      createdAt: {
        [Op.lt]: new Date(),
        [Op.gt]: new Date(new Date() - 24 * 60 * 60 * 1000)
      }
    },
    // createdAt < [timestamp] AND createdAt > [timestamp]

    {
      [Op.or]: [
        {
          title: {
            [Op.like]: 'Boat%'
          }
        },
        {
          description: {
            [Op.like]: '%boat%'
          }
        }
      ]
    }
    // title LIKE 'Boat%' OR description LIKE '%boat%'
);

// 使用 Op.not 示例:
Project.findAll({
  where: {
    name: 'Some Project',
    [Op.not]: [
      { id: [1,2,3] },
      {
        description: {
          [Op.like]: 'Hello%'
        }
      }
    ]
  }
});

// 上面将生成：

SELECT *
FROM `Projects`
WHERE (
  `Projects`.`name` = 'a project'
  AND NOT (
    `Projects`.`id` IN (1,2,3)
    OR
    `Projects`.`description` LIKE 'Hello%'
  )
)
```
## 高级查询
```js
Post.findAll({
  where: sequelize.where(sequelize.fn('char_length', sequelize.col('content')), 7)
});
// SELECT ... FROM "posts" AS "post" WHERE char_length("content") = 7
```
```js
Post.findAll({
  where: {
    [Op.or]: [
      sequelize.where(sequelize.fn('char_length', sequelize.col('content')), 7),
      {
        content: {
          [Op.like]: 'Hello%'
        }
      },
      {
        [Op.and]: [
          { status: 'draft' },
          sequelize.where(sequelize.fn('char_length', sequelize.col('content')), {
            [Op.gt]: 10
          })
        ]
      }
    ]
  }
});

// 上面生成了以下SQL：

// SELECT
//   ...
// FROM "posts" AS "post"
// WHERE (
//   char_length("content") = 7
//   OR
//   "post"."content" LIKE 'Hello%'
//   OR (
//     "post"."status" = 'draft'
//     AND
//     char_length("content") > 10
//   )
// )
```

## 仅限 Postgres 的范围运算符
```js
[Op.contains]: 2,            // @> '2'::integer  (PG range 包含元素运算符)
[Op.contains]: [1, 2],       // @> [1, 2)        (PG range 包含范围运算符)
[Op.contained]: [1, 2],      // <@ [1, 2)        (PG range 包含于运算符)
[Op.overlap]: [1, 2],        // && [1, 2)        (PG range 重叠(有共同点)运算符)
[Op.adjacent]: [1, 2],       // -|- [1, 2)       (PG range 相邻运算符)
[Op.strictLeft]: [1, 2],     // << [1, 2)        (PG range 左严格运算符)
[Op.strictRight]: [1, 2],    // >> [1, 2)        (PG range 右严格运算符)
[Op.noExtendRight]: [1, 2],  // &< [1, 2)        (PG range 未延伸到右侧运算符)
[Op.noExtendLeft]: [1, 2],   // &> [1, 2)        (PG range 未延伸到左侧运算符)
```

# Getters, Setters & Virtuals

## getters 获取器

```js
const User = sequelize.define('user', {
  // 假设我们想要以大写形式查看每个用户名,
  // 即使它们在数据库本身中不一定是大写的
  username: {
    type: DataTypes.STRING,
    get() {
      const rawValue = this.getDataValue(username);
      return rawValue ? rawValue.toUpperCase() : null;
    }
  }
});

const user = User.build({ username: 'SuperUser123' });
console.log(user.username); // 'SUPERUSER123'
console.log(user.getDataValue(username)); // 'SuperUser123'
```

## setters 设置器

```js
const User = sequelize.define('user', {
  username: DataTypes.STRING,
  password: {
    type: DataTypes.STRING,
    set(value) {
      // 在数据库中以明文形式存储密码是很糟糕的.
      // 使用适当的哈希函数来加密哈希值更好.
      this.setDataValue('password', hash(value));
    }
  }
});
const user = User.build({ username: 'someone', password: 'NotSo§tr0ngP4$SW0RD!' });
console.log(user.password); // '7cfc84b8ea898bb72462e78b4643cfccd77e9f05678ec2ce78754147ba947acc'
console.log(user.getDataValue(password)); // '7cfc84b8ea898bb72462e78b4643cfccd77e9f05678ec2ce78754147ba947acc'

//-----------
const User = sequelize.define('user', {
  username: DataTypes.STRING,
  password: {
    type: DataTypes.STRING,
    set(value) {
      // 在数据库中以明文形式存储密码是很糟糕的.
      // 使用适当的哈希函数来加密哈希值更好.
      // 使用用户名作为盐更好.
      this.setDataValue('password', hash(this.username + value));
    }
  }
});
```

## virtuals 虚拟字段

```js
//虚拟字段是 Sequelize 在后台填充的字段,但实际上它们不存在于数据库中.
const { DataTypes } = require("sequelize");

const User = sequelize.define('user', {
  firstName: DataTypes.TEXT,
  lastName: DataTypes.TEXT,
  fullName: {
    type: DataTypes.VIRTUAL,
    get() {
      return `${this.firstName} ${this.lastName}`;
    },
    set(value) {
      throw new Error('不要尝试设置 `fullName` 的值!');
    }
  }
});

const user = await User.create({ firstName: 'John', lastName: 'Doe' });
console.log(user.fullName); // 'John Doe'
```

# Validations & Constraints - 验证 & 约束

在本教程中,你将学习如何在 Sequelize 中设置模型的验证和约束.对于本教程,将假定以下设置：

```js
const { Sequelize, Op, Model, DataTypes } = require("sequelize");
const sequelize = new Sequelize("sqlite::memory:");

const User = sequelize.define("user", {
  username: {
    type: DataTypes.TEXT,
    allowNull: false,
    unique: true
  },
  hashedPassword: {
    type: DataTypes.STRING(64),
    is: /^[0-9a-f]{64}$/i
  }
});

(async () => {
  await sequelize.sync({ force: true });
  // 这是代码
})();
```

## 验证和约束的区别

验证是在纯 JavaScript 中在 Sequelize 级别执行的检查. 如果你提供自定义验证器功能,它们可能会非常复杂,也可能是 Sequelize 提供的内置验证器之一. 如果验证失败,则根本不会将 SQL 查询发送到数据库.

另一方面,约束是在 SQL 级别定义的规则. 约束的最基本示例是唯一约束. 如果约束检查失败,则数据库将引发错误,并且 Sequelize 会将错误转发给 JavaScript(在此示例中,抛出 `SequelizeUniqueConstraintError`). 请注意,在这种情况下,与验证不同,它执行了 SQL 查询.

## 唯一约束

下面的代码示例在 `username` 字段上定义了唯一约束：

```js
/* ... */ {
  username: {
    type: DataTypes.TEXT,
    allowNull: false,
    unique: true
  },
} /* ... */
```

同步此模型后(例如,通过调用`sequelize.sync`),在表中将 `username` 字段创建为` `name` TEXT UNIQUE`,如果尝试插入已存在的用户名将抛出 `SequelizeUniqueConstraintError`.

## 允许/禁止 null 值

默认情况下,`null` 是模型每一列的允许值. 可以通过为列设置 `allowNull: false` 参数来禁用它,就像在我们的代码示例的 `username` 字段中所做的一样：

```js
/* ... */ {
  username: {
    type: DataTypes.TEXT,
    allowNull: false,
    unique: true
  },
} /* ... */
```

如果没有 `allowNull: false`, 那么调用 `User.create({})` 将会生效.

### 关于 `allowNull` 实现的说明

按照本教程开头所述,`allowNull` 检查是 Sequelize 中唯一由 *验证* 和 *约束* 混合而成的检查. 这是因为：

- 如果试图将 `null` 设置到不允许为 null 的字段,则将抛出`ValidationError` ,而且 *不会执行任何 SQL 查询*.
- 另外,在 `sequelize.sync` 之后,具有 `allowNull: false` 的列将使用 `NOT NULL` SQL 约束进行定义. 这样,尝试将值设置为 `null` 的直接 SQL 查询也将失败.

## 验证器

使用模型验证器,可以为模型的每个属性指定 格式/内容/继承 验证. 验证会自动在 `create`, `update` 和 `save` 时运行. 你还可以调用 `validate()` 来手动验证实例.

### 按属性验证

你可以定义你的自定义验证器,也可以使用由 [validator.js (10.11.0)](https://github.com/chriso/validator.js) 实现的多个内置验证器,如下所示.

```js
sequelize.define('foo', {
  bar: {
    type: DataTypes.STRING,
    validate: {
      is: /^[a-z]+$/i,          // 匹配这个 RegExp
      is: ["^[a-z]+$",'i'],     // 与上面相同,但是以字符串构造 RegExp
      not: /^[a-z]+$/i,         // 不匹配 RegExp
      not: ["^[a-z]+$",'i'],    // 与上面相同,但是以字符串构造 RegExp
      isEmail: true,            // 检查 email 格式 (foo@bar.com)
      isUrl: true,              // 检查 url 格式 (http://foo.com)
      isIP: true,               // 检查 IPv4 (129.89.23.1) 或 IPv6 格式
      isIPv4: true,             // 检查 IPv4 格式 (129.89.23.1)
      isIPv6: true,             // 检查 IPv6 格式
      isAlpha: true,            // 只允许字母
      isAlphanumeric: true,     // 将仅允许使用字母数字,因此 '_abc' 将失败
      isNumeric: true,          // 只允许数字
      isInt: true,              // 检查有效的整数
      isFloat: true,            // 检查有效的浮点数
      isDecimal: true,          // 检查任何数字
      isLowercase: true,        // 检查小写
      isUppercase: true,        // 检查大写
      notNull: true,            // 不允许为空
      isNull: true,             // 只允许为空
      notEmpty: true,           // 不允许空字符串
      equals: 'specific value', // 仅允许 'specific value'
      contains: 'foo',          // 强制特定子字符串
      notIn: [['foo', 'bar']],  // 检查值不是这些之一
      isIn: [['foo', 'bar']],   // 检查值是其中之一
      notContains: 'bar',       // 不允许特定的子字符串
      len: [2,10],              // 仅允许长度在2到10之间的值
      isUUID: 4,                // 只允许 uuid
      isDate: true,             // 只允许日期字符串
      isAfter: "2011-11-05",    // 仅允许特定日期之后的日期字符串
      isBefore: "2011-11-05",   // 仅允许特定日期之前的日期字符串
      max: 23,                  // 仅允许值 <= 23
      min: 23,                  // 仅允许值 >= 23
      isCreditCard: true,       // 检查有效的信用卡号

      // 自定义验证器的示例:
      isEven(value) {
        if (parseInt(value) % 2 !== 0) {
          throw new Error('Only even values are allowed!');
        }
      }
      isGreaterThanOtherField(value) {
        if (parseInt(value) <= parseInt(this.otherField)) {
          throw new Error('Bar must be greater than otherField.');
        }
      }
    }
  }
});
```

请注意,在需要将多个参数传递给内置验证函数的情况下,要传递的参数必须位于数组中. 但是,如果要传递单个数组参数,例如,`isIn` 可接受的字符串数组,则将其解释为多个字符串参数,而不是一个数组参数. 要解决此问题,请传递一个单长度的参数数组,例如上面所示的 `[['foo', 'bar']]` .

要使用自定义错误消息而不是 [validator.js](https://github.com/chriso/validator.js) 提供的错误消息,请使用对象而不是纯值或参数数组,例如验证器 不需要参数就可以给自定义消息

```js
isInt: {
  msg: "必须是价格的整数"
}
```

或者如果还需要传递参数,则添加一个 `args` 属性：

```js
isIn: {
  args: [['en', 'zh']],
  msg: "必须为英文或中文"
}
```

使用自定义验证器功能时,错误消息将是抛出的 `Error` 对象所持有的任何消息.

有关内置验证方法的更多详细信息,请参见[validator.js 项目](https://github.com/chriso/validator.js).

**提示:** 你还可以为日志记录部分定义自定义功能. 只需传递一个函数. 第一个参数是记录的字符串.

### `allowNull` 与其他验证器的交互

如果将模型的特定字段设置为不允许为 null(使用 `allowNull: false`),并且该值已设置为 `null`,则将跳过所有验证器,并抛出 `ValidationError`.

另一方面,如果将其设置为允许 null(使用 `allowNull: true`),并且该值已设置为 `null`,则仅会跳过内置验证器,而自定义验证器仍将运行.

举例来说,这意味着你可以拥有一个字符串字段,该字段用于验证其长度在5到10个字符之间,但也允许使用 `null` (因为当该值为 `null` 时,长度验证器将被自动跳过)：

```js
class User extends Model {}
User.init({
  username: {
    type: DataTypes.STRING,
    allowNull: true,
    validate: {
      len: [5, 10]
    }
  }
}, { sequelize });
```

你也可以使用自定义验证器有条件地允许 `null` 值,因为不会跳过它：

```js
class User extends Model {}
User.init({
  age: Sequelize.INTEGER,
  name: {
    type: DataTypes.STRING,
    allowNull: true,
    validate: {
      customValidator(value) {
        if (value === null && this.age !== 10) {
          throw new Error("除非年龄为10,否则名称不能为 null");
        }
      })
    }
  }
}, { sequelize });
```

你可以通过设置 `notNull` 验证器来自定义 `allowNull` 错误消息：

```js
class User extends Model {}
User.init({
  name: {
    type: DataTypes.STRING,
    allowNull: false,
    validate: {
      notNull: {
        msg: '请输入你的名字'
      }
    }
  }
}, { sequelize });
```

### 模型范围内的验证

还可以定义验证,来在特定于字段的验证器之后检查模型. 例如,使用此方法,可以确保既未设置 `latitude` 和 `longitude`,又未同时设置两者. 如果设置了一个但未设置另一个,则失败.

使用模型对象的上下文调用模型验证器方法,如果它们抛出错误,则认为失败,否则将通过. 这与自定义字段特定的验证器相同.

所收集的任何错误消息都将与字段验证错误一起放入验证结果对象中,其关键字以 `validate` 选项对象中验证方法失败的键命名. 即便在任何时候每种模型验证方法都只有一个错误消息,但它会在数组中显示为单个字符串错误,以最大程度地提高与字段错误的一致性.

一个例子:

```js
class Place extends Model {}
Place.init({
  name: Sequelize.STRING,
  address: Sequelize.STRING,
  latitude: {
    type: DataTypes.INTEGER,
    validate: {
      min: -90,
      max: 90
    }
  },
  longitude: {
    type: DataTypes.INTEGER,
    validate: {
      min: -180,
      max: 180
    }
  },
}, {
  sequelize,
  validate: {
    bothCoordsOrNone() {
      if ((this.latitude === null) !== (this.longitude === null)) {
        throw new Error('Either both latitude and longitude, or neither!');
      }
    }
  }
})
```

在这种简单的情况下,如果只给定了纬度或经度,而不是同时给出两者, 则不能验证对象. 如果我们尝试构建一个超出范围的纬度且没有经度的对象,则`somePlace.validate()` 可能会返回：

```js
{
  'latitude': ['Invalid number: latitude'],
  'bothCoordsOrNone': ['Either both latitude and longitude, or neither!']
}
```

也可以使用在单个属性上定义的自定义验证程序(例如 `latitude` 属性,通过检查 `(value === null) !== (this.longitude === null)` )来完成此类验证, 但模型范围内的验证方法更为简洁.

# Raw Queries - 原始查询

由于常常使用简单的方式来执行原始/已经准备好的SQL查询,因此可以使用 `sequelize.query` 方法.

默认情况下,函数将返回两个参数 - 一个结果数组,以及一个包含元数据(例如受影响的行数等)的对象. 请注意,由于这是一个原始查询,所以元数据都是具体的方言. 某些方言返回元数据 "within" 结果对象(作为数组上的属性). 但是,将永远返回两个参数,但对于MSSQL和MySQL,它将是对同一对象的两个引用.

```js
const [results, metadata] = await sequelize.query("UPDATE users SET y = 42 WHERE x = 12");
// 结果将是一个空数组,元数据将包含受影响的行数.
```

在不需要访问元数据的情况下,你可以传递一个查询类型来告诉后续如何格式化结果. 例如,对于一个简单的选择查询你可以做:

```js
const { QueryTypes } = require('sequelize');
const users = await sequelize.query("SELECT * FROM `users`", { type: QueryTypes.SELECT });
// 我们不需要在这里分解结果 - 结果会直接返回
```

还有其他几种查询类型可用. [详细了解来源](https://github.com/sequelize/sequelize/blob/master/lib/query-types.js)

第二种选择是模型. 如果传递模型,返回的数据将是该模型的实例.

```js
// Callee 是模型定义. 这样你就可以轻松地将查询映射到预定义的模型
const projects = await sequelize.query('SELECT * FROM projects', {
  model: Projects,
  mapToModel: true // 如果你有任何映射字段,则在此处传递 true
});
// 现在,`projects` 的每个元素都是 Project 的一个实例
```

查看 [Query API 参考](https://github.com/demopark/sequelize-docs-Zh-CN/blob/master/core-concepts/class/lib/sequelize.js~Sequelize.html#instance-method-query)中的更多参数. 以下是一些例子:

```js
const { QueryTypes } = require('sequelize');
await sequelize.query('SELECT 1', {
  // 用于记录查询的函数(或false)
  // 将调用发送到服务器的每个SQL查询.
  logging: console.log,

  // 如果plain为true,则sequelize将仅返回结果集的第一条记录. 
  // 如果是false,它将返回所有记录.
  plain: false,

  // 如果你没有查询的模型定义,请将此项设置为true.
  raw: false,

  // 你正在执行的查询类型. 查询类型会影响结果在传回之前的格式.
  type: QueryTypes.SELECT
});

// 注意第二个参数为null！
// 即使我们在这里声明了一个被调用对象,
// raw: true 也会取代并返回一个原始对象.

console.log(await sequelize.query('SELECT * FROM projects', { raw: true }));
```

## "Dotted" 属性 和 `nest` 参数

如果表的属性名称包含点,则可以通过设置 `nest: true` 参数将生成的对象变为嵌套对象. 这可以通过 [dottie.js](https://github.com/mickhansen/dottie.js/) 在后台实现. 见下文：

- 不使用 `nest: true`:

  ```js
  const { QueryTypes } = require('sequelize');
  const records = await sequelize.query('select 1 as `foo.bar.baz`', {
    type: QueryTypes.SELECT
  });
  console.log(JSON.stringify(records[0], null, 2));
  ```

  ```js
  {
    "foo.bar.baz": 1
  }
  ```

- 使用 `nest: true`:

  ```js
  const { QueryTypes } = require('sequelize');
  const records = await sequelize.query('select 1 as `foo.bar.baz`', {
    nest: true,
    type: QueryTypes.SELECT
  });
  console.log(JSON.stringify(records[0], null, 2));
  ```

  ```js
  {
    "foo": {
      "bar": {
        "baz": 1
      }
    }
  }
  ```

## 替换

查询中的替换可以通过两种不同的方式完成:使用命名参数(以`:`开头),或者由`？`表示的未命名参数. 替换在options对象中传递.

- 如果传递一个数组, `?` 将按照它们在数组中出现的顺序被替换
- 如果传递一个对象, `:key` 将替换为该对象的键. 如果对象包含在查询中找不到的键,则会抛出异常,反之亦然.

```js
const { QueryTypes } = require('sequelize');

await sequelize.query(
  'SELECT * FROM projects WHERE status = ?',
  {
    replacements: ['active'],
    type: QueryTypes.SELECT
  }
);

await sequelize.query(
  'SELECT * FROM projects WHERE status = :status',
  {
    replacements: { status: 'active' },
    type: QueryTypes.SELECT
  }
);
```

数组替换将自动处理,以下查询将搜索状态与值数组匹配的项目.

```js
const { QueryTypes } = require('sequelize');

await sequelize.query(
  'SELECT * FROM projects WHERE status IN(:status)',
  {
    replacements: { status: ['active', 'inactive'] },
    type: QueryTypes.SELECT
  }
);
```

要使用通配符运算符 `％`,请将其附加到你的替换中. 以下查询与名称以 'ben' 开头的用户相匹配.

```js
const { QueryTypes } = require('sequelize');

await sequelize.query(
  'SELECT * FROM users WHERE name LIKE :search_name',
  {
    replacements: { search_name: 'ben%' },
    type: QueryTypes.SELECT
  }
);
```

## 绑定参数

绑定参数就像替换. 除非替换被转义并在查询发送到数据库之前通过后续插入到查询中,而将绑定参数发送到SQL查询文本之外的数据库. 查询可以具有绑定参数或替换.绑定参数由 `$1, $2, ... (numeric)` 或 `$key (alpha-numeric)` 引用.这是独立于方言的.

- 如果传递一个数组, `$1` 被绑定到数组中的第一个元素 (`bind[0]`).
- 如果传递一个对象, `$key` 绑定到 `object['key']`. 每个键必须以非数字字符开始. `$1` 不是一个有效的键,即使 `object['1']` 存在.
- 在这两种情况下 `$$` 可以用来转义一个 `$` 字符符号.

数组或对象必须包含所有绑定的值,或者Sequelize将抛出异常. 这甚至适用于数据库可能忽略绑定参数的情况.

数据库可能会增加进一步的限制. 绑定参数不能是SQL关键字,也不能是表或列名. 引用的文本或数据也忽略它们. 在PostgreSQL中,如果不能从上下文 `$1::varchar` 推断类型,那么也可能需要对其进行类型转换.

```js
const { QueryTypes } = require('sequelize');

await sequelize.query(
  'SELECT *, "text with literal $$1 and literal $$status" as t FROM projects WHERE status = $1',
  {
    bind: ['active'],
    type: QueryTypes.SELECT
  }
);

await sequelize.query(
  'SELECT *, "text with literal $$1 and literal $$status" as t FROM projects WHERE status = $status',
  {
    bind: { status: 'active' },
    type: QueryTypes.SELECT
  }
);
```

# Associations - 关联

Sequelize 支持标准关联关系: [一对一](https://en.wikipedia.org/wiki/One-to-one_(data_model)), [一对多](https://en.wikipedia.org/wiki/One-to-many_(data_model)) 和 [多对多](https://en.wikipedia.org/wiki/Many-to-many_(data_model)).

为此,Sequelize 提供了 **四种** 关联类型,并将它们组合起来以创建关联：

- `HasOne` 关联类型
- `BelongsTo` 关联类型
- `HasMany` 关联类型
- `BelongsToMany` 关联类型

该指南将讲解如何定义这四种类型的关联,然后讲解如何将它们组合来定义三种标准关联类型([一对一](https://en.wikipedia.org/wiki/One-to-one_(data_model)), [一对多](https://en.wikipedia.org/wiki/One-to-many_(data_model)) 和 [多对多](https://en.wikipedia.org/wiki/Many-to-many_(data_model))).

## 定义 Sequelize 关联

四种关联类型的定义非常相似. 假设我们有两个模型 `A` 和 `B`. 告诉 Sequelize 两者之间的关联仅需要调用一个函数：

```js
const A = sequelize.define('A', /* ... */);
const B = sequelize.define('B', /* ... */);

A.hasOne(B); // A 有一个 B
A.belongsTo(B); // A 属于 B
A.hasMany(B); // A 有多个 B
A.belongsToMany(B, { through: 'C' }); // A 属于多个 B , 通过联结表 C
```

它们都接受一个对象作为第二个参数(前三个参数是可选的,而对于包含 `through` 属性的 `belongsToMany` 是必需的)： They all accept an options object as a second parameter

```js
A.hasOne(B, { /* 参数 */ });
A.belongsTo(B, { /* 参数 */ });
A.hasMany(B, { /* 参数 */ });
A.belongsToMany(B, { through: 'C', /* 参数 */ });
```

关联的定义顺序是有关系的. 换句话说,对于这四种情况,定义顺序很重要. 在上述所有示例中,`A` 称为 **源** 模型,而 `B` 称为 **目标** 模型. 此术语很重要.

`A.hasOne(B)` 关联意味着 `A` 和 `B` 之间存在一对一的关系,外键在目标模型(`B`)中定义.

`A.belongsTo(B)`关联意味着 `A` 和 `B` 之间存在一对一的关系,外键在源模型中定义(`A`).

`A.hasMany(B)` 关联意味着 `A` 和 `B` 之间存在一对多关系,外键在目标模型(`B`)中定义.

这三个调用将导致 Sequelize 自动将外键添加到适当的模型中(除非它们已经存在).

`A.belongsToMany(B, { through: 'C' })` 关联意味着将表 `C` 用作[联结表](https://en.wikipedia.org/wiki/Associative_entity),在 `A` 和 `B` 之间存在多对多关系. 具有外键(例如,`aId` 和 `bId`). Sequelize 将自动创建此模型 `C`(除非已经存在),并在其上定义适当的外键.

*注意：在上面的 `belongsToMany` 示例中,字符串(`'C'`)被传递给 `through` 参数. 在这种情况下,Sequelize 会自动使用该名称生成模型. 但是,如果已经定义了模型,也可以直接传递模型.*

这些是每种关联类型中涉及的主要思想. 但是,这些关系通常成对使用,以便 Sequelize 更好地使用. 这将在后文中看到.

## 创建标准关系

如前所述,Sequelize 关联通常成对定义. 综上所述：

- 创建一个 **一对一** 关系, `hasOne` 和 `belongsTo` 关联一起使用;
- 创建一个 **一对多** 关系, `hasMany` he `belongsTo` 关联一起使用;
- 创建一个多对多关系, 两个belongsToMany调用一起使用.
  - 注意: 还有一个 *超级多对多* 关系,一次使用六个关联,将在[高级多对多关系指南](https://github.com/demopark/sequelize-docs-Zh-CN/blob/master/core-concepts/advanced-association-concepts/advanced-many-to-many.md)中进行讨论.

接下来将进行详细介绍. 本章末尾将讨论使用这些成对而不是单个关联的优点.

## 一对一关系

### 哲理

在深入探讨使用 Sequelize 的各个方面之前,退后一步来考虑一对一关系会发生什么是很有用的.

假设我们有两个模型,`Foo` 和 `Bar`.我们要在 Foo 和 Bar 之间建立一对一的关系.我们知道在关系数据库中,这将通过在其中一个表中建立外键来完成.因此,在这种情况下,一个非常关键的问题是：我们希望该外键在哪个表中？换句话说,我们是要 `Foo` 拥有 `barId` 列,还是 `Bar` 应当拥有 `fooId` 列？

原则上,这两个选择都是在 Foo 和 Bar 之间建立一对一关系的有效方法.但是,当我们说 *"Foo 和 Bar 之间存在一对一关系"* 时,尚不清楚该关系是 *强制性* 的还是可选的.换句话说,Foo 是否可以没有 Bar 而存在？ Foo 的 Bar 可以存在吗？这些问题的答案有助于帮我们弄清楚外键列在哪里.

### 目标

对于本示例的其余部分,我们假设我们有两个模型,即 `Foo` 和 `Bar`. 我们想要在它们之间建立一对一的关系,以便 `Bar` 获得 `fooId` 列.

### 实践

实现该目标的主要设置如下：

```js
Foo.hasOne(Bar);
Bar.belongsTo(Foo);
```

由于未传递任何参数,因此 Sequelize 将从模型名称中推断出要做什么. 在这种情况下,Sequelize 知道必须将 `fooId` 列添加到 `Bar` 中.

这样,在上述代码之后调用 `Bar.sync()` 将产生以下 SQL(例如,在PostgreSQL上)：

```js
CREATE TABLE IF NOT EXISTS "foos" (
  /* ... */
);
CREATE TABLE IF NOT EXISTS "bars" (
  /* ... */
  "fooId" INTEGER REFERENCES "foos" ("id") ON DELETE SET NULL ON UPDATE CASCADE
  /* ... */
);
```

### 参数

可以将各种参数作为关联调用的第二个参数传递.

#### `onDelete` 和 `onUpdate`

例如,要配置 `ON DELETE` 和 `ON UPDATE` 行为,你可以执行以下操作：

```js
Foo.hasOne(Bar, {
  onDelete: 'RESTRICT',
  onUpdate: 'RESTRICT'
});
Bar.belongsTo(Foo);
```

可用的参数为 `RESTRICT`, `CASCADE`, `NO ACTION`, `SET DEFAULT` 和 `SET NULL`.

一对一关联的默认值, `ON DELETE` 为 `SET NULL` 而 `ON UPDATE` 为 `CASCADE`.

#### 自定义外键

上面显示的 `hasOne` 和 `belongsTo` 调用都会推断出要创建的外键应称为 `fooId`. 如要使用其他名称,例如 `myFooId`：

```js
// 方法 1
Foo.hasOne(Bar, {
  foreignKey: 'myFooId'
});
Bar.belongsTo(Foo);

// 方法 2
Foo.hasOne(Bar, {
  foreignKey: {
    name: 'myFooId'
  }
});
Bar.belongsTo(Foo);

// 方法 3
Foo.hasOne(Bar);
Bar.belongsTo(Foo, {
  foreignKey: 'myFooId'
});

// 方法 4
Foo.hasOne(Bar);
Bar.belongsTo(Foo, {
  foreignKey: {
    name: 'myFooId'
  }
});
```

如上所示,`foreignKey` 参数接受一个字符串或一个对象. 当接收到一个对象时,该对象将用作列的定义,就像在标准的 `sequelize.define` 调用中所做的一样. 因此,指定诸如 `type`, `allowNull`, `defaultValue` 等参数就可以了.

例如,要使用 `UUID` 作为外键数据类型而不是默认值(`INTEGER`),只需执行以下操作：

```js
const { DataTypes } = require("Sequelize");

Foo.hasOne(Bar, {
  foreignKey: {
    // name: 'myFooId'
    type: DataTypes.UUID
  }
});
Bar.belongsTo(Foo);
```

#### 强制性与可选性关联

默认情况下,该关联被视为可选. 换句话说,在我们的示例中,`fooId` 允许为空,这意味着一个 Bar 可以不存在 Foo 而存在. 只需在外键选项中指定 `allowNull: false` 即可更改此设置：

```js
Foo.hasOne(Bar, {
  foreignKey: {
    allowNull: false
  }
});
// "fooId" INTEGER NOT NULL REFERENCES "foos" ("id") ON DELETE RESTRICT ON UPDATE RESTRICT
```

## 一对多关系

### 原理

一对多关联将一个源与多个目标连接,而所有这些目标仅与此单个源连接.

这意味着,与我们必须选择放置外键的一对一关联不同,在一对多关联中只有一个选项. 例如,如果一个 Foo 有很多 Bar(因此每个 Bar 都属于一个 Foo),那么唯一明智的方式就是在 `Bar` 表中有一个 `fooId` 列. 而反过来是不可能的,因为一个 `Foo` 会有很多 `Bar`.

### 目标

在这个例子中,我们有模型 `Team` 和 `Player`. 我们要告诉 Sequelize,他们之间存在一对多的关系,这意味着一个 Team 有 Player ,而每个 Player 都属于一个 Team.

### 实践

这样做的主要方法如下：

```js
Team.hasMany(Player);
Player.belongsTo(Team);
```

同样,实现此目标的主要方法是使用一对 Sequelize 关联(`hasMany` 和 `belongsTo`).

例如,在 PostgreSQL 中,以上设置将在 `sync()` 之后产生以下 SQL：

```js
CREATE TABLE IF NOT EXISTS "Teams" (
  /* ... */
);
CREATE TABLE IF NOT EXISTS "Players" (
  /* ... */
  "TeamId" INTEGER REFERENCES "Teams" ("id") ON DELETE SET NULL ON UPDATE CASCADE,
  /* ... */
);
```

### 参数

在这种情况下要应用的参数与一对一情况相同. 例如,要更改外键的名称并确保该关系是强制性的,我们可以执行以下操作：

```js
Team.hasMany(Player, {
  foreignKey: 'clubId'
});
Player.belongsTo(Team);
```

如同一对一关系, `ON DELETE` 默认为 `SET NULL` 而 `ON UPDATE` 默认为 `CASCADE`.

## 多对多关系

### 原理

多对多关联将一个源与多个目标相连,而所有这些目标又可以与第一个目标之外的其他源相连.

不能像其他关系那样通过向其中一个表添加一个外键来表示这一点. 取而代之的是使用[联结模型](https://en.wikipedia.org/wiki/Associative_entity)的概念. 这将是一个额外的模型(以及数据库中的额外表),它将具有两个外键列并跟踪关联. 联结表有时也称为 *join table* 或 *through table*.

### 目标

对于此示例,我们将考虑模型 `Movie` 和 `Actor`. 一位 actor 可能参与了许多 movies,而一部 movie 中有许多 actors 参与了其制作. 跟踪关联的联结表将被称为 `ActorMovies`,其中将包含外键 `movieId` 和 `actorId`.

### 实践

在 Sequelize 中执行此操作的主要方法如下：

```js
const Movie = sequelize.define('Movie', { name: DataTypes.STRING });
const Actor = sequelize.define('Actor', { name: DataTypes.STRING });
Movie.belongsToMany(Actor, { through: 'ActorMovies' });
Actor.belongsToMany(Movie, { through: 'ActorMovies' });
```

因为在 `belongsToMany` 的 `through` 参数中给出了一个字符串,所以 Sequelize 将自动创建 `ActorMovies` 模型作为联结模型. 例如,在 PostgreSQL 中：

```js
CREATE TABLE IF NOT EXISTS "ActorMovies" (
  "createdAt" TIMESTAMP WITH TIME ZONE NOT NULL,
  "updatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
  "MovieId" INTEGER REFERENCES "Movies" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "ActorId" INTEGER REFERENCES "Actors" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  PRIMARY KEY ("MovieId","ActorId")
);
```

除了字符串以外,还支持直接传递模型,在这种情况下,给定的模型将用作联结模型(并且不会自动创建任何模型). 例如：

```js
const Movie = sequelize.define('Movie', { name: DataTypes.STRING });
const Actor = sequelize.define('Actor', { name: DataTypes.STRING });
const ActorMovies = sequelize.define('ActorMovies', {
  MovieId: {
    type: DataTypes.INTEGER,
    references: {
      model: Movie, // 'Movies' 也可以使用
      key: 'id'
    }
  },
  ActorId: {
    type: DataTypes.INTEGER,
    references: {
      model: Actor, // 'Actors' 也可以使用
      key: 'id'
    }
  }
});
Movie.belongsToMany(Actor, { through: 'ActorMovies' });
Actor.belongsToMany(Movie, { through: 'ActorMovies' });
```

上面的代码在 PostgreSQL 中产生了以下 SQL,与上面所示的代码等效：

```js
CREATE TABLE IF NOT EXISTS "ActorMovies" (
  "MovieId" INTEGER NOT NULL REFERENCES "Movies" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "ActorId" INTEGER NOT NULL REFERENCES "Actors" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
  "createdAt" TIMESTAMP WITH TIME ZONE NOT NULL,
  "updatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
  UNIQUE ("MovieId", "ActorId"),     -- 注意: Sequelize 产生了这个 UNIQUE 约束,但是
  PRIMARY KEY ("MovieId","ActorId")  -- 这没有关系,因为它也是 PRIMARY KEY
);
```

### 参数

与一对一和一对多关系不同,对于多对多关系,`ON UPDATE` 和 `ON DELETE` 的默认值为 `CASCADE`.

## 基本的涉及关联的查询

了解了定义关联的基础知识之后,我们可以查看涉及关联的查询. 最常见查询是 *read* 查询(即 SELECT). 稍后,将展示其他类型的查询.

为了研究这一点,我们将思考一个例子,其中有船和船长,以及它们之间的一对一关系. 我们将在外键上允许 null(默认值),这意味着船可以在没有船长的情况下存在,反之亦然.

```js
// 这是我们用于以下示例的模型的设置
const Ship = sequelize.define('ship', {
  name: DataTypes.TEXT,
  crewCapacity: DataTypes.INTEGER,
  amountOfSails: DataTypes.INTEGER
}, { timestamps: false });
const Captain = sequelize.define('captain', {
  name: DataTypes.TEXT,
  skillLevel: {
    type: DataTypes.INTEGER,
    validate: { min: 1, max: 10 }
  }
}, { timestamps: false });
Captain.hasOne(Ship);
Ship.belongsTo(Captain);
```

### 获取关联 - 预先加载 vs 延迟加载

预先加载和延迟加载的概念是理解获取关联如何在 Sequelize 中工作的基础. 延迟加载是指仅在确实需要时才获取关联数据的技术. 另一方面,预先加载是指从一开始就通过较大的查询一次获取所有内容的技术.

#### 延迟加载示例

```js
const awesomeCaptain = await Captain.findOne({
  where: {
    name: "Jack Sparrow"
  }
});
// 用获取到的 captain 做点什么
console.log('Name:', awesomeCaptain.name);
console.log('Skill Level:', awesomeCaptain.skillLevel);
// 现在我们需要有关他的 ship 的信息!
const hisShip = await awesomeCaptain.getShip();
// 用 ship 做点什么
console.log('Ship Name:', hisShip.name);
console.log('Amount of Sails:', hisShip.amountOfSails);
```

请注意,在上面的示例中,我们进行了两个查询,仅在要使用它时才获取关联的 ship. 如果我们可能需要也可能不需要这艘 ship,或者我们只想在少数情况下有条件地取回它,这会特别有用; 这样,我们可以仅在必要时提取,从而节省时间和内存.

注意：上面使用的 `getShip()` 实例方法是 Sequelize 自动添加到 Captain 实例的方法之一. 还有其他方法, 你将在本指南的后面部分进一步了解它们.

#### 预先加载示例

```js
const awesomeCaptain = await Captain.findOne({
  where: {
    name: "Jack Sparrow"
  },
  include: Ship
});
// 现在 ship 跟着一起来了
console.log('Name:', awesomeCaptain.name);
console.log('Skill Level:', awesomeCaptain.skillLevel);
console.log('Ship Name:', awesomeCaptain.ship.name);
console.log('Amount of Sails:', awesomeCaptain.ship.amountOfSails);
```

如上所示,通过使用 include 参数 在 Sequelize 中执行预先加载. 观察到这里只对数据库执行了一个查询(与实例一起带回关联的数据).

这只是 Sequelize 中预先加载的简单介绍. 还有更多内容,你可以在[预先加载的专用指南](https://github.com/demopark/sequelize-docs-Zh-CN/blob/master/core-concepts/advanced-association-concepts/eager-loading.md)中学习

### 创建, 更新和删除

上面显示了查询有关关联的数据的基础知识. 对于创建,更新和删除,你可以：

- 直接使用标准模型查询：

  ```js
  // 示例：使用标准方法创建关联的模型
  Bar.create({
    name: 'My Bar',
    fooId: 5
  });
  // 这将创建一个属于 ID 5 的 Foo 的 Bar
  // 这里没有什么特别的东西
  ```

- 或使用关联模型可用的 *[特殊方法/混合](https://github.com/demopark/sequelize-docs-Zh-CN/blob/master/core-concepts/＃special-methods-mixins-to-instances)* ,这将在本文稍后进行解释.

**注意:** [`save()`实例方法](https://sequelize.org/master/class/lib/model.js~Model.html#instance-method-save) 并不知道关联关系. 如果你修改了 *父级* 对象预先加载的 *子级* 的值,那么在父级上调用 `save()` 将会忽略子级上发生的修改.

## 关联别名 & 自定义外键

在以上所有示例中,Sequelize 自动定义了外键名称. 例如,在船和船长示例中,Sequelize 在 Ship 模型上自动定义了一个 `captainId` 字段. 然而,想要自定义外键也是很容易的.

让我们以简化的形式考虑 Ship 和 Captain 模型,仅着眼于当前主题,如下所示(较少的字段)：

```js
const Ship = sequelize.define('ship', { name: DataTypes.TEXT }, { timestamps: false });
const Captain = sequelize.define('captain', { name: DataTypes.TEXT }, { timestamps: false });
```

有三种方法可以为外键指定不同的名称：

- 通过直接提供外键名称
- 通过定义别名
- 通过两个方法同时进行

### 回顾: 默认设置

通过简单地使用 `Ship.belongsTo(Captain)`,sequelize 将自动生成外键名称：

```js
Ship.belongsTo(Captain); // 这将在 Ship 中创建 `captainId` 外键.

// 通过将模型传递给 `include` 来完成预先加载:
console.log((await Ship.findAll({ include: Captain })).toJSON());
// 或通过提供关联的模型名称:
console.log((await Ship.findAll({ include: 'captain' })).toJSON());

// 同样,实例获得用于延迟加载的 `getCaptain()` 方法：
const ship = Ship.findOne();
console.log((await ship.getCaptain()).toJSON());
```

### 直接提供外键名称

可以直接在关联定义的参数中提供外键名称,如下所示：

```js
Ship.belongsTo(Captain, { foreignKey: 'bossId' }); // 这将在 Ship 中创建 `bossId` 外键.

// 通过将模型传递给 `include` 来完成预先加载:
console.log((await Ship.findAll({ include: Captain })).toJSON());
// 或通过提供关联的模型名称:
console.log((await Ship.findAll({ include: 'Captain' })).toJSON());

// 同样,实例获得用于延迟加载的 `getCaptain()` 方法:
const ship = Ship.findOne();
console.log((await ship.getCaptain()).toJSON());
```

### 定义别名

定义别名比简单指定外键的自定义名称更强大. 通过一个示例可以更好地理解这一点：

```js
Ship.belongsTo(Captain, { as: 'leader' }); // 这将在 Ship 中创建 `leaderId` 外键.

// 通过将模型传递给 `include` 不能再触发预先加载:
console.log((await Ship.findAll({ include: Captain })).toJSON()); // 引发错误
// 相反,你必须传递别名:
console.log((await Ship.findAll({ include: 'leader' })).toJSON());
// 或者,你可以传递一个指定模型和别名的对象:
console.log((await Ship.findAll({
  include: {
    model: Captain,
    as: 'leader'
  }
})).toJSON());

// 同样,实例获得用于延迟加载的 `getLeader()`方法:
const ship = Ship.findOne();
console.log((await ship.getLeader()).toJSON());
```

当你需要在同一模型之间定义两个不同的关联时,别名特别有用. 例如,如果我们有`Mail` 和 `Person` 模型,则可能需要将它们关联两次,以表示邮件的 `sender` 和 `receiver`. 在这种情况下,我们必须为每个关联使用别名,因为否则,诸如 `mail.getPerson()` 之类的调用将是模棱两可的. 使用 `sender` 和 `receiver` 别名,我们将有两种可用的可用方法：`mail.getSender()` 和 `mail.getReceiver()`,它们都返回一个`Promise<Person>`.

在为 `hasOne` 或 `belongsTo` 关联定义别名时,应使用单词的单数形式(例如上例中的 `leader`). 另一方面,在为 `hasMany` 和 `belongsToMany` 定义别名时,应使用复数形式. [高级多对多关联指南](https://github.com/demopark/sequelize-docs-Zh-CN/blob/master/core-concepts/advanced-association-concepts/advanced-many-to-many.md)中介绍了定义多对多关系(带有`belongsToMany`)的别名.

### 两者都做

我们可以定义别名,也可以直接定义外键:

```js
Ship.belongsTo(Captain, { as: 'leader', foreignKey: 'bossId' }); // 这将在 Ship 中创建 `bossId` 外键.

// 由于定义了别名,因此仅通过将模型传递给 `include`,预先加载将不起作用:
console.log((await Ship.findAll({ include: Captain })).toJSON()); // 引发错误
// 相反,你必须传递别名:
console.log((await Ship.findAll({ include: 'leader' })).toJSON());
// 或者,你可以传递一个指定模型和别名的对象:
console.log((await Ship.findAll({
  include: {
    model: Captain,
    as: 'leader'
  }
})).toJSON());

// 同样,实例获得用于延迟加载的 `getLeader()` 方法:
const ship = Ship.findOne();
console.log((await ship.getLeader()).toJSON());
```

## 添加到实例的特殊方法

当两个模型之间定义了关联时,这些模型的实例将获得特殊的方法来与其关联的另一方进行交互.

例如,如果我们有两个模型 `Foo` 和 `Bar`,并且它们是关联的,则它们的实例将具有以下可用的方法,具体取决于关联类型：

### `Foo.hasOne(Bar)`

- `fooInstance.getBar()`
- `fooInstance.setBar()`
- `fooInstance.createBar()`

示例:

```js
const foo = await Foo.create({ name: 'the-foo' });
const bar1 = await Bar.create({ name: 'some-bar' });
const bar2 = await Bar.create({ name: 'another-bar' });
console.log(await foo.getBar()); // null
await foo.setBar(bar1);
console.log((await foo.getBar()).name); // 'some-bar'
await foo.createBar({ name: 'yet-another-bar' });
const newlyAssociatedBar = await foo.getBar();
console.log(newlyAssociatedBar.name); // 'yet-another-bar'
await foo.setBar(null); // Un-associate
console.log(await foo.getBar()); // null
```

### `Foo.belongsTo(Bar)`

来自 `Foo.hasOne(Bar)` 的相同内容:

- `fooInstance.getBar()`
- `fooInstance.setBar()`
- `fooInstance.createBar()`

### `Foo.hasMany(Bar)`

- `fooInstance.getBars()`
- `fooInstance.countBars()`
- `fooInstance.hasBar()`
- `fooInstance.hasBars()`
- `fooInstance.setBars()`
- `fooInstance.addBar()`
- `fooInstance.addBars()`
- `fooInstance.removeBar()`
- `fooInstance.removeBars()`
- `fooInstance.createBar()`

示例:

```js
const foo = await Foo.create({ name: 'the-foo' });
const bar1 = await Bar.create({ name: 'some-bar' });
const bar2 = await Bar.create({ name: 'another-bar' });
console.log(await foo.getBars()); // []
console.log(await foo.countBars()); // 0
console.log(await foo.hasBar(bar1)); // false
await foo.addBars([bar1, bar2]);
console.log(await foo.countBars()); // 2
await foo.addBar(bar1);
console.log(await foo.countBars()); // 2
console.log(await foo.hasBar(bar1)); // true
await foo.removeBar(bar2);
console.log(await foo.countBars()); // 1
await foo.createBar({ name: 'yet-another-bar' });
console.log(await foo.countBars()); // 2
await foo.setBars([]); // 取消关联所有先前关联的 Bars
console.log(await foo.countBars()); // 0
```

getter 方法接受参数,就像通常的 finder 方法(例如`findAll`)一样：

```js
const easyTasks = await project.getTasks({
  where: {
    difficulty: {
      [Op.lte]: 5
    }
  }
});
const taskTitles = (await project.getTasks({
  attributes: ['title'],
  raw: true
})).map(task => task.title);
```

### `Foo.belongsToMany(Bar, { through: Baz })`

来自 `Foo.hasMany(Bar)` 的相同内容:

- `fooInstance.getBars()`
- `fooInstance.countBars()`
- `fooInstance.hasBar()`
- `fooInstance.hasBars()`
- `fooInstance.setBars()`
- `fooInstance.addBar()`
- `fooInstance.addBars()`
- `fooInstance.removeBar()`
- `fooInstance.removeBars()`
- `fooInstance.createBar()`

### 注意: 方法名称

如上面的示例所示,Sequelize 赋予这些特殊方法的名称是由前缀(例如,get,add,set)和模型名称(首字母大写)组成的. 必要时,可以使用复数形式,例如在 `fooInstance.setBars()` 中. 同样,不规则复数也由 Sequelize 自动处理. 例如,`Person` 变成 `People` 或者 `Hypothesis` 变成 `Hypotheses`.

如果定义了别名,则将使用别名代替模型名称来形成方法名称. 例如：

```js
Task.hasOne(User, { as: 'Author' });
```

- `taskInstance.getAuthor()`
- `taskInstance.setAuthor()`
- `taskInstance.createAuthor()`

## 为什么关联是成对定义的？

如前所述,就像上面大多数示例中展示的,Sequelize 中的关联通常成对定义：

- 创建一个 **一对一** 关系, `hasOne` 和 `belongsTo` 关联一起使用;
- 创建一个 **一对多** 关系, `hasMany` he `belongsTo` 关联一起使用;
- 创建一个 **多对多** 关系, 两个 `belongsToMany` 调用一起使用.

当在两个模型之间定义了 Sequelize 关联时,只有 *源* 模型 *知晓关系*. 因此,例如,当使用 `Foo.hasOne(Bar)`(当前,`Foo` 是源模型,而 `Bar` 是目标模型)时,只有 `Foo` 知道该关联的存在. 这就是为什么在这种情况下,如上所示,`Foo` 实例获得方法 `getBar()`, `setBar()` 和 `createBar()` 而另一方面,`Bar` 实例却没有获得任何方法.

类似地,对于 `Foo.hasOne(Bar)`,由于 `Foo` 了解这种关系,我们可以像 `Foo.findOne({ include: Bar })` 中那样执行预先加载,但不能执行 `Bar.findOne({ include: Foo })`.

因此,为了充分发挥 Sequelize 的作用,我们通常成对设置关系,以便两个模型都 *互相知晓*.

实际示范:

- 如果我们未定义关联对,则仅调用 `Foo.hasOne(Bar)`:

  ```js
  // 这有效...
  await Foo.findOne({ include: Bar });
  
  // 但这会引发错误:
  await Bar.findOne({ include: Foo });
  // SequelizeEagerLoadingError: foo is not associated to bar!
  ```

- 如果我们按照建议定义关联对, 即, `Foo.hasOne(Bar)` 和 `Bar.belongsTo(Foo)`:

  ```js
  // 这有效
  await Foo.findOne({ include: Bar });
  
  // 这也有效!
  await Bar.findOne({ include: Foo });
  ```

## 涉及相同模型的多个关联

在 Sequelize 中,可以在同一模型之间定义多个关联. 你只需要为它们定义不同的别名：

```js
Team.hasOne(Game, { as: 'HomeTeam', foreignKey: 'homeTeamId' });
Team.hasOne(Game, { as: 'AwayTeam', foreignKey: 'awayTeamId' });
Game.belongsTo(Team);
```

## 创建引用非主键字段的关联

在以上所有示例中,通过引用所涉及模型的主键(在我们的示例中为它们的ID)定义了关联. 但是,Sequelize 允许你定义一个关联,该关联使用另一个字段而不是主键字段来建立关联.

此其他字段必须对此具有唯一的约束(否则,这将没有意义).

### 对于 `belongsTo` 关系

首先,回想一下 `A.belongsTo(B)` 关联将外键放在 *源模型* 中(即,在 `A` 中).

让我们再次使用"船和船长"的示例. 此外,我们将假定船长姓名是唯一的：

```js
const Ship = sequelize.define('ship', { name: DataTypes.TEXT }, { timestamps: false });
const Captain = sequelize.define('captain', {
  name: { type: DataTypes.TEXT, unique: true }
}, { timestamps: false });
```

这样,我们不用在 Ship 上保留 `captainId`,而是可以保留 `captainName` 并将其用作关联跟踪器. 换句话说,我们的关系将引用目标模型上的另一列：`name` 列,而不是从目标模型(Captain)中引用 `id`. 为了说明这一点,我们必须定义一个 *目标键*. 我们还必须为外键本身指定一个名称：

```js
Ship.belongsTo(Captain, { targetKey: 'name', foreignKey: 'captainName' });
// 这将在源模型(Ship)中创建一个名为 `captainName` 的外键,
// 该外键引用目标模型(Captain)中的 `name` 字段.
```

现在我们可以做类似的事情:

```js
await Captain.create({ name: "Jack Sparrow" });
const ship = await Ship.create({ name: "Black Pearl", captainName: "Jack Sparrow" });
console.log((await ship.getCaptain()).name); // "Jack Sparrow"
```

### 对于 `hasOne` 和 `hasMany` 关系

可以将完全相同的想法应用于 `hasOne` 和 `hasMany` 关联,但是在定义关联时,我们提供了 `sourceKey`,而不是提供 `targetKey`. 这是因为与 `belongsTo` 不同,`hasOne` 和 `hasMany`关联将外键保留在目标模型上：

```js
const Foo = sequelize.define('foo', {
  name: { type: DataTypes.TEXT, unique: true }
}, { timestamps: false });
const Bar = sequelize.define('bar', {
  title: { type: DataTypes.TEXT, unique: true }
}, { timestamps: false });
const Baz = sequelize.define('baz', { summary: DataTypes.TEXT }, { timestamps: false });
Foo.hasOne(Bar, { sourceKey: 'name', foreignKey: 'fooName' });
Bar.hasMany(Baz, { sourceKey: 'title', foreignKey: 'barTitle' });
// [...]
await Bar.setFoo("Foo's Name Here");
await Baz.addBar("Bar's Title Here");
```

### 对于 `belongsToMany` 关系

同样的想法也可以应用于 `belongsToMany` 关系. 但是,与其他情况下(其中只涉及一个外键)不同,`belongsToMany` 关系涉及两个外键,这些外键保留在额外的表(联结表)上.

请考虑以下设置：

```js
const Foo = sequelize.define('foo', {
  name: { type: DataTypes.TEXT, unique: true }
}, { timestamps: false });
const Bar = sequelize.define('bar', {
  title: { type: DataTypes.TEXT, unique: true }
}, { timestamps: false });
```

有四种情况需要考虑：

- 我们可能希望使用默认的主键为 `Foo` 和 `Bar` 进行多对多关系：

```js
Foo.belongsToMany(Bar, { through: 'foo_bar' });
// 这将创建具有字段 `fooId` 和 `barID` 的联结表 `foo_bar`.
```

- 我们可能希望使用默认主键 `Foo` 的多对多关系,但使用 `Bar` 的不同字段：

```js
Foo.belongsToMany(Bar, { through: 'foo_bar', targetKey: 'title' });
// 这将创建具有字段 `fooId` 和 `barTitle` 的联结表 `foo_bar`.
```

- 我们可能希望使用 `Foo` 的不同字段和 `Bar` 的默认主键进行多对多关系：

```js
Foo.belongsToMany(Bar, { through: 'foo_bar', sourceKey: 'name' });
// 这将创建具有字段 `fooName` 和 `barId` 的联结表 `foo_bar`.
```

- 我们可能希望使用不同的字段为 `Foo` 和 `Bar` 使用多对多关系：

```js
Foo.belongsToMany(Bar, { through: 'foo_bar', sourceKey: 'name', targetKey: 'title' });
// 这将创建带有字段 `fooName` 和 `barTitle` 的联结表 `foo_bar`.
```

### 注意

不要忘记关联中引用的字段必须具有唯一性约束. 否则,将引发错误(对于 SQLite 有时还会发出诡异的错误消息,例如 `SequelizeDatabaseError: SQLITE_ERROR: foreign key mismatch - "ships" referencing "captains"`).

在 `sourceKey` 和 `targetKey` 之间做出决定的技巧只是记住每个关系在何处放置其外键. 如本指南开头所述：

- `A.belongsTo(B)` 将外键保留在源模型中(`A`),因此引用的键在目标模型中,因此使用了 `targetKey`.
- `A.hasOne(B)` 和 `A.hasMany(B)` 将外键保留在目标模型(`B`)中,因此引用的键在源模型中,因此使用了 `sourceKey`.
- `A.belongsToMany(B)` 包含一个额外的表(联结表),因此 `sourceKey` 和 `targetKey` 均可用,其中 `sourceKey` 对应于`A`(源)中的某个字段而 `targetKey` 对应于 `B`(目标)中的某个字段.