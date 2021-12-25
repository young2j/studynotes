/*
 * File: context.hpp
 * Created Date: 2021-12-19 11:42:17
 * Author: ysj
 * Description:  状态上下文
 */
#pragma once
#include <iostream>
#include "state.hpp"
using namespace std;

class OrderContext
{
private:
    OrderState *state;

public:
    bool hasPaid;
    bool elapsed30min;

public:
    // 初始化
    OrderContext(OrderState *state)
    {
        this->state = state;
        this->hasPaid = false;
        this->elapsed30min = false;
    };
    // 获取当前状态
    OrderState *getState()
    {
        return this->state;
    };
    //设置当前状态
    void setState(OrderState *state)
    {
        this->state = state;
    };
    // 当前状态下的行为
    void handle()
    {
        state->handle(this);
    };
};