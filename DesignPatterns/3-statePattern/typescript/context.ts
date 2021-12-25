/**
 * -------------------------------------------------------
 * File: context.ts
 * Created Date: 2021-12-19 07:55:04
 * Author: ysj
 * Description: 状态上下文类
 * -------------------------------------------------------
 */

import { IOrderState } from './state';

export default class OrderContext {
  private _state: IOrderState;

  public hasPaid: boolean;
  public elapsed30min: boolean;
  // 初始化
  constructor(state: IOrderState) {
    this._state = state;
    this.hasPaid = false;
    this.elapsed30min = false;
  }

  // 获取当前状态
  getState(): IOrderState {
    return this._state;
  }
  // 设置当前状态
  setState(state: IOrderState) {
    this._state = state;
  }
  // 当前状态下的行为
  handle() {
    this._state.handle(this);
  }
}
