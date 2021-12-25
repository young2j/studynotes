/*
 * File: state.hpp
 * Created Date: 2021-12-19 11:42:48
 * Author: ysj
 * Description:  状态类
 */

#pragma once
#include <iostream>
#include "context.hpp"

using namespace std;

// 抽象订单状态类
class OrderState
{
public:
    virtual void handle(OrderContext *ctx) = 0;
};

/* 订单支付中 */
class OrderPaying : public OrderState
{
public:
    void handle(OrderContext *ctx)
    {
        cout << "订单支付中..." << endl;
        if (ctx->hasPaid)
        {
            ctx->setState(new OrderPaid());
            ctx->handle();
        }
        else if (ctx->elapsed30min)
        {
            ctx->setState(new OrderExpired());
            ctx->handle();
        }
        else
        {
            ctx->setState(new OrderUnpay());
            ctx->handle();
        }
    }
};

/* 订单未支付 */
class OrderUnpay : public OrderState
{
public:
    void handle(OrderContext *ctx)
    {
        if (ctx->hasPaid)
        {
            ctx->setState(new OrderPaid());
            ctx->handle();
        }
        else if (ctx->elapsed30min)
        {
            ctx->setState(new OrderExpired());
            ctx->handle();
        }
        else
        {
            cout << "订单未支付..." << endl;
        }
    }
};

/* 订单已支付 */
class OrderPaid : public OrderState
{
public:
    void handle(OrderContext *ctx)
    {
        cout << "订单已支付。" << endl;
    }
};

/* 订单已过期 */
class OrderExpired : public OrderState
{
public:
    void handle(OrderContext *ctx)
    {
        cout << "订单已过期。" << endl;
    }
};