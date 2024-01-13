const FlexItems = () => {
  return (
    <>
      <div className="columns-5">
        <div className="flex items-stretch">
          <div className="bg-lime-300 border-2">items-stretch</div>
          <div className="bg-lime-300 border-2">items-stretch</div>
          <div className="bg-lime-300 border-2">items-stretch</div>
        </div>

        <div className="flex items-start">
          <div className="bg-lime-300 border-2">items-start</div>
          <div className="bg-lime-300 border-2">items-start</div>
          <div className="bg-lime-300 border-2">items-start</div>
        </div>

        <div className="flex items-center">
          <div className="bg-lime-300 border-2">items-center</div>
          <div className="bg-lime-300 border-2">items-center</div>
          <div className="bg-lime-300 border-2">items-center</div>
        </div>

        <div className="flex items-end">
          <div className="bg-lime-300 border-2">items-end</div>
          <div className="bg-lime-300 border-2">items-end</div>
          <div className="bg-lime-300 border-2">items-end</div>
        </div>

        <div className="flex items-baseline">
          <div className="bg-lime-300 border-2">items-baseline</div>
          <div className="bg-lime-300 border-2">items-baseline</div>
          <div className="bg-lime-300 border-2">items-baseline</div>
        </div>
      </div>
    </>
  );
};

export default FlexItems;
