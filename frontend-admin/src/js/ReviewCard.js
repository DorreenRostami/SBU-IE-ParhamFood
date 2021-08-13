import React, { useState } from "react";
import Button from "react-bootstrap/Button";
import Card from "react-bootstrap/Card";
import axios from "axios";

export default function ReviewCard(props) {
  const [newReply, setNewReply] = useState("");
  const id = props.id;

  function postReply() {
    var info = {
      restaurant_id: id,
      review_id: props.reviewID,
      reply: newReply,
    };
    axios.post("http://localhost:1323/postreply", info);
    //props.action(); // to do
  }

  return (
    <div
      className="ReviewCard"
      style={{
        width: "20rem",
        textAlign: "center",
        display: "flex",
        justifycontent: "center",
        alignitems: "center",
        backgroundColor: "#EA738DFF",
        margin: "20px",
      }}
    >
      <Card
        style={{ width: "18rem", backgroundColor: "#FCF6F5FF", margin: "15px" }}
      >
        <Card.Body>
          <Card.Title>Review ID: {props.reviewID}</Card.Title>
          <Card.Text>Customer ID: {props.customerID}</Card.Text>
          <Card.Text>Review Text: {props.reviewText}</Card.Text>
          <Card.Text>Stars Number: {props.stars}</Card.Text>
          <Card.Text>Reply Text: {props.replyText}</Card.Text>

          <div>
            <input
              type="text"
              name="reply"
              onChange={(e) => setNewReply(e.target.value)}
              style={{ margin: "10px" }}
              placeholder="Reply to this review..."
            />
            <Button
              variant="primary"
              style={{ marginBottom: "10px" }}
              onClick={postReply}
            >
              Post
            </Button>
          </div>
        </Card.Body>
      </Card>
    </div>
  );
}
