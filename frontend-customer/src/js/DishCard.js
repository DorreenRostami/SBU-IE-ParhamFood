import React, { useState, useHistory } from "react";
import Button from "react-bootstrap/Button";
import Card from "react-bootstrap/Card";
import axios from "axios";
import "../css/Menu.css";

export default function DishCard(props) {
  const [dishNumber, setDishNumber] = useState(0);

  var status;
  if (props.available) {
    status = "available";
  } else {
    status = "not available";
  }

  return (
    <div
      className="Dishcard"
      style={{
        width: "20rem",
        textAlign: "center",
        display: "flex",
        justifycontent: "center",
        alignitems: "center",
        backgroundColor: "#EA738DFF",
        margin: "20px",
        marginLeft: "38%",
      }}
    >
      <Card
        style={{
          width: "18rem",
          backgroundColor: "#FCF6F5FF",
          margin: "20px",
        }}
      >
        <Card.Body>
          <Card.Title>{props.name}</Card.Title>
          <Card.Text>Food price: {props.price}</Card.Text>
          <Card.Text>This food is {status}.</Card.Text>
          <input
            type="number"
            name="dishNumber"
            id="dishNumber"
            onChange={(e) => {
              setDishNumber(e.target.value);
            }}
            style={{ margin: "10px" }}
            placeholder="Enter the number to order"
          />
        </Card.Body>
      </Card>
    </div>
  );
}
