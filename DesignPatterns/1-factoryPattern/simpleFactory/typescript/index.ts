/**
 * -------------------------------------------------------
 * File: index.ts
 * Created Date: 2021-11-24 01:42:10
 * Author: ysj
 * Description: ts 简单工厂
 * -------------------------------------------------------
 */

/**一个接口 */
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
/**简单工厂类*/
class SimpleFactory {
  createCombo(comboType: string) {
    switch (comboType) {
      case 'trial':
        return new TrialCombo();
      case 'basic':
        return new BasicCombo();
      case 'premium':
        return new PremiumCombo();
      default:
        return new TrialCombo();
    }
  }
}

/**客户端调用 */
const factory: SimpleFactory = new SimpleFactory();

const trialCombo = factory.createCombo('trial');
trialCombo.task();
const basicCombo = factory.createCombo('basic');
basicCombo.task();
const premiumCombo = factory.createCombo('premium');
premiumCombo.task();
