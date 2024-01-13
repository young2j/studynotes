import Header from "../Header";
import Margin from "./margin";
import Padding from "./padding";
import Spacing from "./spacing";

const Space = () => {
  return (
    <>
      <Header>padding</Header>
      <Padding />
      <Header>margin</Header>
      <Margin />
      <Header>spacing</Header>
      <Spacing />
    </>
  );
};

export default Space;
