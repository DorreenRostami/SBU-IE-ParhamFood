import React from "react";
import Card from "react-bootstrap/Card";
import { useHistory } from "react-router-dom";
import Button from "react-bootstrap/Button";

export default function RestaurantCard(props) {
  const history = useHistory();

  function order() {
    history.push({
      pathname: "/profile/restaurantpage",
      state: {
        res_id: props.restaurant_id,
        customer_id: props.customer_id,
      },
    });
    console.log(props.restaurant_id);
  }

  function seeReveiws() {
    history.push({
      pathname: "/profile/restaurantreviewpage",
      state: {
        res_id: props.restaurant_id,
        customer_id: props.customer_id,
      },
    });
  }

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
          <Card.Title>Restaurant ID: {props.restaurant_id}</Card.Title>
          <Card.Text>Restaurant Name: {props.name}</Card.Text>
          <Card.Text>Restaurant District: {props.district}</Card.Text>
          <Card.Text>Restarant Address: {props.address}</Card.Text>

          <div>
            <Button
              variant="primary"
              style={{ marginBottom: "10px" }}
              onClick={order}
            >
              Get menu and Order from this restaraunt
            </Button>
          </div>

          <div>
            <Button
              variant="primary"
              style={{ marginBottom: "10px" }}
              onClick={seeReveiws}
            >
              See reviews for this restaraunt
            </Button>
          </div>
        </Card.Body>
      </Card>
    </div>
  );
}
