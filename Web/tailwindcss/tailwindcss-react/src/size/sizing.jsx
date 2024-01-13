const Sizing = () => {
  return (
    <>
      <div className="rows-2 space-y-1">
        <div className="row-span-1 flex space-x-1">
          <div className="size-16 bg-lime-400">size-16</div>
          <div className="size-40 bg-lime-400">siez-40</div>
          <div className="size-1/2 bg-lime-400">size-1/2</div>
        </div>

        <div className="space-y-1">
          <div className="size-min  bg-lime-400">size-min</div>
          <div className="size-full  bg-lime-400">size-full</div>
          <div className="size-fit  bg-lime-400">size-fit</div>
        </div>
      </div>
    </>
  );
};

export default Sizing;
