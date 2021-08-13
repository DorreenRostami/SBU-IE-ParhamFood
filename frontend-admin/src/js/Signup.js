import React, { useState } from "react";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import "../css/Signup.css";
import axios from "axios";
import { Link, useHistory } from "react-router-dom";

export default function Signup() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [name, setName] = useState(""); // restaurant name
  const [district, setDistrict] = useState("");
  const [address, setAddress] = useState("");
  const [open, setOpen] = useState(0);
  const [close, setClose] = useState(0);
  const [fixedCost, setFixedCost] = useState(0);
  const [fixedTime, setFixedTime] = useState(0);

  const [id, setId] = useState(0);

  function validateForm() {
    // return email.length > 0 && password.length > 7 && /^(?:[A-Za-z]+|\d+)$/.test(password);
    return email.length > 0 && password.length > 7;
  }

  let history = useHistory();

  function handleSubmit(event) {
    event.preventDefault();

    const info = {
      email: email,
      password: password,
      name: name,
      district: district,
      address: address,
      open: parseInt(open, 10),
      close: parseInt(close, 10),
      fixed_cost: parseInt(fixedCost, 10),
      fixed_minute: parseInt(fixedTime, 10),
    };
    axios.post("http://localhost:1323/signupadmin", info).then(
      (response) => {
        history.push({
          pathname: "/profile",
          state: {
            restaurant_id: response.data.restaurant_id,
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
        <Form.Group size="lg" controlId="email">
          <Form.Label>Email: </Form.Label>
          <Form.Control
            autoFocus
            placeholder="name@example.com"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
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
          <Form.Label>Restaurant Name: </Form.Label>
          <Form.Control
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>

        <Form.Group size="lg" controlId="district">
          <Form.Label>Restaurant District: </Form.Label>
          <Form.Control
            type="text"
            value={district}
            onChange={(e) => setDistrict(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>

        <Form.Group size="lg" controlId="address">
          <Form.Label>Restaurant Address: </Form.Label>
          <Form.Control
            type="text"
            value={address}
            onChange={(e) => setAddress(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>

        <Form.Group size="lg" controlId="open">
          <Form.Label>Restaurant Open Time: </Form.Label>
          <Form.Control
            type="number"
            value={open}
            onChange={(e) => setOpen(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>

        <Form.Group size="lg" controlId="close">
          <Form.Label>Restaurant Close Time: </Form.Label>
          <Form.Control
            type="number"
            value={close}
            onChange={(e) => setClose(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>

        <Form.Group size="lg" controlId="fixedCost">
          <Form.Label>Restaurant Fixed Cost: </Form.Label>
          <Form.Control
            type="number"
            value={fixedCost}
            onChange={(e) => setFixedCost(e.target.value)}
            style={{ display: "block", margin: "2px" }}
          />
        </Form.Group>

        <Form.Group size="lg" controlId="fixedTime">
          <Form.Label>Restaurant Fixed Time: </Form.Label>
          <Form.Control
            type="number"
            value={fixedTime}
            onChange={(e) => setFixedTime(e.target.value)}
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
