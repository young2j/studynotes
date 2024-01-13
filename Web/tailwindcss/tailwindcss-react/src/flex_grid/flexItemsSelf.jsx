const FlexItemsSelf = () => {
  return (
    <>
      <div className="columns-5">
        <div className="flex items-stretch">
          <div className="bg-lime-300 border-2">items-stretch</div>
          <div className="self-stretch bg-lime-300 border-2">self-stretch</div>
          <div className="bg-lime-300 border-2">items-stretch</div>
        </div>

        <div className="flex items-stretch">
          <div className="bg-lime-300 border-2">items-stretch</div>
          <div className="self-start bg-lime-300 border-2">self-start</div>
          <div className="bg-lime-300 border-2">items-stretch</div>
        </div>

        <div className="flex items-stretch">
          <div className="bg-lime-300 border-2">items-stretch</div>
          <div className="self-center bg-lime-300 border-2">self-center</div>
          <div className="bg-lime-300 border-2">items-stretch</div>
        </div>

        <div className="flex items-stretch">
          <div className="bg-lime-300 border-2">items-stretch</div>
          <div className="self-end bg-lime-300 border-2">self-end</div>
          <div className="bg-lime-300 border-2">items-stretch</div>
        </div>

        <div className="flex items-stretch">
          <div className="bg-lime-300 border-2">items-stretch</div>
          <div className="self-auto bg-lime-300 border-2">self-auto</div>
          <div className="bg-lime-300 border-2">items-stretch</div>
        </div>
      </div>
    </>
  );
};

export default FlexItemsSelf;
