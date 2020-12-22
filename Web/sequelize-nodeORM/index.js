const { Sequelize, DataTypes,Model} = require('sequelize')
const sequelize = new Sequelize('sqlite::memory:',{
    define:{
        freezeTableName:true //强制所有实例表名与模型名相同
    }
})

const User = sequelize.define('User',{
    firstName:{
        type:DataTypes.STRING,
        allowNull:false
    },
    lastName:{
        type:DataTypes.STRING,
        allowNull:false
    },
    adult:{
        type:DataTypes.BOOLEAN,
        allowNull:true //默认
    }
},{
    freezeTableName:true, //强制表名与模型名相同
})

console.log(User === sequelize.models.User)

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
    adult:{
        type:DataTypes.BOOLEAN,
        allowNull:true //默认
    }
},{
    sequelize, //需要传递连接实例
    modelName:'User'
})

console.log(User === sequelize.models.User)
