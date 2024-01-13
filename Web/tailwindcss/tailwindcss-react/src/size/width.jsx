const Width = () => {
  return (
    <>
      <div className="rows-4 space-y-1">
        <div className="row-span-1 flex space-x-1">
          <div className="w-24 h-40 bg-lime-400">w-24</div>
          <div className="w-48 h-40 bg-lime-400">w-48</div>
          <div className="w-1/2 h-40 bg-lime-400">w-1/2</div>
        </div>

        <div className="space-y-1">
          <div className="w-screen h-40 bg-lime-400">w-screen</div>
          <div className="w-full h-40 bg-lime-400">w-full</div>
          <div className="w-min h-40 bg-lime-400">w-min</div>
        </div>

        <div className="space-y-1">
          <div className="min-w-24 w-24 h-40 bg-lime-400">min-w-24</div>
          <div className="min-w-48 w-1/2 h-40 bg-lime-400">min-w-48</div>
          <div className="min-w-full h-40 bg-lime-400">min-w-full</div>
        </div>

        <div className="space-y-1">
          <div className="max-w-24 h-40 bg-lime-400">max-w-24</div>
          <div className="max-w-48 h-40 bg-lime-400">max-w-48</div>
          <div className="max-w-full h-40 bg-lime-400">max-w-full</div>
        </div>
      </div>
    </>
  );
};

export default Width;
