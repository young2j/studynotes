// eslint-disable-next-line react/prop-types
const Header = ({ children, level = 2 }) => {
  return (
    <>
      {level === 1 ? (
        <h1 className="text-center w-full mx-auto border-t border-b bg-blue-400">
          {children}
        </h1>
      ) : (
        <h2 className="text-center w-full mx-auto border-t border-b bg-blue-200">
          {children}
        </h2>
      )}
    </>
  );
};
export default Header;
