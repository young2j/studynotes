const FlexShrink = () => {
  return (
    <>
      <div className="flex">
        <div className="flex-none bg-lime-300 border-2 w-14">flex-none</div>
        <div className="shrink bg-lime-300 border-2 w-14">shrink</div>
        <div className="shrink-0 bg-lime-300 border-2 w-14">shrink-0</div>
        <div className="flex-1 bg-lime-300 border-2 w-14">flex-1</div>
      </div>
    </>
  );
};

export default FlexShrink;
