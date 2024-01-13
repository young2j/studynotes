import Header from "../Header";
import Flex from "./flex";
import FlexBasis from "./flexBasis";
import FlexCol from "./flexCol";
import FlexGrow from "./flexGrow";
import FlexItems from "./flexItems";
import FlexItemsSelf from "./flexItemsSelf";
import FlexJustify from "./flexJustify";
import FlexOrder from "./flexOrder";
import FlexShrink from "./flexShrink";
import FlexWrap from "./flexWrap";
import GridAutoCol from "./gridAutoCol";
import GridAutoRow from "./gridAutoRow";
import GridCol from "./gridCol";
import GridContent from "./gridContent";
import GridGap from "./gridGap";
import GridJustifyItems from "./gridJustifyItems";
import GridJustifySelf from "./gridJustifySelf";
import GridPlaceContent from "./gridPlaceContent";
import GridPlaceItems from "./gridPlaceItems";
import GridPlaceSelf from "./gridPlaceSelf";
import GridRow from "./gridRow";

const FlexGrid = () => {
  return (
    <>
      <Header level={1}>弹性盒&网格</Header>
      <Header>flex-basis</Header>
      <FlexBasis />
      <Header>flex-direction</Header>
      <FlexCol />
      <Header>flex-wrap</Header>
      <FlexWrap />
      <Header>flex</Header>
      <Flex />
      <Header>flex-grow</Header>
      <FlexGrow />
      <Header>flex-shrink</Header>
      <FlexShrink />
      <Header>order</Header>
      <FlexOrder />
      <Header>flex-justify</Header>
      <FlexJustify />
      <Header>flex-items</Header>
      <FlexItems />
      <Header>flex-items-self</Header>
      <FlexItemsSelf />

      <br />
      <br />
      <br />
      <Header>grid-col</Header>
      <GridCol />
      <Header>grid-row</Header>
      <GridRow />
      <Header>grid-auto-cols</Header>
      <GridAutoCol />
      <Header>grid-auto-rows</Header>
      <GridAutoRow />
      <Header>grid-gap</Header>
      <GridGap />
      <Header>grid-justify-items</Header>
      <GridJustifyItems />
      <Header>grid-justify-self</Header>
      <GridJustifySelf />
      <Header>grid-content</Header>
      <GridContent />
      <Header>grid-place-content</Header>
      <GridPlaceContent />
      <Header>grid-place-items</Header>
      <GridPlaceItems />
      <Header>grid-place-self</Header>
      <GridPlaceSelf />
    </>
  );
};

export default FlexGrid;
