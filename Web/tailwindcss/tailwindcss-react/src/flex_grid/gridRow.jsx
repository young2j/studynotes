const GridRow = () => {
  return (
    <>
      <div className="grid grid-rows-4 grid-flow-col gap-4">
        <div className="bg-lime-300 border-2">01</div>
        <div className="col-span-2 bg-lime-300 border-2">02</div>
        <div className="bg-lime-300 border-2">03</div>
        <div className="bg-lime-300 border-2">04</div>
        <div className="bg-lime-300 border-2">05</div>
        <div className="col-span-2 bg-lime-300 border-2">06</div>
        <div className="bg-lime-300 border-2">07</div>
        <div className="bg-lime-300 border-2">08</div>
        <div className="grid grid-rows-subgrid gap-4 row-span-4">
          <div className="row-start-2 bg-lime-300 border-2">09</div>
          <div className="row-end-4 bg-lime-300 border-2">10</div>
        </div>
      </div>
    </>
  );
};

export default GridRow;
