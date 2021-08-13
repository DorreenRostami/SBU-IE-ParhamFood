import Login from "./Login";
import Signup from "./Signup";
import Profile from "./Profile";
import RestaurantPage from "./RestaurantPage";
import RestaurantReviewPage from "./RestaurantReviewPage";
import Edit from "./Edit";
import OrdersHistory from "./OrdersHistory";

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

            <Route exact path="/profile/edit">
              <Edit />
            </Route>

            <Route exact path="/profile/restaurantpage">
              <RestaurantPage />
            </Route>

            <Route exact path="/profile/restaurantreviewpage">
              <RestaurantReviewPage />
            </Route>

            <Route exact path="/profile/ordershistory">
              <OrdersHistory />
            </Route>
          </Switch>
        </div>
      </div>
    </Router>
  );
}

export default App;
