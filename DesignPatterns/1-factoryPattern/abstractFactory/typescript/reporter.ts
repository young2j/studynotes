/**
 * -------------------------------------------------------
 * File: reporter.ts
 * Created Date: 2021-11-29 04:19:29
 * Author: ysj
 * Description: 报告类
 * -------------------------------------------------------
 */

/**一个接口 */
export interface IReporter {
  reportType: string;
  exportReport(): void;
}

/**具体实现类 */
export class TrialReporter implements IReporter {
  reportType = '简要excel报告';
  exportReport() {
    console.log(`导出${this.reportType}`);
  }
}

/**具体实现类 */
export class BasicReporter implements IReporter {
  reportType = '详细excel报告';
  exportReport() {
    console.log(`导出${this.reportType}`);
  }
}

/**具体实现类 */
export class PremiumReporter implements IReporter {
  reportType = '精美pdf报告';
  exportReport() {
    console.log(`导出${this.reportType}`);
  }
}
