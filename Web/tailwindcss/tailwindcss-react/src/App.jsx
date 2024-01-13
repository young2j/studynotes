import { Link } from "react-router-dom";

function App() {
  return (
    <>
      <div className="min-h-screen min-w-screen flex items-center justify-center">
        <div className="container mx-auto my-auto w-60">
          <ol className="list-decimal">
            <li>
              <Link to="/layout">布局</Link>
            </li>
            <li>
              <Link to="/flex_grid">弹性盒&网格</Link>
            </li>
            <li>
              <Link to="/space">间距</Link>
            </li>
            <li>
              <Link to="/size">尺寸</Link>
            </li>
          </ol>
        </div>
      </div>
    </>
  );
}

export default App;
