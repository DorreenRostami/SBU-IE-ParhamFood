import React, { useState, useEffect } from "react";
import "../css/Profile.css";

import { Link, useLocation } from "react-router-dom";

export default function Profile() {
  const [id, setId] = useState(0);
  const location = useLocation();
  useEffect(() => {
    setId(location.state.restaurant_id);
  }, []);

  return (
    <div className="Profile" style={{ textAlign: "center" }}>
      <div
        style={{ backgroundColor: "#ADD8E6", width: "20rem", margin: "1rem" }}
      >
        <Link
          to={{
            pathname: "/profile/menu",
            state: {
              restaurant_id: id,
            },
          }}
        >
          Menu
        </Link>
      </div>
      <div
        style={{ backgroundColor: "#ADD8E6", width: "20rem", margin: "1rem" }}
      >
        <Link
          to={{
            pathname: "/profile/orders",
            state: {
              restaurant_id: id,
            },
          }}
        >
          Orders
        </Link>
      </div>
      <div
        style={{ backgroundColor: "#ADD8E6", width: "20rem", margin: "1rem" }}
      >
        <Link
          to={{
            pathname: "/profile/reviews",
            state: {
              restaurant_id: id,
            },
          }}
        >
          Reviews
        </Link>
      </div>

      <div
        style={{ backgroundColor: "#ADD8E6", width: "20rem", margin: "1rem" }}
      >
        <Link
          to={{
            pathname: "/profile/edit",
            state: {
              restaurant_id: id,
            },
          }}
        >
          Get and Edit Restaurant Information
        </Link>
      </div>
    </div>
  );
}
