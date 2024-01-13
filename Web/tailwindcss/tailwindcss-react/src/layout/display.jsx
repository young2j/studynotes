const Display = () => {
  return (
    <>
      <div>
        When controlling the flow of text, using the CSS property
        <span className="inline text-red-600">display: inline</span>
        will cause the text inside the element to wrap normally. While using the
        property <span className="inline-block text-red-600">display: inline-block</span>
        will wrap the element to prevent the text inside from extending beyond
        its parent. Lastly, using the property{" "}
        <span className="block text-red-600">display: block</span>
        will put the element on its own line and fill its parent.
      </div>
    </>
  );
};
export default Display;
