import React, { useState, useEffect } from "react";
import "../css/Profile.css";
import axios from "axios";
import ReviewCard from "./ReviewCard";
import Button from "react-bootstrap/Button";

import { Link, useLocation } from "react-router-dom";

export default function RestaurantReviewPage() {
  const [id, setId] = useState(0);
  const [resId, setResId] = useState(0);
  const location = useLocation();

  const [review, setReview] = useState([]);
  const [newReviewText, setNewReviewText] = useState("");
  const [newReviewStars, setNewReviewStars] = useState(0);

  useEffect(() => {
    setId(location.state.customer_id);
    setResId(location.state.res_id);
    const info = { restaurant_id: location.state.res_id };
    axios
      .post("http://localhost:1323/restaurantreviews", info)
      .then((response) => {
        setReview(response.data.reviews);
      });
  }, []);

  var reviews = [];
  var i;
  var reviewID;
  var customerID;
  var text;
  var stars;
  var reply;

  for (i = 0; i < review.length; i++) {
    reviewID = review[i].review_id;
    customerID = review[i].customer_id;
    text = review[i].text;
    stars = review[i].stars;
    reply = review[i].reply;

    reviews.push(
      <ReviewCard
        reviewID={reviewID}
        customerID={customerID}
        text={text}
        stars={stars}
        reply={reply}
      />
    );
  }

  function postReview() {
    if (parseInt(newReviewStars, 10) <= 5 && newReviewText.length <= 140) {
      const info = {
        customer_id: id,
        restaurant_id: resId,
        text: newReviewText,
        stars: parseInt(newReviewStars, 10),
      };
      axios.post("http://localhost:1323/postreview", info);
    }
  }

  return (
    <div className="Profile">
      <div
        style={{
          backgroundColor: "#ADD8E6",
          width: "20rem",
          margin: "2rem",
        }}
      >
        <Link
          to={{
            pathname: "/profile",
            state: {
              customer_id: id,
            },
          }}
        >
          Back to Profile
        </Link>
      </div>
      <div>
        <input
          type="text"
          name="text"
          onChange={(e) => setNewReviewText(e.target.value)}
          style={{ margin: "10px" }}
          placeholder="Enter your text review..."
        />
        <input
          type="number"
          name="stars"
          onChange={(e) => setNewReviewStars(e.target.value)}
          style={{ margin: "10px" }}
          placeholder="Enter number of stars..."
        />
        <Button
          variant="primary"
          style={{ margin: "10px" }}
          onClick={postReview}
        >
          Post Review
        </Button>
      </div>
      <div>{reviews}</div>
    </div>
  );
}
