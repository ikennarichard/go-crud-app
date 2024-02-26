import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from "./pages/Home";
import Layout from "./pages/Layout";
import Employee from "./components/UpdateEmployee";

function App() {
  return (
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Layout />}>
            <Route index element={<Home />} />
            <Route path="/employee/:id" element={<Employee />} />
          </Route>
        </Routes>
      </BrowserRouter>
  );
}

export default App;
