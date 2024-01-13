const FlexJustify = () => {
  return (
    <>
      <div>
        <div className="flex justify-start">
          <div className="bg-lime-300 border-2">01</div>
          <div className="bg-lime-300 border-2">02</div>
          <div className="bg-lime-300 border-2">03</div>
        </div>

        <div className="flex justify-center">
          <div className="bg-lime-300 border-2">04</div>
          <div className="bg-lime-300 border-2">05</div>
          <div className="bg-lime-300 border-2">06</div>
        </div>

        <div className="flex justify-end">
          <div className="bg-lime-300 border-2">07</div>
          <div className="bg-lime-300 border-2">08</div>
          <div className="bg-lime-300 border-2">09</div>
        </div>
      </div>
    </>
  );
};

export default FlexJustify;
