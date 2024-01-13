import Header from "../Header";
import Height from "./height";
import Sizing from "./sizing";
import Width from "./width";

const Size = () => {
  return (
    <>
      <Header>width</Header>
      <Width />
      <Header>height</Header>
      <Height />
      <Header>size</Header>
      <Sizing />
    </>
  );
};

export default Size;
