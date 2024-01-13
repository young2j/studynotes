const GridContent = () => {
  return (
    <>
      <div className="grid grid-cols-1 grid-flow-rows">
        <div className="h-36 grid grid-cols-2 gap-4 content-start border-2 border-red-100">
          <div className="bg-lime-300 border-2">content-start</div>
          <div className="bg-lime-300 border-2">content-start</div>
          <div className="bg-lime-300 border-2">content-start</div>
        </div>

        <div className="h-36 grid grid-cols-2 gap-4 content-center border-2 border-red-100">
          <div className="bg-lime-300 border-2">content-center</div>
          <div className="bg-lime-300 border-2">content-center</div>
          <div className="bg-lime-300 border-2">content-center</div>
        </div>

        <div className="h-36 grid grid-cols-2 gap-4 content-end border-2 border-red-100">
          <div className="bg-lime-300 border-2">content-end</div>
          <div className="bg-lime-300 border-2">content-end</div>
          <div className="bg-lime-300 border-2">content-end</div>
        </div>

        <div className="h-36 grid grid-cols-2 gap-4 content-between border-2 border-red-100">
          <div className="bg-lime-300 border-2">content-between</div>
          <div className="bg-lime-300 border-2">content-between</div>
          <div className="bg-lime-300 border-2">content-between</div>
        </div>

        <div className="h-36 grid grid-cols-2 gap-4 content-around border-2 border-red-100">
          <div className="bg-lime-300 border-2">content-around</div>
          <div className="bg-lime-300 border-2">content-around</div>
          <div className="bg-lime-300 border-2">content-around</div>
        </div>

        <div className="h-36 grid grid-cols-2 gap-4 content-evenly border-2 border-red-100">
          <div className="bg-lime-300 border-2">content-evenly</div>
          <div className="bg-lime-300 border-2">content-evenly</div>
          <div className="bg-lime-300 border-2">content-evenly</div>
        </div>

        <div className="h-36 grid grid-cols-2 gap-4 content-normal border-2 border-red-100">
          <div className="bg-lime-300 border-2">content-normal</div>
          <div className="bg-lime-300 border-2">content-normal</div>
          <div className="bg-lime-300 border-2">content-normal</div>
        </div>

        <div className="h-36 grid grid-cols-2 gap-4 content-stretch border-2 border-red-100">
          <div className="bg-lime-300 border-2">content-stretch</div>
          <div className="bg-lime-300 border-2">content-stretch</div>
          <div className="bg-lime-300 border-2">content-stretch</div>
        </div>
      </div>
    </>
  );
};

export default GridContent;
