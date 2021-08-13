import React, { useState, useHistory } from "react";
import Button from "react-bootstrap/Button";
import Card from "react-bootstrap/Card";
import axios from "axios";
import "../css/Menu.css";

export default function Dishcard(props) {
  const [newPrice, setNewPrice] = useState(0);
  const [newName, setNewName] = useState("");

  function deleteDish() {
    const info = {
      restaurant_id: props.id,
      name: props.name,
      price: 0,
      available: true,
    };
    axios.post("http://localhost:1323/deletedish", info);
    props.action();
  }

  function changeAvailable() {
    if (props.availibility == "available") {
      var info = {
        restaurant_id: props.id,
        name: props.name,
        price: props.price,
        available: false,
      };
    } else {
      info = {
        restaurant_id: props.id,
        name: props.name,
        price: props.price,
        available: true,
      };
    }
    axios.post("http://localhost:1323/updatedishpa", info);
    //props.action(); // to do
  }

  function changePrice() {
    var temp = parseInt(newPrice, 10);

    if (props.availibility == "available") {
      var info = {
        restaurant_id: props.id,
        name: props.name,
        price: temp,
        available: true,
      };
    } else {
      info = {
        restaurant_id: props.id,
        name: props.name,
        price: temp,
        available: false,
      };
    }
    axios.post("http://localhost:1323/updatedishpa", info);
    //props.action(); // to do
  }

  function changeName() {
    var info = {
      restaurant_id: props.id,
      old_name: props.name,
      new_name: newName,
    };
    axios.post("http://localhost:1323/updatedishname", info);
    //props.action(); // to do
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
      }}
    >
      <Card
        style={{ width: "18rem", backgroundColor: "#FCF6F5FF", margin: "20px" }}
      >
        <Card.Body>
          <Card.Title>{props.name}</Card.Title>
          <Card.Text>Food price: {props.price}</Card.Text>
          <Card.Text>This food is {props.availibility}.</Card.Text>
          <Button
            variant="primary"
            style={{ marginTop: "10px" }}
            onClick={deleteDish}
          >
            Delete this dish
          </Button>
          <Button
            variant="primary"
            style={{ margin: "10px" }}
            onClick={changeAvailable}
          >
            available/unavailable
          </Button>
          <div>
            <input
              type="text"
              name="newName"
              onChange={(e) => setNewName(e.target.value)}
              style={{ margin: "10px" }}
              placeholder="Enter new name"
            />
            <Button
              variant="primary"
              style={{ marginBottom: "10px" }}
              onClick={changeName}
            >
              Edit name
            </Button>
          </div>
          <div>
            <input
              type="number"
              name="newPrice"
              id="newPrice"
              onChange={(e) => {
                setNewPrice(e.target.value);
              }}
              style={{ margin: "10px" }}
              placeholder="Enter new price"
            />
            <Button
              variant="primary"
              style={{ marginBottom: "10px" }}
              onClick={changePrice}
            >
              Edit price
            </Button>
          </div>
        </Card.Body>
      </Card>
    </div>
  );
}
