/**
 * -------------------------------------------------------
 * File: combo.ts
 * Created Date: 2021-11-29 04:18:18
 * Author: ysj
 * Description: 套餐类
 * -------------------------------------------------------
 */

/**一个接口 */
export interface ICombo {
  name: string;
  price: number;
  task(): void;
}

/**具体实现类*/
export class TrialCombo implements ICombo {
  name = '体验版';
  price = 299;
  task() {
    console.log(`${this.name}-> 执行价值${this.price}元的服务`);
  }
}

/**具体实现类*/
export class BasicCombo implements ICombo {
  name = '基础版';
  price = 599;
  task() {
    console.log(`${this.name}-> 执行价值${this.price}元的服务`);
  }
}
/**具体实现类*/
export class PremiumCombo implements ICombo {
  name = '高级版';
  price = 1999;
  task() {
    console.log(`${this.name}-> 执行价值${this.price}元的服务`);
  }
}
