const FlexOrder = () => {
  return (
    <>
      <div className="flex justify-between">
        <div className="order-last bg-lime-300 border-2 w-14">01(order-last)</div>
        <div className="bg-lime-300 border-2 w-14">02</div>
        <div className="bg-lime-300 border-2 w-14">03</div>
      </div>
    </>
  );
};

export default FlexOrder;
