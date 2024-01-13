const Positon = () => {
  return (
    <>
      <div className="static">
        <p>Static parent</p>
        <div className="absolute bottom-0 left-0 ">
          <p>Absolute child</p>
        </div>
      </div>

      <div className="static ">
        <div className="static ">
          <p>Static child</p>
        </div>
        <div className="inline-block ">
          <p>Static sibling</p>
        </div>

        <div className="absolute ">
          <p>Absolute child</p>
        </div>
        <div className="inline-block ">
          <p>Static sibling</p>
        </div>
      </div>
    </>
  );
};

export default Positon;
