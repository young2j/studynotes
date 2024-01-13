const FlexGrow = () => {
  return (
    <>
      <div className="flex">
        <div className="flex-none bg-lime-300 border-2 w-14">flex-none</div>
        <div className="grow bg-lime-300 border-2 w-14">grow</div>
        <div className="grow-0 bg-lime-300 border-2 w-14">grow-0</div>
      </div>
    </>
  );
};

export default FlexGrow;
