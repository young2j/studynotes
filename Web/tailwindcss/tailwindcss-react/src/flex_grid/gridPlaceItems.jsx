const GridPlaceItems = () => {
  return (
    <>
      <div className="grid grid-cols-4 grid-flow-rows">
        <div className="h-80 w-80 grid grid-cols-2 gap-4 place-items-start border-2 border-red-100">
          <div className="bg-lime-300 border-2">place-items-start</div>
          <div className="bg-lime-300 border-2">place-items-start</div>
          <div className="bg-lime-300 border-2">place-items-start</div>
        </div>

        <div className="h-80 w-80 grid grid-cols-2 gap-4 place-items-center border-2 border-red-100">
          <div className="bg-lime-300 border-2">place-items-center</div>
          <div className="bg-lime-300 border-2">place-items-center</div>
          <div className="bg-lime-300 border-2">place-items-center</div>
        </div>


        <div className="h-80 w-80 grid grid-cols-2 gap-4 place-items-end border-2 border-red-100">
          <div className="bg-lime-300 border-2">place-items-end</div>
          <div className="bg-lime-300 border-2">place-items-end</div>
          <div className="bg-lime-300 border-2">place-items-end</div>
        </div>

        <div className="h-80 w-80 grid grid-cols-2 gap-4 place-items-baseline border-2 border-red-100">
          <div className="bg-lime-300 border-2">place-items-baseline</div>
          <div className="bg-lime-300 border-2">place-items-baseline</div>
          <div className="bg-lime-300 border-2">place-items-baseline</div>
        </div>

        <div className="h-80 w-80 grid grid-cols-2 gap-4 place-items-stretch border-2 border-red-100">
          <div className="bg-lime-300 border-2">place-items-stretch</div>
          <div className="bg-lime-300 border-2">place-items-stretch</div>
          <div className="bg-lime-300 border-2">place-items-stretch</div>
        </div>
      </div>
    </>
  );
};

export default GridPlaceItems;
