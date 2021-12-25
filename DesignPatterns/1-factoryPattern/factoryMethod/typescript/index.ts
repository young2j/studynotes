/**
 * -------------------------------------------------------
 * File: index.ts
 * Created Date: 2021-11-24 01:42:10
 * Author: ysj
 * Description: ts 工厂方法
 * -------------------------------------------------------
 */

/**接口 */
interface ICombo {
  name: string;
  price: number;
  task(): void;
}

/**具体实现类*/
class TrialCombo implements ICombo {
  name = '体验版';
  price = 299;
  task() {
    console.log(`${this.name}-> 执行价值${this.price}元的服务`);
  }
}

/**具体实现类*/
class BasicCombo implements ICombo {
  name = '基础版';
  price = 599;
  task() {
    console.log(`${this.name}-> 执行价值${this.price}元的服务`);
  }
}
/**具体实现类*/
class PremiumCombo implements ICombo {
  name = '高级版';
  price = 1999;
  task() {
    console.log(`${this.name}-> 执行价值${this.price}元的服务`);
  }
}
/**工厂接口*/
interface IFactory {
  createCombo(): ICombo;
}

/**具体工厂类 */
class TrialComboFactory implements IFactory {
  createCombo() {
    return new TrialCombo();
  }
}

/**具体工厂类 */
class BasicComboFactory implements IFactory {
  createCombo() {
    return new BasicCombo();
  }
}

/**具体工厂类 */
class PremiumComboFactory implements IFactory {
  createCombo() {
    return new PremiumCombo();
  }
}

/**客户端调用 */
const trialComboFactory = new TrialComboFactory();
const trialCombo = trialComboFactory.createCombo();
trialCombo.task();

const basicComboFactory = new BasicComboFactory();
const basicCombo = basicComboFactory.createCombo();
basicCombo.task();

const premiumComboFactory = new PremiumComboFactory();
const premiumCombo = premiumComboFactory.createCombo();
premiumCombo.task();
