import React, { useEffect, useState } from "react";
import ReviewCard from "./ReviewCard";

import axios from "axios";
import { Link, useLocation } from "react-router-dom";

export default function Reviews() {
  const [id, setId] = useState(0);
  const [review, setReview] = useState([]);

  const location = useLocation();
  useEffect(() => {
    const info = { restaurant_id: location.state.restaurant_id };
    setId(location.state.restaurant_id);
    axios.post("http://localhost:1323/getreviews", info).then((response) => {
      setReview(response.data.reviews);
    });
  }, []);

  var reviews = [];
  var i;
  var reviewID;
  var customerID;
  var reviewText;
  var stars;
  var replyText;
  for (i = 0; i < review.length; i++) {
    reviewID = review[i].review_id;
    customerID = review[i].customer_id;
    reviewText = review[i].text;
    stars = review[i].stars;
    replyText = review[i].reply;

    reviews.push(
      <ReviewCard
        reviewID={reviewID}
        customerID={customerID}
        reviewText={reviewText}
        stars={stars}
        replyText={replyText}
        id={id}
      />
    );
  }

  return (
    <div className="Reviews">
      <div>
        <Link
          to={{
            pathname: "/profile",
            state: {
              restaurant_id: id,
            },
          }}
          style={{ margin: "4rem" }}
        >
          Back to Profile
        </Link>
      </div>
      {reviews}
    </div>
  );
}
