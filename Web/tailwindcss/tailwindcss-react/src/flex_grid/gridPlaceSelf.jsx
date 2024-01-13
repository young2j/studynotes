const GridPlaceSelf = () => {
  return (
    <>
      <div className="grid grid-cols-4 grid-flow-rows">
        <div className="h-80 w-80 grid grid-cols-2 gap-4  border-2 border-red-100">
          <div className="bg-lime-300 border-2">01</div>
          <div className="place-self-auto bg-lime-300 border-2">self-auto</div>
          <div className="bg-lime-300 border-2">03</div>
        </div>

        <div className="h-80 w-80 grid grid-cols-2 gap-4  border-2 border-red-100">
          <div className="bg-lime-300 border-2">01</div>
          <div className="place-self-start bg-lime-300 border-2">
            self-start
          </div>
          <div className="bg-lime-300 border-2">02</div>
        </div>

        <div className="h-80 w-80 grid grid-cols-2 gap-4  border-2 border-red-100">
          <div className="bg-lime-300 border-2">01</div>
          <div className="place-self-center bg-lime-300 border-2">
            self-center
          </div>
          <div className="bg-lime-300 border-2">02</div>
        </div>

        <div className="h-80 w-80 grid grid-cols-2 gap-4  border-2 border-red-100">
          <div className="bg-lime-300 border-2">01</div>
          <div className="place-self-end bg-lime-300 border-2">self-end</div>
          <div className="bg-lime-300 border-2">02</div>
        </div>

        <div className="h-80 w-80 grid grid-cols-2 gap-4  border-2 border-red-100">
          <div className="bg-lime-300 border-2">01</div>
          <div className="place-self-stretch bg-lime-300 border-2">
            self-stretch
          </div>
          <div className="bg-lime-300 border-2">02</div>
        </div>
      </div>
    </>
  );
};

export default GridPlaceSelf;
