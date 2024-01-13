const Spacing = () => {
  return (
    <>
      <div className="grid grid-cols-2">
        <div className="col-span-1 flex space-x-1">
          <div className="w-40 h-40 bg-lime-400">space-x-1</div>
          <div className="w-40 h-40 bg-lime-400">space-x-1</div>
          <div className="w-40 h-40 bg-lime-400">space-x-1</div>
        </div>
        <div className="flex flex-col space-y-1">
          <div className="w-40 h-40 bg-lime-400">space-y-1</div>
          <div className="w-40 h-40 bg-lime-400">space-y-1</div>
          <div className="w-40 h-40 bg-lime-400">space-y-1</div>
        </div>
      </div>
    </>
  );
};

export default Spacing;
