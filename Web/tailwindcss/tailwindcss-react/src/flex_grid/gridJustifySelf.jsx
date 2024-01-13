const GridJustifySelf = () => {
  return (
    <>
      <div className="grid grid-cols-4">
        <div className="grid justify-items-stretch">
          <div className="bg-lime-300 border-2">01</div>
          <div className="justify-self-start bg-lime-300 border-2">02</div>
          <div className="bg-lime-300 border-2">03</div>
        </div>

        <div className="grid justify-items-stretch">
          <div className="bg-lime-300 border-2">04</div>
          <div className="justify-self-center bg-lime-300 border-2">05</div>
          <div className="bg-lime-300 border-2">06</div>
        </div>

        <div className="grid justify-items-stretch">
          <div className="bg-lime-300 border-2">07</div>
          <div className="justify-self-stretch bg-lime-300 border-2">08</div>
          <div className="bg-lime-300 border-2">09</div>
        </div>

        <div className="grid justify-items-stretch">
          <div className="bg-lime-300 border-2">10</div>
          <div className="justify-self-end bg-lime-300 border-2">11</div>
          <div className="bg-lime-300 border-2">12</div>
        </div>
      </div>
    </>
  );
};

export default GridJustifySelf;
