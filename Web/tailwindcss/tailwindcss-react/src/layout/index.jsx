import Header from "../Header";
import Aspect from "./aspect";
import BoxDecoration from "./boxDecoration";
import BoxSizing from "./boxSizing";
import Break from "./break";
import ClearFloat from "./clear";
import Columns from "./columns";
import Container from "./container";
import Display from "./display";
import Float from "./float";
import Isolate from "./isolate";
import ObjectFit from "./objectFit";
import ObjectSide from "./objectSide";
import Overflow from "./overflow";
import Overscroll from "./overscroll";
import Placement from "./placement";
import Positon from "./position";
import Visibility from "./visibility";
import ZIndex from "./zIndex";

const Layout = () => (
  <>
    <Header level={1}>布局</Header>
    <Header>纵横比</Header>
    <Aspect />
    <Header>容器</Header>
    <Container />
    <Header>列</Header>
    <Columns />
    <Header>中断</Header>
    <Break />
    <Header>盒子装饰中断</Header>
    <BoxDecoration />
    <Header>盒子尺寸</Header>
    <BoxSizing />
    <Header>展示</Header>
    <Display />
    <Header>浮动</Header>
    <Float />
    <Header>清除</Header>
    <ClearFloat />
    <Header>隔离</Header>
    <Isolate />
    <Header>对象适应</Header>
    <ObjectFit />
    <Header>对象位置</Header>
    <ObjectSide />
    <Header>溢出</Header>
    <Overflow />
    <Header>过度滚动</Header>
    <Overscroll />
    <Header>位置</Header>
    <Positon />
    <Header>上右下左</Header>
    <Placement />
    <Header>可见性</Header>
    <Visibility />
    <Header>Z-index</Header>
    <ZIndex />
  </>
);

export default Layout;
