import React, { useState } from "react";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import "../css/Signup.css";
import axios from "axios";
import { Link, useHistory } from "react-router-dom";

export default function Signup() {
  const [phoneNumber, setPhoneNumber] = useState("");
  const [password, setPassword] = useState("");

  const [name, setName] = useState(""); // customer name
  const [district, setDistrict] = useState("");
  const [address, setAddress] = useState("");

  const [id, setId] = useState(0);

  function validateForm() {
    // return phoneNumber.length > 0 && password.length > 7 && /^(?:[A-Za-z]+|\d+)$/.test(password);
    return phoneNumber.length > 0 && password.length > 7;
  }

  let history = useHistory();

  function handleSubmit(event) {
    event.preventDefault();

    const info = {
      mobile: phoneNumber,
      password: password,
      name: name,
      district: district,
      address: address,
    };
    axios.post("http://localhost:1323/signupcustomer", info).then(
      (response) => {
        history.push({
          pathname: "/profile",
          state: {
            customer_id: response.data.id,
          },
        });
      },
      (error) => {
        console.log(error);
      }
    );
  }

  return (
    <div className="Signup">
      <Form
        onSubmit={handleSubmit}
        style={{ backgroundColor: "#EA738DFF", padding: "20px" }}
      >
        <Form.Group size="lg" controlId="phoneNumber">
          <Form.Label>Phone Number: </Form.Label>
          <Form.Control
            autoFocus
            type="phoneNumber"
            value={phoneNumber}
            onChange={(e) => setPhoneNumber(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>

        <Form.Group size="lg" controlId="password">
          <Form.Label>Password: </Form.Label>
          <Form.Control
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>

        <Form.Group size="lg" controlId="name">
          <Form.Label>Your Name: </Form.Label>
          <Form.Control
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>

        <Form.Group size="lg" controlId="district">
          <Form.Label>Your District: </Form.Label>
          <Form.Control
            type="text"
            value={district}
            onChange={(e) => setDistrict(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>

        <Form.Group size="lg" controlId="address">
          <Form.Label>Your Address: </Form.Label>
          <Form.Control
            type="text"
            value={address}
            onChange={(e) => setAddress(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>

        <Button block size="lg" type="submit" disabled={!validateForm()}>
          Signup
        </Button>
      </Form>

      <div className="loginLink" style={{ textAlign: "center" }}>
        <Link to="/">Back to Login</Link>
      </div>
    </div>
  );
}
