const Height = () => {
  return (
    <>
      <div className="rows-4 space-y-1">
        <div className="row-span-1 flex space-x-1">
          <div className="h-24 w-40 bg-lime-400">h-24</div>
          <div className="h-48 w-40 bg-lime-400">h-48</div>
          <div className="h-1/2 w-40 bg-lime-400">h-1/2</div>
        </div>

        <div className="space-y-1">
          <div className="h-screen w-40 bg-lime-400">h-screen</div>
          <div className="h-full w-40 bg-lime-400">h-full</div>
          <div className="h-min w-40 bg-lime-400">h-min</div>
        </div>

        <div className="space-y-1">
          <div className="min-h-24 h-24 w-40 bg-lime-400">min-h-24</div>
          <div className="min-h-48 h-1/2 w-40 bg-lime-400">min-h-48</div>
          <div className="min-h-full w-40 bg-lime-400">min-h-full</div>
        </div>

        <div className="space-y-1">
          <div className="max-h-24 w-40 bg-lime-400">max-h-24</div>
          <div className="max-h-48 w-40 bg-lime-400">max-h-48</div>
          <div className="max-h-full w-40 bg-lime-400">max-h-full</div>
        </div>
      </div>
    </>
  );
};

export default Height;
