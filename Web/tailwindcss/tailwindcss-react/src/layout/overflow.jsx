const Overflow = () => {
  return (
    <>
      <div className="container mx-auto overflow-scroll h-16 w-2/3">
        <p>这是一段内容，如果它的高度超过了容器的高度，就会出现滚动条。</p>
        <p>这是一段内容，如果它的宽度超过了容器的宽度，就会出现滚动条。</p>
        <p>您可以继续添加更多的内容，以测试滚动效果。</p>
      </div>
    </>
  );
};

export default Overflow;
