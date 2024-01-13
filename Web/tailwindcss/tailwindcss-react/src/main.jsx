import React from "react";
import ReactDOM from "react-dom/client";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import App from "./App.jsx";
import "./index.css";
import Layout from "./layout";
import FlexGrid from "./flex_grid/index.jsx";
import Space from "./space/index.jsx";
import Size from "./size/index.jsx";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
  },
  {
    path: "/layout",
    element: <Layout />,
  },
  {
    path: "/flex_grid",
    element: <FlexGrid />,
  },
  {
    path: "/space",
    element: <Space />,
  },
  {
    path: "/size",
    element: <Size />,
  },
]);

ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
