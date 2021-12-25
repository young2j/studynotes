/**
 * -------------------------------------------------------
 * File: index.ts
 * Created Date: 2021-12-19 07:55:33
 * Author: ysj
 * Description: ts 状态模式
 * -------------------------------------------------------
 */
import { OrderPaying } from './state';
import OrderContext from './context';

const initState = new OrderPaying();
// 支付中->未支付
const ctx = new OrderContext(initState);
ctx.handle();

// 未支付->已过期
ctx.elapsed30min = true;
ctx.handle();

// 已过期->X 已支付
ctx.hasPaid = true;
ctx.handle();

// 支付中-> 已支付
ctx.setState(initState);
ctx.hasPaid = true;
ctx.handle();
