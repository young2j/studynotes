/**
 * -------------------------------------------------------
 * File: state.ts
 * Created Date: 2021-12-19 07:55:20
 * Author: ysj
 * Description: 状态类
 * -------------------------------------------------------
 */

import OrderContext from './context';

// 订单状态接口
export interface IOrderState {
  handle(ctx: OrderContext): void;
}

// 订单支付中
export class OrderPaying implements IOrderState {
  handle(ctx: OrderContext): void {
    console.log('订单支付中...');
    if (ctx.hasPaid) {
      ctx.setState(new OrderPaid());
      ctx.handle();
    } else if (ctx.elapsed30min) {
      ctx.setState(new OrderExpired());
      ctx.handle();
    } else {
      ctx.setState(new OrderUnpay());
      ctx.handle();
    }
  }
}

// 订单未支付
export class OrderUnpay implements IOrderState {
  handle(ctx: OrderContext): void {
    if (ctx.hasPaid) {
      ctx.setState(new OrderPaid());
      ctx.handle();
    } else if (ctx.elapsed30min) {
      ctx.setState(new OrderExpired());
      ctx.handle();
    } else {
      console.log('订单未支付...');
    }
  }
}

// 订单已支付
export class OrderPaid implements IOrderState {
  handle(ctx: OrderContext): void {
    console.log('订单已支付。');
  }
}

// 订单已过期
export class OrderExpired implements IOrderState {
  handle(ctx: OrderContext): void {
    console.log('订单已过期。');
  }
}
