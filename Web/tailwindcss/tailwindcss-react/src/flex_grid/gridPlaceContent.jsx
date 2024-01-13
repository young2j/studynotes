const GridPlaceContent = () => {
  return (
    <>
      <div className="grid grid-cols-4 grid-flow-rows">
        <div className="h-80 w-80 grid grid-cols-2 gap-4 place-cotent-start border-2 border-red-100">
          <div className="bg-lime-300 border-2">place-cotent-start</div>
          <div className="bg-lime-300 border-2">place-cotent-start</div>
          <div className="bg-lime-300 border-2">place-cotent-start</div>
        </div>

        <div className="h-80 w-80 grid grid-cols-2 gap-4 place-content-center border-2 border-red-100">
          <div className="bg-lime-300 border-2">place-content-center</div>
          <div className="bg-lime-300 border-2">place-content-center</div>
          <div className="bg-lime-300 border-2">place-content-center</div>
        </div>

        <div className="h-80 w-80 grid grid-cols-2 gap-4 place-content-end border-2 border-red-100">
          <div className="bg-lime-300 border-2">place-content-end</div>
          <div className="bg-lime-300 border-2">place-content-end</div>
          <div className="bg-lime-300 border-2">place-content-end</div>
        </div>

        <div className="h-80 w-80 grid grid-cols-2 gap-4 place-content-between border-2 border-red-100">
          <div className="bg-lime-300 border-2">place-content-between</div>
          <div className="bg-lime-300 border-2">place-content-between</div>
          <div className="bg-lime-300 border-2">place-content-between</div>
        </div>

        <div className="h-80 w-80 grid grid-cols-2 gap-4 place-content-around border-2 border-red-100">
          <div className="bg-lime-300 border-2">place-content-around</div>
          <div className="bg-lime-300 border-2">place-content-around</div>
          <div className="bg-lime-300 border-2">place-content-around</div>
        </div>

        <div className="h-80 w-80 grid grid-cols-2 gap-4 place-content-evenly border-2 border-red-100">
          <div className="bg-lime-300 border-2">place-content-evenly</div>
          <div className="bg-lime-300 border-2">place-content-evenly</div>
          <div className="bg-lime-300 border-2">place-content-evenly</div>
        </div>

        <div className="h-80 w-80 grid grid-cols-2 gap-4 place-content-baseline border-2 border-red-100">
          <div className="bg-lime-300 border-2">place-content-baseline</div>
          <div className="bg-lime-300 border-2">place-content-baseline</div>
          <div className="bg-lime-300 border-2">place-content-baseline</div>
        </div>

        <div className="h-80 w-80 grid grid-cols-2 gap-4 place-content-stretch border-2 border-red-100">
          <div className="bg-lime-300 border-2">place-content-stretch</div>
          <div className="bg-lime-300 border-2">place-content-stretch</div>
          <div className="bg-lime-300 border-2">place-content-stretch</div>
        </div>
      </div>
    </>
  );
};

export default GridPlaceContent;
