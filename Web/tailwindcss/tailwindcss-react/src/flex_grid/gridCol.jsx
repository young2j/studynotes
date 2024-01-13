const GridCol = () => {
  return (
    <>
      <div className="grid grid-cols-4 gap-4">
        <div className="bg-lime-300 border-2">01</div>
        <div className="col-span-2 bg-lime-300 border-2">02</div>
        <div className="bg-lime-300 border-2">03</div>
        <div className="bg-lime-300 border-2">04</div>
        <div className="bg-lime-300 border-2">05</div>
        <div className="col-span-2 bg-lime-300 border-2">06</div>
        <div className="bg-lime-300 border-2">07</div>
        <div className="bg-lime-300 border-2">08</div>
        <div className="grid grid-cols-subgrid gap-4 col-span-4">
          <div className="col-start-2 bg-lime-300 border-2">09</div>
          <div className="col-end-15 bg-lime-300 border-2">10</div>
        </div>
      </div>
    </>
  );
};

export default GridCol;
