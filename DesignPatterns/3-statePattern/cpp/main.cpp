/*
 * File: main.cpp
 * Created Date: 2021-12-19 11:40:30
 * Author: ysj
 * Description:  cpp状态模式
 */
#include <iostream>
#include "state.hpp"
#include "context.hpp"

int main()
{
    OrderState *initState = new OrderPaying();
    // 支付中->未支付
    OrderContext *ctx = new OrderContext(initState);
    ctx->handle();

    // 未支付->已过期
    ctx->elapsed30min = true;
    ctx->handle();

    // 已过期->X 已支付
    ctx->hasPaid = true;
    ctx->handle();

    // 支付中-> 已支付
    ctx->setState(initState);
    ctx->hasPaid = true;
    ctx->handle();

    return 0;
}
