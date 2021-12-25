/**
 * -------------------------------------------------------
 * File: factory.ts
 * Created Date: 2021-11-29 04:20:16
 * Author: ysj
 * Description: 工厂类
 * -------------------------------------------------------
 */

import { ICombo, TrialCombo, BasicCombo, PremiumCombo } from './combo';
import {
  IReporter,
  TrialReporter,
  BasicReporter,
  PremiumReporter,
} from './reporter';

/**工厂接口*/
export interface IFactory {
  createCombo(): ICombo;
  createReporter(): IReporter;
}

/**具体工厂类 */
export class TrialFactory implements IFactory {
  createCombo() {
    return new TrialCombo();
  }
  createReporter() {
    return new TrialReporter();
  }
}

/**具体工厂类 */
export class BasicFactory implements IFactory {
  createCombo() {
    return new BasicCombo();
  }
  createReporter() {
    return new BasicReporter();
  }
}

/**具体工厂类 */
export class PremiumFactory implements IFactory {
  createCombo() {
    return new PremiumCombo();
  }
  createReporter() {
    return new PremiumReporter();
  }
}
