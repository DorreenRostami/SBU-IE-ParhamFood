import Login from "./Login";
import Signup from "./Signup";
import Profile from "./Profile";
import Menu from "./Menu";
import Orders from "./Orders";
import Edit from "./Edit";
import Reviews from "./Reviews";

import { BrowserRouter as Router, Route, Switch } from "react-router-dom";

function App() {
  return (
    <Router>
      <div className="App">
        <div className="content">
          <Switch>
            <Route exact path="/">
              <Login />
            </Route>

            <Route exact path="/signup">
              <Signup />
            </Route>

            <Route exact path="/profile">
              <Profile />
            </Route>

            <Route exact path="/profile/menu">
              <Menu />
            </Route>

            <Route exact path="/profile/orders">
              <Orders />
            </Route>

            <Route exact path="/profile/edit">
              <Edit />
            </Route>

            <Route exact path="/profile/reviews">
              <Reviews />
            </Route>
          </Switch>
        </div>
      </div>
    </Router>
  );
}

export default App;
