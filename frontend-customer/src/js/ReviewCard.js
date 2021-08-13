import React from "react";
import Card from "react-bootstrap/Card";

export default function ReviewCard(props) {
  return (
    <div
      className="RestaurantCard"
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
          <Card.Text>Text: {props.text}</Card.Text>
          <Card.Text>Stars: {props.stars}</Card.Text>
          <Card.Text>Restaurant Reply: {props.reply}</Card.Text>
        </Card.Body>
      </Card>
    </div>
  );
}
